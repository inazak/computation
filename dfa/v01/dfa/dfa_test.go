package dfa

import (
  "testing"
)

func TestSimpleDFANode(t *testing.T) {

  //            +- a -+      +- b -+         
  //            |     |      |     |         
  //            v     |      v     |        
  // [s0] - a -> [s1] +- b -> [s2] +- c -> [s3*]
  //
  // * is accept

  s0 := MakeDFANode("s0")
  s1 := MakeDFANode("s1")
  s2 := MakeDFANode("s2")
  s3 := MakeDFANode("s3")

  s0.AddTransition('a', s1)
  s1.AddTransition('a', s1)
  s1.AddTransition('b', s2)
  s2.AddTransition('b', s2)
  s2.AddTransition('c', s3)
  s3.SetAcceptNode(true)

  //t.Logf("DFADump:\n%s", s0.Dump())

  ps := []struct{
    String   string
    Expected bool
  }{
    { String:   "abc",
      Expected: true, },
    { String:   "a",
      Expected: false, },
    { String:   "aaabbbc",
      Expected: true, },
    { String:   "aabbcc",
      Expected: false, },
    { String:   "xyz",
      Expected: false, },
  }

  for _, p := range ps {
    result, _ := s0.Accept(p.String)
    if result != p.Expected {
      t.Errorf("expected=%v, got=%v", p.Expected, result)
    }
  }
}



