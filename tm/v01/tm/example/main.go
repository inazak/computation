package main

import (
  "fmt"
  . "github.com/inazak/computation/tm/v01/tm"
)

func main() {

  s0 := MakeTMNode("s0")
  s1 := MakeTMNode("s1")
  s2 := MakeTMNode("s2")
  s3 := MakeTMNode("s3")
  s4 := MakeTMNode("s4")
  s5 := MakeTMNode("s5")
  s6 := MakeTMNode("s6")
  s7 := MakeTMNode("s7")
  s8 := MakeTMNode("s8")

  s0.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s0.AddTranRule(TranRule{ Expect: '0', Write: '_', Move: RIGHT, Next: s1 })
  s0.AddTranRule(TranRule{ Expect: '1', Write: '_', Move: RIGHT, Next: s2 })

  // read '0' and seek right end
  s1.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s3 })
  s1.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: RIGHT, Next: s1 })

  // read '1' and seek right end
  s2.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s4 })
  s2.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: RIGHT, Next: s2 })
  s2.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: RIGHT, Next: s2 })

  // read '0' and reach right end and check cell is '0'
  s3.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s3.AddTranRule(TranRule{ Expect: '0', Write: '_', Move: LEFT,  Next: s5 })
  s3.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s8 })

  // read '1' and reach right end and check cell is '1'
  s4.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s4.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s8 })
  s4.AddTranRule(TranRule{ Expect: '1', Write: '_', Move: LEFT,  Next: s5 })

  // go to left
  s5.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s5.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s6 })
  s5.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s6 })

  // seek left end and check cell
  s6.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: RIGHT, Next: s0 })
  s6.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s6 })
  s6.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s6 })

  s7.SetAcceptNode(true)
  //s7(accept) has no rule
  //s8(reject) has no rule


  tp := MakeTape([]rune{}, '1', []rune{'0','0','1','0','0','1'})

  node  := s0
  count := 0
  fmt.Printf("[%4d] %s\n", count, tp.DumpTape(8))
  for node.HasNext(tp) {
    node = node.Step(tp)
    count += 1
    fmt.Printf("[%4d] %s\n", count, tp.DumpTape(8))
  }

  if node.IsAcceptNode() {
    fmt.Printf("=> Acceptable")
  } else {
    fmt.Printf("=> NOT Acceptable")
  }
}



