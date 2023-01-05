package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_BDec2_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //BDec2(9900,9900) -> 0012, 0021, 0030, 0036
      NAND 0001 0002 0003
      CONNECT 9900 0001
      CONNECT 9900 0002
      NAND 0004 0005 0006
      CONNECT 9900 0004
      CONNECT 9900 0005
      NAND 0007 0008 0009
      CONNECT 0003 0007
      CONNECT 0006 0008
      NAND 0010 0011 0012
      CONNECT 0009 0010
      CONNECT 0009 0011
      NAND 0013 0014 0015
      CONNECT 9900 0013
      CONNECT 9900 0014
      NAND 0016 0017 0018
      CONNECT 9900 0016
      CONNECT 0015 0017
      NAND 0019 0020 0021
      CONNECT 0018 0019
      CONNECT 0018 0020
      NAND 0022 0023 0024
      CONNECT 9900 0022
      CONNECT 9900 0023
      NAND 0025 0026 0027
      CONNECT 0024 0025
      CONNECT 9900 0026
      NAND 0028 0029 0030
      CONNECT 0027 0028
      CONNECT 0027 0029
      NAND 0031 0032 0033
      CONNECT 9900 0031
      CONNECT 9900 0032
      NAND 0034 0035 0036
      CONNECT 0033 0034
      CONNECT 0033 0035
      //BDec2(9901,9900) -> 0048, 0057, 0066, 0072
      NAND 0037 0038 0039
      CONNECT 9901 0037
      CONNECT 9901 0038
      NAND 0040 0041 0042
      CONNECT 9900 0040
      CONNECT 9900 0041
      NAND 0043 0044 0045
      CONNECT 0039 0043
      CONNECT 0042 0044
      NAND 0046 0047 0048
      CONNECT 0045 0046
      CONNECT 0045 0047
      NAND 0049 0050 0051
      CONNECT 9900 0049
      CONNECT 9900 0050
      NAND 0052 0053 0054
      CONNECT 9901 0052
      CONNECT 0051 0053
      NAND 0055 0056 0057
      CONNECT 0054 0055
      CONNECT 0054 0056
      NAND 0058 0059 0060
      CONNECT 9901 0058
      CONNECT 9901 0059
      NAND 0061 0062 0063
      CONNECT 0060 0061
      CONNECT 9900 0062
      NAND 0064 0065 0066
      CONNECT 0063 0064
      CONNECT 0063 0065
      NAND 0067 0068 0069
      CONNECT 9901 0067
      CONNECT 9900 0068
      NAND 0070 0071 0072
      CONNECT 0069 0070
      CONNECT 0069 0071
      //BDec2(9900,9901) -> 0084, 0093, 0102, 0108
      NAND 0073 0074 0075
      CONNECT 9900 0073
      CONNECT 9900 0074
      NAND 0076 0077 0078
      CONNECT 9901 0076
      CONNECT 9901 0077
      NAND 0079 0080 0081
      CONNECT 0075 0079
      CONNECT 0078 0080
      NAND 0082 0083 0084
      CONNECT 0081 0082
      CONNECT 0081 0083
      NAND 0085 0086 0087
      CONNECT 9901 0085
      CONNECT 9901 0086
      NAND 0088 0089 0090
      CONNECT 9900 0088
      CONNECT 0087 0089
      NAND 0091 0092 0093
      CONNECT 0090 0091
      CONNECT 0090 0092
      NAND 0094 0095 0096
      CONNECT 9900 0094
      CONNECT 9900 0095
      NAND 0097 0098 0099
      CONNECT 0096 0097
      CONNECT 9901 0098
      NAND 0100 0101 0102
      CONNECT 0099 0100
      CONNECT 0099 0101
      NAND 0103 0104 0105
      CONNECT 9900 0103
      CONNECT 9901 0104
      NAND 0106 0107 0108
      CONNECT 0105 0106
      CONNECT 0105 0107
      //BDec2(9901,9901) -> 0120, 0129, 0138, 0144
      NAND 0109 0110 0111
      CONNECT 9901 0109
      CONNECT 9901 0110
      NAND 0112 0113 0114
      CONNECT 9901 0112
      CONNECT 9901 0113
      NAND 0115 0116 0117
      CONNECT 0111 0115
      CONNECT 0114 0116
      NAND 0118 0119 0120
      CONNECT 0117 0118
      CONNECT 0117 0119
      NAND 0121 0122 0123
      CONNECT 9901 0121
      CONNECT 9901 0122
      NAND 0124 0125 0126
      CONNECT 9901 0124
      CONNECT 0123 0125
      NAND 0127 0128 0129
      CONNECT 0126 0127
      CONNECT 0126 0128
      NAND 0130 0131 0132
      CONNECT 9901 0130
      CONNECT 9901 0131
      NAND 0133 0134 0135
      CONNECT 0132 0133
      CONNECT 9901 0134
      NAND 0136 0137 0138
      CONNECT 0135 0136
      CONNECT 0135 0137
      NAND 0139 0140 0141
      CONNECT 9901 0139
      CONNECT 9901 0140
      NAND 0142 0143 0144
      CONNECT 0141 0142
      CONNECT 0141 0143
    `,
    Expect: [][]Pair {
      {
        { "0012", control.HI },
        { "0021", control.LO },
        { "0030", control.LO },
        { "0036", control.LO },

        { "0048", control.LO },
        { "0057", control.HI },
        { "0066", control.LO },
        { "0072", control.LO },

        { "0084", control.LO },
        { "0093", control.LO },
        { "0102", control.HI },
        { "0108", control.LO },

        { "0120", control.LO },
        { "0129", control.LO },
        { "0138", control.LO },
        { "0144", control.HI },
      },
      {
        { "0012", control.HI },
        { "0021", control.LO },
        { "0030", control.LO },
        { "0036", control.LO },

        { "0048", control.LO },
        { "0057", control.HI },
        { "0066", control.LO },
        { "0072", control.LO },

        { "0084", control.LO },
        { "0093", control.LO },
        { "0102", control.HI },
        { "0108", control.LO },

        { "0120", control.LO },
        { "0129", control.LO },
        { "0138", control.LO },
        { "0144", control.HI },
      },
      {
        { "0012", control.HI },
        { "0021", control.LO },
        { "0030", control.LO },
        { "0036", control.LO },

        { "0048", control.LO },
        { "0057", control.HI },
        { "0066", control.LO },
        { "0072", control.LO },

        { "0084", control.LO },
        { "0093", control.LO },
        { "0102", control.HI },
        { "0108", control.LO },

        { "0120", control.LO },
        { "0129", control.LO },
        { "0138", control.LO },
        { "0144", control.HI },
      },
    },
  }

  DoTest(t, tc)
}

