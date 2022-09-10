package model

import (
  "testing"
)

// Latch
//
// a ---> +------+
//        | NAND |----+----------------> out
//   +--> +------+    |
//   |                +--> +------+
//   |                     | NAND |---+
//   |              b ---> +------+   |
//   +--------------------------------+
//
//  a   b   | q
//  --------|----
//  LO  LO  | HI (dont use)
//  LO  HI  | HI
//  HI  LO  | LO
//  HI  HI  | Keep State
// 
func makeLatch(c *Circuit, a, b *Port) *Port {

  x := c.MakePort()
  q := c.Nand(a, x)
  y := c.Nand(q, b)
  c.Connect(y, x)

  return q
}

func TestNand(t *testing.T) {

  c := MakeCircuit()

  a := c.MakePort()
  b := c.MakePort()
  x := c.Nand(a, b)

  ps := []struct{
    InputA   int
    InputB   int
    Expected int
  }{
    { LO, LO, HI },
    { LO, HI, HI },
    { HI, LO, HI },
    { HI, HI, LO },
  }

  c.PreUpdate()

  for _, p := range ps {
    a.State = p.InputA
    b.State = p.InputB
    c.Update()
    if x.State != p.Expected {
      t.Errorf("Nand(%v,%v) expected=%v got=%v",
        p.InputA, p.InputB, p.Expected, x.State)
    }
  }
}

func TestLatch(t *testing.T) {

  c := MakeCircuit()

  a := c.MakePort()
  b := c.MakePort()

  q := makeLatch(c, a, b)

  ps := []struct{
    InputA   int
    InputB   int
    Expected int
  }{
    { LO, HI, HI },
    { HI, LO, LO },
    { HI, HI, LO },
    { LO, HI, HI },
    { HI, HI, HI },
  }

  c.PreUpdate()

  for _, p := range ps {
    a.State = p.InputA
    b.State = p.InputB
    c.Update()
    if q.State != p.Expected {
      t.Errorf("Latch(%v,%v) expected=%v got=%v",
        p.InputA, p.InputB, p.Expected, q.State)
    }
  }
}

// DLatch
//
// if enable is HI(1), then q equals d value.
// if enable is LO(0), then q keep prev value.
// 
//  d   e   | q
//  --------|----
//  LO  LO  | q (latched)
//  HI  LO  | q (latched)
//  LO  HI  | LO
//  HI  HI  | HI
// 
func TestDLatch(t *testing.T) {

  c := MakeCircuit()

  d := c.MakePort() // data
  e := c.MakePort() // enable

  f := c.Nand(d, d) // NOT(d)
  g := c.Nand(d, e)
  h := c.Nand(f, e)

  q := makeLatch(c, g, h)

  ps := []struct{
    InputD   int
    InputE   int
    Expected int
  }{
    { HI, HI, HI },
    { LO, HI, LO },
    { HI, LO, LO }, // d is prev value
    { LO, LO, LO }, // d is prev value
    { HI, LO, LO }, // d is prev value
    { LO, HI, LO },
    { HI, HI, HI },
    { HI, LO, HI }, // d is prev value
    { LO, LO, HI }, // d is prev value
    { HI, LO, HI }, // d is prev value
    { HI, HI, HI },
    { LO, HI, LO },
  }

  c.PreUpdate()

  for _, p := range ps {
    d.State = p.InputD
    e.State = p.InputE
    c.Update()

    if q.State != p.Expected {
      t.Errorf("DLatch(%v,%v) expected=%v got=%v",
        p.InputD, p.InputE, p.Expected, q.State)
    }
  }
}

func TestDFF(t *testing.T) {

  c := MakeCircuit()

  in := c.MakePort()
  e1 := c.Nand(c.clock, c.clock) // NOT(c.clock)
  f1 := c.Nand(in, in) // NOT(in)
  g1 := c.Nand(in, e1)
  h1 := c.Nand(f1, e1)
  q1 := makeLatch(c, g1, h1)

  f2 := c.Nand(q1, q1) // NOT(q1)
  g2 := c.Nand(q1, c.clock)
  h2 := c.Nand(f2, c.clock)
  q  := makeLatch(c, g2, h2)
  nq := c.Nand(q, q) // NOT(q)

  ps := []struct{
    Input      int
    ExpectedQ  int
    ExpectedNQ int
  }{
    { LO, LO, HI },
    { HI, HI, LO },
    { HI, HI, LO },
    { LO, LO, HI },
    { LO, LO, HI },
    { HI, HI, LO },
  }

  //initialize
  in.State = LO

  c.PreUpdate()

  for i, p := range ps {
    prevQ  := q.State
    prevNQ := nq.State

    in.State = p.Input
    c.ClockDown()
    c.Update() //keep q/nq state and store input state

    if q.State != prevQ {
      t.Errorf("DFF q.State after ClockDown[%v] expected=%v got=%v",
        i, prevQ, q.State)
    }
    if nq.State != prevNQ {
      t.Errorf("DFF nq.State after ClockDown[%v] expected=%v got=%v",
        i, prevNQ, nq.State)
    }

    in.State = HI //test changing state at this point has no effect

    c.ClockUp()
    c.Update() //update state

    if q.State != p.ExpectedQ {
      t.Errorf("DFF q.State after ClockUp[%v] expected=%v got=%v",
        i, p.ExpectedQ, q.State)
    }
    if nq.State != p.ExpectedNQ {
      t.Errorf("DFF nq.State after ClockUp[%v] expected=%v got=%v",
        i, p.ExpectedNQ, nq.State)
    }
  }
}

//  +--- n [ NOT ] <----+
//  |                   |
//  +----> [ DFF ] q ---+
//
func TestDFFLoop(t *testing.T) {

  c := MakeCircuit()

  in := c.MakePort()
  e1 := c.Nand(c.clock, c.clock) // NOT(c.clock)
  f1 := c.Nand(in, in) // NOT(in)
  g1 := c.Nand(in, e1)
  h1 := c.Nand(f1, e1)
  q1 := makeLatch(c, g1, h1)

  f2 := c.Nand(q1, q1) // NOT(q1)
  g2 := c.Nand(q1, c.clock)
  h2 := c.Nand(f2, c.clock)
  q  := makeLatch(c, g2, h2)

  n  := c.Nand(q, q) // NOT(q)

  c.Connect(n, in)
  in.State = LO

  c.PreUpdate()

  for i, p := range []int{ HI, LO, HI, LO, HI, } {

    c.Cycle() // clockup-down and Update
    if q.State != p {
      t.Errorf("DFFLoop q.State after Tick[%v] expected=%v got=%v",
        i, p, q.State)
    }
  }
}

