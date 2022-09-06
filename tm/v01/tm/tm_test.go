package tm

import (
  "testing"
)

func TestBinaryIncrement(t *testing.T) {

  s0 := MakeTMNode("s0")
  s1 := MakeTMNode("s1")
  s2 := MakeTMNode("s2")
  s2.SetAcceptNode(true)

  s0.AddTranRule(TranRule{ Expect: '0', Write: '1', Move: RIGHT, Next: s1 })
  s0.AddTranRule(TranRule{ Expect: '1', Write: '0', Move: LEFT,  Next: s0 })
  s0.AddTranRule(TranRule{ Expect: '_', Write: '1', Move: RIGHT, Next: s1 })

  s1.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s2 })

  tp := MakeTape([]rune{ '1','0','1' }, '1', []rune{})

  if tp.DumpTape(4) != "_101[1]____" {
    t.Errorf("MakeTape Expect=%v, got=%v", "_101[1]____", tp.DumpTape(4))
  }

  ok := s0.Run(tp)

  if ! ok {
    t.Errorf("TestBinaryIncrement is not Acceptable")
  }
  if tp.DumpTape(4) != "_110[0]____" {
    t.Errorf("UsedTape Expect=%v, got=%v", "_110[0]____", tp.DumpTape(4))
  }
}


func TestCountABC(t *testing.T) {

  s0 := MakeTMNode("s0")
  s1 := MakeTMNode("s1")
  s2 := MakeTMNode("s2")
  s3 := MakeTMNode("s3")
  s4 := MakeTMNode("s4")
  s5 := MakeTMNode("s5")
  s5.SetAcceptNode(true)

  s0.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s5 })
  s0.AddTranRule(TranRule{ Expect: 'X', Write: 'X', Move: RIGHT, Next: s0 })
  s0.AddTranRule(TranRule{ Expect: 'a', Write: 'X', Move: RIGHT, Next: s1 })

  s1.AddTranRule(TranRule{ Expect: 'X', Write: 'X', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: 'a', Write: 'a', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: 'b', Write: 'X', Move: RIGHT, Next: s2 })

  s2.AddTranRule(TranRule{ Expect: 'X', Write: 'X', Move: RIGHT, Next: s2 })
  s2.AddTranRule(TranRule{ Expect: 'b', Write: 'b', Move: RIGHT, Next: s2 })
  s2.AddTranRule(TranRule{ Expect: 'c', Write: 'X', Move: RIGHT, Next: s3 })

  s3.AddTranRule(TranRule{ Expect: 'c', Write: 'c', Move: RIGHT, Next: s3 })
  s3.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s4 })

  s4.AddTranRule(TranRule{ Expect: 'a', Write: 'a', Move: LEFT,  Next: s4 })
  s4.AddTranRule(TranRule{ Expect: 'b', Write: 'b', Move: LEFT,  Next: s4 })
  s4.AddTranRule(TranRule{ Expect: 'c', Write: 'c', Move: LEFT,  Next: s4 })
  s4.AddTranRule(TranRule{ Expect: 'X', Write: 'X', Move: LEFT,  Next: s4 })
  s4.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: RIGHT, Next: s0 })

  tp := MakeTape([]rune{}, 'a', []rune{'a','b','b','c','c'})

  if tp.DumpTape(6) != "______[a]abbcc_" {
    t.Errorf("MakeTape Expect=%v, got=%v", "______[a]abbcc_", tp.DumpTape(6))
  }

  ok := s0.Run(tp)

  if ! ok {
    t.Errorf("TestCountABC is not Acceptable")
  }
  if tp.DumpTape(6) != "_XXXXX[X]______" {
    t.Errorf("UsedTape Expect=%v, got=%v", "_XXXXX[X]______", tp.DumpTape(6))
  }
}


