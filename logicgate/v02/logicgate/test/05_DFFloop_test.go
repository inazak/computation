package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_DFFLoop_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //DFF(1000, 9909) -> 0030 : NOT(0030) -> 0036 : 0036 -> 1000
      NAND 0001 0002 0003
      CONNECT 9909 0001
      CONNECT 9909 0002
      NAND 0004 0005 0006
      CONNECT 1000 0004
      CONNECT 0003 0005
      NAND 0007 0008 0009
      CONNECT 1000 0007
      CONNECT 1000 0008
      NAND 0010 0011 0012
      CONNECT 0009 0010
      CONNECT 0003 0011
      NAND 0013 0014 0015
      NAND 0016 0017 0018
      CONNECT 0015 0016
      CONNECT 0018 0014
      CONNECT 0006 0013
      CONNECT 0012 0017
      NAND 0019 0020 0021
      CONNECT 0015 0019
      CONNECT 9909 0020
      NAND 0022 0023 0024
      CONNECT 0015 0022
      CONNECT 0015 0023
      NAND 0025 0026 0027
      CONNECT 0024 0025
      CONNECT 9909 0026
      NAND 0028 0029 0030
      NAND 0031 0032 0033
      CONNECT 0030 0031
      CONNECT 0033 0029
      CONNECT 0021 0028
      CONNECT 0027 0032
      NAND 0034 0035 0036
      CONNECT 0030 0034
      CONNECT 0030 0035
      CONNECT 0036 1000
    `,
    Expect: [][]Pair {
      {
        { "1000", control.LO },
      },
      {
        { "1000", control.HI },
      },
      {
        { "1000", control.LO },
      },
    },
  }

  DoTest(t, tc)
}

