package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_Xor_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //Xor(9900, 9900) -> 0012
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9900 0002
      NAND 0004 0005 0006
      CONNECT 9900 0004
      CONNECT 0003 0005
      NAND 0007 0008 0009
      CONNECT 9900 0007
      CONNECT 0003 0008
      NAND 0010 0011 0012
      CONNECT 0006 0010
      CONNECT 0009 0011
      //Xor(9901, 9900) -> 0024
      NAND 0013 0014 0015
      CONNECT 9901 0013
      CONNECT 9900 0014
      NAND 0016 0017 0018
      CONNECT 9901 0016
      CONNECT 0015 0017
      NAND 0019 0020 0021
      CONNECT 9900 0019
      CONNECT 0015 0020
      NAND 0022 0023 0024
      CONNECT 0018 0022
      CONNECT 0021 0023
      //Xor(9900, 9901) -> 0036
      NAND 0025 0026 0027
      CONNECT 9900 0025
      CONNECT 9901 0026
      NAND 0028 0029 0030
      CONNECT 9900 0028
      CONNECT 0027 0029
      NAND 0031 0032 0033
      CONNECT 9901 0031
      CONNECT 0027 0032
      NAND 0034 0035 0036
      CONNECT 0030 0034
      CONNECT 0033 0035
      //Xor(9901, 9901) -> 0048
      NAND 0037 0038 0039
      CONNECT 9901 0037
      CONNECT 9901 0038
      NAND 0040 0041 0042
      CONNECT 9901 0040
      CONNECT 0039 0041
      NAND 0043 0044 0045
      CONNECT 9901 0043
      CONNECT 0039 0044
      NAND 0046 0047 0048
      CONNECT 0042 0046
      CONNECT 0045 0047
    `,
    Expect: [][]Pair {
      {
        { "0012", control.LO },
        { "0024", control.HI },
        { "0036", control.HI },
        { "0048", control.LO },
      },
      {
        { "0012", control.LO },
        { "0024", control.HI },
        { "0036", control.HI },
        { "0048", control.LO },
      },
      {
        { "0012", control.LO },
        { "0024", control.HI },
        { "0036", control.HI },
        { "0048", control.LO },
      },
    },
  }

  DoTest(t, tc)
}

