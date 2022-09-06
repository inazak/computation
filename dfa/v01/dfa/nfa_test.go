package dfa

import (
  "testing"
)

func TestSimpleNFANode(t *testing.T) {

  //            +- b -+      +- c -+         
  //            |     |      |     |         
  //            v     |      v     |        
  // [s0] - a -> [s1] +- b -> [s2] +- c -> [s3*]
  //
  // * is accept

  s0 := MakeNFANode("s0")
  s1 := MakeNFANode("s1")
  s2 := MakeNFANode("s2")
  s3 := MakeNFANode("s3")

  s0.AddTransition('a', s1)
  s1.AddTransition('b', s1)
  s1.AddTransition('b', s2)
  s2.AddTransition('c', s2)
  s2.AddTransition('c', s3)
  s3.SetAcceptNode(true)

  //t.Logf("NFADump:\n%s", s0.Dump())

  ps := []struct{
    String   string
    Expected bool
  }{
    { String:   "abc",
      Expected: true, },
    { String:   "abbbccc",
      Expected: true, },
    { String:   "aabbcc",
      Expected: false, },
    { String:   "ababc",
      Expected: false, },
    { String:   "xyz",
      Expected: false, },
  }

  for _, p := range ps {
    result, _ := s0.Accept(p.String)
    if result != p.Expected {
      t.Errorf("string=%v expected=%v, got=%v", p.String, p.Expected, result)
    }
  }
}


func TestFreemoveNFANode(t *testing.T) {

  //            +- b -+      +- c -+         
  //            |     |      |     |         
  //            v     |      v     |        
  // [s0] - a -> [s1] +- b -> [s2] +- c -> [s3*]
  //             |                         ^
  //             |                         |
  //             +--- (free) --------------+
  //
  // * is accept

  s0 := MakeNFANode("s0")
  s1 := MakeNFANode("s1")
  s2 := MakeNFANode("s2")
  s3 := MakeNFANode("s3")

  s0.AddTransition('a', s1)
  s1.AddTransition('b', s1)
  s1.AddTransition('b', s2)
  s1.AddFreemove(s3)
  s2.AddTransition('c', s2)
  s2.AddTransition('c', s3)
  s3.SetAcceptNode(true)

  //t.Logf("NFADump:\n%s", s0.Dump())

  ps := []struct{
    String   string
    Expected bool
  }{
    { String:   "abc",
      Expected: true, },
    { String:   "a",
      Expected: true, },
    { String:   "ab",
      Expected: true, },
    { String:   "abbb",
      Expected: true, },
    { String:   "aa",
      Expected: false, },
    { String:   "ac",
      Expected: false, },
    { String:   "abbcc",
      Expected: true, },
    { String:   "xyz",
      Expected: false, },
  }

  for _, p := range ps {
    result, _ := s0.Accept(p.String)
    if result != p.Expected {
      t.Errorf("string=%v expected=%v, got=%v", p.String, p.Expected, result)
    }
  }
}




