package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_BDec_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //BDec(9900) -> 0003, 0004
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9900 0002
      CONNECT 9900 0004
      //BDec(9901) -> 0007, 0008
      NAND 0005 0006 0007
      CONNECT 9901 0005
      CONNECT 9901 0006
      CONNECT 9901 0008
    `,
    Expect: [][]Pair {
      {
        { "0003", control.HI },
        { "0004", control.LO },
        { "0007", control.LO },
        { "0008", control.HI },
      },
      {
        { "0003", control.HI },
        { "0004", control.LO },
        { "0007", control.LO },
        { "0008", control.HI },
      },
      {
        { "0003", control.HI },
        { "0004", control.LO },
        { "0007", control.LO },
        { "0008", control.HI },
      },
    },
  }

  DoTest(t, tc)
}

