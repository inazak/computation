# tm

Turing machine simulator

チューリングマシンのシミュレーション


## Example - simulation Turing machine

`example/main.go` で何をやっているかというと、
`1` と `0` で構成された文字列の回文チェック。

`_` は空白セルを表し、`[``]` で囲まれた部分がヘッド位置。下記のような状態からスタート。

```
________[1]001001__
```



### ルールに展開するときには次のようになる

初期状態を `s0` として、入力文字の左端にヘッドがある状態とする。
文字列が空白の場合は `_` を読むので、いきなり受理状態 `s7` に遷移して終わり。
`0` もしくは `1` を読んだ場合は、右に進んで、それぞれ別の状態 `s1` `s2` に遷移する。

```
  s0 := MakeTMNode("s0")
  s1 := MakeTMNode("s1")
  s2 := MakeTMNode("s2")
  s3 := MakeTMNode("s3")
  s4 := MakeTMNode("s4")
  s5 := MakeTMNode("s5")
  s6 := MakeTMNode("s6")
  s7 := MakeTMNode("s7")
  s8 := MakeTMNode("s8")

  s0.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s0.AddTranRule(TranRule{ Expect: '0', Write: '_', Move: RIGHT, Next: s1 })
  s0.AddTranRule(TranRule{ Expect: '1', Write: '_', Move: RIGHT, Next: s2 })
```

`s1` は `0` を読んだ状態ということを覚えたまま、右端まで進み、
右端の終端を超えたら（つまり `_` を読んだら）、左に切り替えして `s3` に遷移する。

```
  s1.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s3 })
  s1.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: RIGHT, Next: s1 })
  s1.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: RIGHT, Next: s1 })
```

`s2` は `1` を読んだ状態ということを覚えたまま、右端まで進み、
右端の終端を超えたら（つまり `_` を読んだら）、左に切り替えして `s4` に遷移する。

```
  s2.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s4 })
  s2.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: RIGHT, Next: s2 })
  s2.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: RIGHT, Next: s2 })
```

`s3` は左端で `0` を読んだ状態で、右端にたどり着いている。
右端が `_` であれば、それは一つ前の `s0` で消し込んだ場所になる。つまり
文字列が奇数だったということで、受理状態 `s7` に遷移。
右端が `0` であれば、左端に対応しているので、消し込んで左に進み、`s5` に遷移。
右端が `1` であれば、左端に対応していないので、受理できず終了となる `s8` に遷移。

```
  s3.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s3.AddTranRule(TranRule{ Expect: '0', Write: '_', Move: LEFT,  Next: s5 })
  s3.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s8 })
```

`s4` は左端で `1` を読んだ状態で、右端にたどり着いている。
右端が `_` であれば、それは一つ前の `s0` で消し込んだ場所になる。つまり
文字列が奇数だったということで、受理状態 `s7` に遷移。
右端が `0` であれば、左端に対応していないので、受理できず終了となる `s8` に遷移。
右端が `1` であれば、左端に対応しているので、消し込んで左に進み、`s5` に遷移。

```
  s4.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s4.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s8 })
  s4.AddTranRule(TranRule{ Expect: '1', Write: '_', Move: LEFT,  Next: s5 })
```

`s5` では `s3` と `s4` から左に一つ進んだ状態をチェックする。
文字が `_` である場合 `s0` で消し込んだ場所に戻ってきたことになる。
つまり残り文字が二文字の `00` か `11` だったということになる。受理状態 `s7` に遷移。
それ以外の場合は、左に進み、`s6` に遷移する。

```
  s5.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: LEFT,  Next: s7 })
  s5.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s6 })
  s5.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s6 })
```

`s6` では、左端である `_` にたどり着いたら、右に進み `s0` に戻るだけ。

```
  s6.AddTranRule(TranRule{ Expect: '_', Write: '_', Move: RIGHT, Next: s0 })
  s6.AddTranRule(TranRule{ Expect: '0', Write: '0', Move: LEFT,  Next: s6 })
  s6.AddTranRule(TranRule{ Expect: '1', Write: '1', Move: LEFT,  Next: s6 })
```

`s7` と `s8` には遷移するルールなし、受理状態の設定のみ。

```
  s7.SetAcceptNode(true)
  //s7(accept) has no rule
  //s8(reject) has no rule
```


### 実行すると次のようになる

これを `1001001` に対して実行する。下記の通り受理状態になる。

```
[   0] ________[1]001001__
[   1] ________[0]01001___
[   2] _______0[0]1001____
[   3] ______00[1]001_____
[   4] _____001[0]01______
[   5] ____0010[0]1_______
[   6] ___00100[1]________
[   7] __001001[_]________
[   8] ___00100[1]________
[   9] ____0010[0]________
[  10] _____001[0]0_______
[  11] ______00[1]00______
[  12] _______0[0]100_____
[  13] ________[0]0100____
[  14] ________[_]00100___
[  15] ________[0]0100____
[  16] ________[0]100_____
[  17] _______0[1]00______
[  18] ______01[0]0_______
[  19] _____010[0]________
[  20] ____0100[_]________
[  21] _____010[0]________
[  22] ______01[0]________
[  23] _______0[1]0_______
[  24] ________[0]10______
[  25] ________[_]010_____
[  26] ________[0]10______
[  27] ________[1]0_______
[  28] _______1[0]________
[  29] ______10[_]________
[  30] _______1[0]________
[  31] ________[1]________
[  32] ________[_]1_______
[  33] ________[1]________
[  34] ________[_]________
[  35] ________[_]________
[  36] ________[_]________
=> Acceptable
```


