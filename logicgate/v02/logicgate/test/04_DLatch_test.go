package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_DLatch_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //DLatch(9900, 9901) -> 0012
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9901 0002
      NAND 0004 0005 0006
      CONNECT 9900 0004
      CONNECT 9900 0005
      NAND 0007 0008 0009
      CONNECT 0006 0007
      CONNECT 9901 0008
      NAND 0010 0011 0012
      NAND 0013 0014 0015
      CONNECT 0012 0013
      CONNECT 0015 0011
      CONNECT 0003 0010
      CONNECT 0009 0014
      //DLatch(9901, 9901) -> 0027
      NAND 0016 0017 0018
      CONNECT 9901 0016
      CONNECT 9901 0017
      NAND 0019 0020 0021
      CONNECT 9901 0019
      CONNECT 9901 0020
      NAND 0022 0023 0024
      CONNECT 0021 0022
      CONNECT 9901 0023
      NAND 0025 0026 0027
      NAND 0028 0029 0030
      CONNECT 0027 0028
      CONNECT 0030 0026
      CONNECT 0018 0025
      CONNECT 0024 0029
    `,
    Expect: [][]Pair {
      {
        { "0012", control.LO },
        { "0027", control.HI },
      },
      {
        { "0012", control.LO },
        { "0027", control.HI },
      },
      {
        { "0012", control.LO },
        { "0027", control.HI },
      },
    },
  }

  DoTest(t, tc)
}

