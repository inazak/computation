# dfa

DFA, NFA, and conversion of NFA to DFA.

決定性有限オートマトン（DFA）、非決定性有限オートマトン（NFA）、とNFAからDFAへの変換。


## Example - conversion of NFA to DFA

`example/main.go` で何をやっているかというと次の通り。

NFAをDFAに変換する考え方は単純。単一のNFAは同一のシンボルでも、複数の遷移先がある、というもの。では単一でなかったらどうなるか。つまり複数のNFAをまとめた集合の単位で考えた場合。集合のNFAはシンボルを取ったら、やっぱり複数の遷移先があるが、シンボルでグルーピングすることはできるようになる。

たとえば、「'a' を取ったら 1 に遷移するNFA」と、「'a' を取ったら 2 に遷移するNFA」を一つの集合に入れた場合、「'a' を取ったら [1, 2] のNFAの集合に遷移する」というように集約できる。シンボルが集約されるということは、シンボルと遷移先が1対1になるので、つまりDFAになる。


例えば`a(b|c)*d+`という正規表現をNFAにしてみる。とりあえず書いてみたら図の通り。実際のところ他にもたくさん表現方法はあるだろう。Freemove（Epsilon transition) を複数入れて試したかった。
![](https://raw.githubusercontent.com/inazak/computation/master/dfa/v01/dfa/misc/NFA.png)

NFAの実装は省略するが、こういう感じで遷移を記録する。
```
  s0 := MakeNFANode("s0")
  s1 := MakeNFANode("s1")
  s2 := MakeNFANode("s2")
  s3 := MakeNFANode("s3")
  s4 := MakeNFANode("s4")
  s5 := MakeNFANode("s5")

  s0.AddTransition('a', s1)
  s0.AddTransition('a', s2)
  s1.AddTransition('b', s3)
  s1.AddTransition('c', s4)
  s3.AddFreemove(s1)
  s4.AddFreemove(s1)
  s3.AddTransition('d', s5)
  s4.AddTransition('d', s5)
  s2.AddTransition('d', s5)
  s5.AddTransition('d', s5)
  s5.SetAcceptNode(true)
```

このNFAからDFAを生成するとどうなったか。
```
  dfa := s0.ToDFA()
```

このようになった、言われてみればその通り。変換するまでもないと言われている気分になる。
![](https://raw.githubusercontent.com/inazak/computation/master/dfa/v01/dfa/misc/DFA.png)

文字列のパターンを何通りか試してみる、NFAと変換後のDFAともに問題なし。
```
[00] string=a expected=false => NFA:ok, DFA:ok
[01] string=b expected=false => NFA:ok, DFA:ok
[02] string=c expected=false => NFA:ok, DFA:ok
[03] string=d expected=false => NFA:ok, DFA:ok
[04] string=aa expected=false => NFA:ok, DFA:ok
[05] string=ad expected=true => NFA:ok, DFA:ok
[06] string=adc expected=false => NFA:ok, DFA:ok
[07] string=abd expected=true => NFA:ok, DFA:ok
[08] string=acd expected=true => NFA:ok, DFA:ok
[09] string=add expected=true => NFA:ok, DFA:ok
[10] string=abbccdd expected=true => NFA:ok, DFA:ok
[11] string=abbbcccddd expected=true => NFA:ok, DFA:ok
[12] string=abcbcbd expected=true => NFA:ok, DFA:ok
[13] string=abcbcbcd expected=true => NFA:ok, DFA:ok
[14] string=abcbbccbcd expected=true => NFA:ok, DFA:ok
[15] string=abcdcbd expected=false => NFA:ok, DFA:ok
```


