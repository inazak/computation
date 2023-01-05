package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_Nand_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //Nand(9900, 9900) -> 0003
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9900 0002
      //Nand(9901, 9900) -> 0006
      NAND 0004 0005 0006
      CONNECT 9901 0004
      CONNECT 9900 0005
      //Nand(9900, 9901) -> 0009
      NAND 0007 0008 0009
      CONNECT 9900 0007
      CONNECT 9901 0008
      //Nand(9901, 9901) -> 0012
      NAND 0010 0011 0012
      CONNECT 9901 0010
      CONNECT 9901 0011
    `,
    Expect: [][]Pair {
      {
        { "0003", control.HI },
        { "0006", control.HI },
        { "0009", control.HI },
        { "0012", control.LO },
      },
      {
        { "0003", control.HI },
        { "0006", control.HI },
        { "0009", control.HI },
        { "0012", control.LO },
      },
      {
        { "0003", control.HI },
        { "0006", control.HI },
        { "0009", control.HI },
        { "0012", control.LO },
      },
    },
  }

  DoTest(t, tc)
}

