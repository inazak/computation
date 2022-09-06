package main

import (
  "fmt"
  . "github.com/inazak/computation/dfa/v01/dfa"
)

var OKNG map[bool]string
func init() {
  OKNG = make(map[bool]string)
  OKNG[true]  = "ok"
  OKNG[false] = "NG"
}

func main() {

  // REGEXP: a(b|c)*d+
  // NFA:
  //                 freemove
  //                +-------------+
  //                |             |             +- d -+
  //                |     +- b -> [s3] --+      |     |
  //                v     |              |      v     |
  // [s0] -+-- a -> [s1]--+              +- d -> [s5*]+
  //       |           ^  |              |
  //       |           |  +- c -> [s4] --+
  //       |           |          |      |
  //       |           +----------+      |
  //       |           freemove          |
  //       |                             |
  //       +-- a ---------------> [s2] --+
  //
  // * is accept

  s0 := MakeNFANode("s0")
  s1 := MakeNFANode("s1")
  s2 := MakeNFANode("s2")
  s3 := MakeNFANode("s3")
  s4 := MakeNFANode("s4")
  s5 := MakeNFANode("s5")

  s0.AddTransition('a', s1)
  s0.AddTransition('a', s2)
  s1.AddTransition('b', s3)
  s1.AddTransition('c', s4)
  s3.AddFreemove(s1)
  s4.AddFreemove(s1)
  s3.AddTransition('d', s5)
  s4.AddTransition('d', s5)
  s2.AddTransition('d', s5)
  s5.AddTransition('d', s5)
  s5.SetAcceptNode(true)

  dfa := s0.ToDFA()

  fmt.Printf("NFADump:\n%s", s0.Dump())
  fmt.Printf("---------------\n")
  fmt.Printf("DFADump:\n%s", dfa.Dump())
  fmt.Printf("---------------\n")

  ps := []struct{
    String   string
    Expected bool
  }{
    { String:   "a",
      Expected: false, },
    { String:   "b",
      Expected: false, },
    { String:   "c",
      Expected: false, },
    { String:   "d",
      Expected: false, },
    { String:   "aa",
      Expected: false, },
    { String:   "ad",
      Expected: true, },
    { String:   "adc",
      Expected: false, },
    { String:   "abd",
      Expected: true, },
    { String:   "acd",
      Expected: true, },
    { String:   "add",
      Expected: true, },
    { String:   "abbccdd",
      Expected: true, },
    { String:   "abbbcccddd",
      Expected: true, },
    { String:   "abcbcbd",
      Expected: true, },
    { String:   "abcbcbcd",
      Expected: true, },
    { String:   "abcbbccbcd",
      Expected: true, },
    { String:   "abcdcbd",
      Expected: false, },
  }

  fmt.Printf("Check Pattern:\n")
  for i, p := range ps {
    nfaresult, _ := s0.Accept(p.String)
    dfaresult, _ := dfa.Accept(p.String)

    fmt.Printf("[%02d] string=%v expected=%v => NFA:%v, DFA:%v\n",
      i, p.String, p.Expected,
      OKNG[nfaresult == p.Expected],
      OKNG[dfaresult == p.Expected])
  }
}

/* Result

NFADump:
s0  - a -> s2
s0  - a -> s1
s2  - d -> s5*
s1  - b -> s3
s1  - c -> s4
s5* - d -> s5*
s3  - d -> s5*
s3  -----> s1
s4  - d -> s5*
s4  -----> s1
---------------
DFADump:
s0  - a -> s1
s1  - c -> s2
s1  - b -> s3
s1  - d -> s4*
s2  - c -> s2
s2  - d -> s4*
s2  - b -> s3
s3  - c -> s2
s3  - d -> s4*
s3  - b -> s3
s4* - d -> s4*
---------------
Check Pattern:
[00] string=a expected=false => NFA:ok, DFA:ok
[01] string=b expected=false => NFA:ok, DFA:ok
[02] string=c expected=false => NFA:ok, DFA:ok
[03] string=d expected=false => NFA:ok, DFA:ok
[04] string=aa expected=false => NFA:ok, DFA:ok
[05] string=ad expected=true => NFA:ok, DFA:ok
[06] string=adc expected=false => NFA:ok, DFA:ok
[07] string=abd expected=true => NFA:ok, DFA:ok
[08] string=acd expected=true => NFA:ok, DFA:ok
[09] string=add expected=true => NFA:ok, DFA:ok
[10] string=abbccdd expected=true => NFA:ok, DFA:ok
[11] string=abbbcccddd expected=true => NFA:ok, DFA:ok
[12] string=abcbcbd expected=true => NFA:ok, DFA:ok
[13] string=abcbcbcd expected=true => NFA:ok, DFA:ok
[14] string=abcbbccbcd expected=true => NFA:ok, DFA:ok
[15] string=abcdcbd expected=false => NFA:ok, DFA:ok

*/

