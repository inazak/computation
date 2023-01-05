package test

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


func Test_Register4_01(t *testing.T) {

  tc := TestCase {
    Script: `
      LO 9900
      HI 9901
      CLOCK 9909
      //Register4(9901, 9901, 9901, 9901, 9909, 9901, 9900) -> 0067, 0161, 0255, 0349
      NAND 0002 0003 0004
      CONNECT 9909 0002
      CONNECT 9909 0003
      NAND 0005 0006 0007
      CONNECT 0001 0005
      CONNECT 0004 0006
      NAND 0008 0009 0010
      CONNECT 0001 0008
      CONNECT 0001 0009
      NAND 0011 0012 0013
      CONNECT 0010 0011
      CONNECT 0004 0012
      NAND 0014 0015 0016
      CONNECT 0007 0014
      CONNECT 0007 0015
      NAND 0017 0018 0019
      CONNECT 9900 0017
      CONNECT 9900 0018
      NAND 0020 0021 0022
      CONNECT 0016 0020
      CONNECT 0019 0021
      NAND 0023 0024 0025
      CONNECT 9900 0023
      CONNECT 9900 0024
      NAND 0026 0027 0028
      CONNECT 0013 0026
      CONNECT 0025 0027
      NAND 0029 0030 0031
      CONNECT 0028 0029
      CONNECT 0028 0030
      NAND 0032 0033 0034
      NAND 0035 0036 0037
      CONNECT 0034 0035
      CONNECT 0037 0033
      CONNECT 0022 0032
      CONNECT 0031 0036
      NAND 0038 0039 0040
      CONNECT 0034 0038
      CONNECT 9909 0039
      NAND 0041 0042 0043
      CONNECT 0034 0041
      CONNECT 0034 0042
      NAND 0044 0045 0046
      CONNECT 0043 0044
      CONNECT 9909 0045
      NAND 0047 0048 0049
      CONNECT 0040 0047
      CONNECT 0040 0048
      NAND 0050 0051 0052
      CONNECT 9900 0050
      CONNECT 9900 0051
      NAND 0053 0054 0055
      CONNECT 0049 0053
      CONNECT 0052 0054
      NAND 0056 0057 0058
      CONNECT 9900 0056
      CONNECT 9900 0057
      NAND 0059 0060 0061
      CONNECT 0046 0059
      CONNECT 0058 0060
      NAND 0062 0063 0064
      CONNECT 0061 0062
      CONNECT 0061 0063
      NAND 0065 0066 0067
      NAND 0068 0069 0070
      CONNECT 0067 0068
      CONNECT 0070 0066
      CONNECT 0055 0065
      CONNECT 0064 0069
      NAND 0071 0072 0073
      CONNECT 9901 0071
      CONNECT 9901 0072
      NAND 0074 0075 0076
      CONNECT 0067 0074
      CONNECT 0073 0075
      NAND 0077 0078 0079
      CONNECT 0076 0077
      CONNECT 0076 0078
      NAND 0080 0081 0082
      CONNECT 9901 0080
      CONNECT 9901 0081
      NAND 0083 0084 0085
      CONNECT 0082 0083
      CONNECT 0082 0084
      NAND 0086 0087 0088
      CONNECT 0079 0086
      CONNECT 0079 0087
      NAND 0089 0090 0091
      CONNECT 0085 0089
      CONNECT 0085 0090
      NAND 0092 0093 0094
      CONNECT 0088 0092
      CONNECT 0091 0093
      CONNECT 0094 0001
      NAND 0096 0097 0098
      CONNECT 9909 0096
      CONNECT 9909 0097
      NAND 0099 0100 0101
      CONNECT 0095 0099
      CONNECT 0098 0100
      NAND 0102 0103 0104
      CONNECT 0095 0102
      CONNECT 0095 0103
      NAND 0105 0106 0107
      CONNECT 0104 0105
      CONNECT 0098 0106
      NAND 0108 0109 0110
      CONNECT 0101 0108
      CONNECT 0101 0109
      NAND 0111 0112 0113
      CONNECT 9900 0111
      CONNECT 9900 0112
      NAND 0114 0115 0116
      CONNECT 0110 0114
      CONNECT 0113 0115
      NAND 0117 0118 0119
      CONNECT 9900 0117
      CONNECT 9900 0118
      NAND 0120 0121 0122
      CONNECT 0107 0120
      CONNECT 0119 0121
      NAND 0123 0124 0125
      CONNECT 0122 0123
      CONNECT 0122 0124
      NAND 0126 0127 0128
      NAND 0129 0130 0131
      CONNECT 0128 0129
      CONNECT 0131 0127
      CONNECT 0116 0126
      CONNECT 0125 0130
      NAND 0132 0133 0134
      CONNECT 0128 0132
      CONNECT 9909 0133
      NAND 0135 0136 0137
      CONNECT 0128 0135
      CONNECT 0128 0136
      NAND 0138 0139 0140
      CONNECT 0137 0138
      CONNECT 9909 0139
      NAND 0141 0142 0143
      CONNECT 0134 0141
      CONNECT 0134 0142
      NAND 0144 0145 0146
      CONNECT 9900 0144
      CONNECT 9900 0145
      NAND 0147 0148 0149
      CONNECT 0143 0147
      CONNECT 0146 0148
      NAND 0150 0151 0152
      CONNECT 9900 0150
      CONNECT 9900 0151
      NAND 0153 0154 0155
      CONNECT 0140 0153
      CONNECT 0152 0154
      NAND 0156 0157 0158
      CONNECT 0155 0156
      CONNECT 0155 0157
      NAND 0159 0160 0161
      NAND 0162 0163 0164
      CONNECT 0161 0162
      CONNECT 0164 0160
      CONNECT 0149 0159
      CONNECT 0158 0163
      NAND 0165 0166 0167
      CONNECT 9901 0165
      CONNECT 9901 0166
      NAND 0168 0169 0170
      CONNECT 0161 0168
      CONNECT 0167 0169
      NAND 0171 0172 0173
      CONNECT 0170 0171
      CONNECT 0170 0172
      NAND 0174 0175 0176
      CONNECT 9901 0174
      CONNECT 9901 0175
      NAND 0177 0178 0179
      CONNECT 0176 0177
      CONNECT 0176 0178
      NAND 0180 0181 0182
      CONNECT 0173 0180
      CONNECT 0173 0181
      NAND 0183 0184 0185
      CONNECT 0179 0183
      CONNECT 0179 0184
      NAND 0186 0187 0188
      CONNECT 0182 0186
      CONNECT 0185 0187
      CONNECT 0188 0095
      NAND 0190 0191 0192
      CONNECT 9909 0190
      CONNECT 9909 0191
      NAND 0193 0194 0195
      CONNECT 0189 0193
      CONNECT 0192 0194
      NAND 0196 0197 0198
      CONNECT 0189 0196
      CONNECT 0189 0197
      NAND 0199 0200 0201
      CONNECT 0198 0199
      CONNECT 0192 0200
      NAND 0202 0203 0204
      CONNECT 0195 0202
      CONNECT 0195 0203
      NAND 0205 0206 0207
      CONNECT 9900 0205
      CONNECT 9900 0206
      NAND 0208 0209 0210
      CONNECT 0204 0208
      CONNECT 0207 0209
      NAND 0211 0212 0213
      CONNECT 9900 0211
      CONNECT 9900 0212
      NAND 0214 0215 0216
      CONNECT 0201 0214
      CONNECT 0213 0215
      NAND 0217 0218 0219
      CONNECT 0216 0217
      CONNECT 0216 0218
      NAND 0220 0221 0222
      NAND 0223 0224 0225
      CONNECT 0222 0223
      CONNECT 0225 0221
      CONNECT 0210 0220
      CONNECT 0219 0224
      NAND 0226 0227 0228
      CONNECT 0222 0226
      CONNECT 9909 0227
      NAND 0229 0230 0231
      CONNECT 0222 0229
      CONNECT 0222 0230
      NAND 0232 0233 0234
      CONNECT 0231 0232
      CONNECT 9909 0233
      NAND 0235 0236 0237
      CONNECT 0228 0235
      CONNECT 0228 0236
      NAND 0238 0239 0240
      CONNECT 9900 0238
      CONNECT 9900 0239
      NAND 0241 0242 0243
      CONNECT 0237 0241
      CONNECT 0240 0242
      NAND 0244 0245 0246
      CONNECT 9900 0244
      CONNECT 9900 0245
      NAND 0247 0248 0249
      CONNECT 0234 0247
      CONNECT 0246 0248
      NAND 0250 0251 0252
      CONNECT 0249 0250
      CONNECT 0249 0251
      NAND 0253 0254 0255
      NAND 0256 0257 0258
      CONNECT 0255 0256
      CONNECT 0258 0254
      CONNECT 0243 0253
      CONNECT 0252 0257
      NAND 0259 0260 0261
      CONNECT 9901 0259
      CONNECT 9901 0260
      NAND 0262 0263 0264
      CONNECT 0255 0262
      CONNECT 0261 0263
      NAND 0265 0266 0267
      CONNECT 0264 0265
      CONNECT 0264 0266
      NAND 0268 0269 0270
      CONNECT 9901 0268
      CONNECT 9901 0269
      NAND 0271 0272 0273
      CONNECT 0270 0271
      CONNECT 0270 0272
      NAND 0274 0275 0276
      CONNECT 0267 0274
      CONNECT 0267 0275
      NAND 0277 0278 0279
      CONNECT 0273 0277
      CONNECT 0273 0278
      NAND 0280 0281 0282
      CONNECT 0276 0280
      CONNECT 0279 0281
      CONNECT 0282 0189
      NAND 0284 0285 0286
      CONNECT 9909 0284
      CONNECT 9909 0285
      NAND 0287 0288 0289
      CONNECT 0283 0287
      CONNECT 0286 0288
      NAND 0290 0291 0292
      CONNECT 0283 0290
      CONNECT 0283 0291
      NAND 0293 0294 0295
      CONNECT 0292 0293
      CONNECT 0286 0294
      NAND 0296 0297 0298
      CONNECT 0289 0296
      CONNECT 0289 0297
      NAND 0299 0300 0301
      CONNECT 9900 0299
      CONNECT 9900 0300
      NAND 0302 0303 0304
      CONNECT 0298 0302
      CONNECT 0301 0303
      NAND 0305 0306 0307
      CONNECT 9900 0305
      CONNECT 9900 0306
      NAND 0308 0309 0310
      CONNECT 0295 0308
      CONNECT 0307 0309
      NAND 0311 0312 0313
      CONNECT 0310 0311
      CONNECT 0310 0312
      NAND 0314 0315 0316
      NAND 0317 0318 0319
      CONNECT 0316 0317
      CONNECT 0319 0315
      CONNECT 0304 0314
      CONNECT 0313 0318
      NAND 0320 0321 0322
      CONNECT 0316 0320
      CONNECT 9909 0321
      NAND 0323 0324 0325
      CONNECT 0316 0323
      CONNECT 0316 0324
      NAND 0326 0327 0328
      CONNECT 0325 0326
      CONNECT 9909 0327
      NAND 0329 0330 0331
      CONNECT 0322 0329
      CONNECT 0322 0330
      NAND 0332 0333 0334
      CONNECT 9900 0332
      CONNECT 9900 0333
      NAND 0335 0336 0337
      CONNECT 0331 0335
      CONNECT 0334 0336
      NAND 0338 0339 0340
      CONNECT 9900 0338
      CONNECT 9900 0339
      NAND 0341 0342 0343
      CONNECT 0328 0341
      CONNECT 0340 0342
      NAND 0344 0345 0346
      CONNECT 0343 0344
      CONNECT 0343 0345
      NAND 0347 0348 0349
      NAND 0350 0351 0352
      CONNECT 0349 0350
      CONNECT 0352 0348
      CONNECT 0337 0347
      CONNECT 0346 0351
      NAND 0353 0354 0355
      CONNECT 9901 0353
      CONNECT 9901 0354
      NAND 0356 0357 0358
      CONNECT 0349 0356
      CONNECT 0355 0357
      NAND 0359 0360 0361
      CONNECT 0358 0359
      CONNECT 0358 0360
      NAND 0362 0363 0364
      CONNECT 9901 0362
      CONNECT 9901 0363
      NAND 0365 0366 0367
      CONNECT 0364 0365
      CONNECT 0364 0366
      NAND 0368 0369 0370
      CONNECT 0361 0368
      CONNECT 0361 0369
      NAND 0371 0372 0373
      CONNECT 0367 0371
      CONNECT 0367 0372
      NAND 0374 0375 0376
      CONNECT 0370 0374
      CONNECT 0373 0375
      CONNECT 0376 0283
    `,
    Expect: [][]Pair {
      {
      //Register4(9901, 9901, 9901, 9901, 9909, 9901, 9900) -> 0067, 0161, 0255, 0349 
        { "0067", control.HI },
        { "0161", control.HI },
        { "0255", control.HI },
        { "0349", control.HI },
      },
      {
      //Register4(9901, 9901, 9901, 9901, 9909, 9901, 9900) -> 0067, 0161, 0255, 0349 
        { "0067", control.HI },
        { "0161", control.HI },
        { "0255", control.HI },
        { "0349", control.HI },
      },
      {
      //Register4(9901, 9901, 9901, 9901, 9909, 9901, 9900) -> 0067, 0161, 0255, 0349 
        { "0067", control.HI },
        { "0161", control.HI },
        { "0255", control.HI },
        { "0349", control.HI },
      },
    },
  }

  DoTest(t, tc)
}

