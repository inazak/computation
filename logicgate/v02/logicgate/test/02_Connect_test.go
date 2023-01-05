package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_Connect_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      CONNECT 9901 1000
      CONNECT 9901 1001
      CONNECT 9900 1002
      CONNECT 9900 1003
      CONNECT 9901 1004
      CONNECT 9901 1005
      CONNECT 9900 1006
      CONNECT 9900 1007
    `,
    Expect: [][]Pair {
      {
        { "1000", control.HI },
        { "1001", control.HI },
        { "1002", control.LO },
        { "1003", control.LO },
        { "1004", control.HI },
        { "1005", control.HI },
        { "1006", control.LO },
        { "1007", control.LO },
      },
      {
        { "1000", control.HI },
        { "1001", control.HI },
        { "1002", control.LO },
        { "1003", control.LO },
        { "1004", control.HI },
        { "1005", control.HI },
        { "1006", control.LO },
        { "1007", control.LO },
      },
      {
        { "1000", control.HI },
        { "1001", control.HI },
        { "1002", control.LO },
        { "1003", control.LO },
        { "1004", control.HI },
        { "1005", control.HI },
        { "1006", control.LO },
        { "1007", control.LO },
      },
    },
  }

  DoTest(t, tc)
}

