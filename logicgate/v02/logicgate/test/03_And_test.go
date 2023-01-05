package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_And_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //And(9900, 9900) -> 0006
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9900 0002
      NAND 0004 0005 0006
      CONNECT 0003 0004
      CONNECT 0003 0005
      //And(9901, 9900) -> 0012
      NAND 0007 0008 0009
      CONNECT 9901 0007
      CONNECT 9900 0008
      NAND 0010 0011 0012
      CONNECT 0009 0010
      CONNECT 0009 0011
      //And(9900, 9901) -> 0018
      NAND 0013 0014 0015
      CONNECT 9900 0013
      CONNECT 9901 0014
      NAND 0016 0017 0018
      CONNECT 0015 0016
      CONNECT 0015 0017
      //And(9901, 9901) -> 0024
      NAND 0019 0020 0021
      CONNECT 9901 0019
      CONNECT 9901 0020
      NAND 0022 0023 0024
      CONNECT 0021 0022
      CONNECT 0021 0023
    `,
    Expect: [][]Pair {
      {
        { "0006", control.LO },
        { "0012", control.LO },
        { "0018", control.LO },
        { "0024", control.HI },
      },
      {
        { "0006", control.LO },
        { "0012", control.LO },
        { "0018", control.LO },
        { "0024", control.HI },
      },
      {
        { "0006", control.LO },
        { "0012", control.LO },
        { "0018", control.LO },
        { "0024", control.HI },
      },
    },
  }

  DoTest(t, tc)
}

