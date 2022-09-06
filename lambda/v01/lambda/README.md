# lambda

untyped lambda calculus interpreter for fun.

FizzBuzz with lambda calculation is based on "Understanding Computation".
I tried from the implementation of lambda calculation,
but something seems to be useless, and the reduction is slow and there is no help for it.

I plan to rewrite it again.

ラムダ計算で FizzBuzz という話が 『アンダースタンディング コンピュテーション』の
「無からのプログラミング」で紹介されている。ラムダ計算の実装からやってみたが、
どこかが無駄なようで、簡約が遅くてどうしようもない。でもとりあえず動いた。

そのうち、もう一回書き直す予定。


## Demo

![](https://raw.githubusercontent.com/inazak/computation/master/lambda/v01/lambda/misc/demo1.gif)


## How to use

コマンドラインの起動オプションは次の通り。`-f FILEPATH` でファイルを指定しない場合は、
REPLで起動する。起動時に評価、出力のモードをいくつか選べる。

```
Usage: 
    lambda [OPTIONS]

Options:
    -d, -debug          ... run with debug print.
    -m, -mode MODENAME  ... select eval & output mode. default is 'reduce'.

      MODENAME: lambda  ... print lambda expression.
                expand  ... expand variable and print lambda expression.
                reduce  ... reduce expression and print.
                 index  ... reduce expression and print with symbol index.
                 ascii  ... reduce expression and print ASCII charactor.

    -f, -file FILEPATH  ... read file from FILEPATH, parse, eval, and print.
```

REPLでは簡約モード `reduce` になっているので、適当に入力すれば、簡約した結果を表示する。
終了するときは `:q` を入力する。

```
in REPL:
    :q, :quit          ... quit REPL.
    :l, :lambda        ... switch to lambda mode.
    :e, :expand        ... switch to expand mode.
    :i, :index         ... switch to index mode.
    :r, :reduce        ... switch to reduce mode.
    :a, :ascii         ... switch to ascii mode.
    :f, :file FILEPATH ... read file and parse, eval.
```

シンボルは英字小文字で入力する。lambda記号は `^` を使用する。例えば次の通り。
簡約戦略は最左最外で関数の内部も簡約する。

```
> (^x.^y.(x y) z)
^y.(z y)

> ^x.(^y.y x)
^x.x
```

数字はそのまま Church Number として展開される。例えば次の通り。

```
> 0
^f.^x.x

> 1
^f.^x.(f x)
```

式を `=` を使って英字大文字から始まる別名に紐付けることができる。例えば次の通り。

```
> A = ^x.(x x)

> (A z)
(z z)
```

定義済みの別名は次の通り。詳細は `ast/builtin.go` にある。

```
Succ, Pred, True, False, If, IsZero, Pair, Left, Right, Slide,
Empty, IsEmpty, First, Rest, Unshift, Push, Range, Fold, Map,
Y, Add, Sub, Div, Mod, Sum, LessOrEq, FizzBuzz,
```

例えば次の通り。

```
> ((Add 1) 2)
^f.^x.(f (f (f x)))
```

例えば定義済の `Sum` は不動点コンビネータ `Y` を使った総和を求める。
次のように、同じ名前で定義し直しもできてしまうが、`Pair` とか `Succ` とかを
再定義してしまうと、他が動かなくなってしまうので、やめておくこと。

```
> Sum = (Y ^f.^n.(((If (IsZero n)) 0) ((Add n) (f (Pred n)))))


> (Sum 3)
^f.^x.(f (f (f (f (f (f x))))))
```

モードを変更すると、出力を変更できる。例えば次の通り。
なお、モード `index` では、簡約した結果を De Bruijn Index で表示する。

```
> :lambda
> ((Add 1) 2)
((Add 1) 2)

> :expand
> ((Add 1) 2)
((^x.^y.((x ^n.^f.^x.(f ((n f) x))) y) ^f.^x.(f x)) ^f.^x.(f (f x)))

> :reduce
> ((Add 1) 2)
^f.^x.(f (f (f x)))

> :index
> ((Add 1) 2)
^.^.(2 (2 (2 1)))
```

モード `ascii` では、簡約した結果が Church Number もしくは Church Number の
リストである場合、ASCIIコードとして文字にして出力する。
また英字、数字については前に `'` をつけることで、その Church Numberに展開される。
例えば次の通り。

```
> :ascii
> ((Add 64) 1)
A

> 'A
A

> ((Unshift ((Unshift Empty) 'B)) 'A)
AB
```

オプション `-debug` を使って起動すると、展開と簡約のデバッグログが出力される。次の通り。

```
> (Succ 1)
[debug] begin Evaluator.reduce
[debug] expr=(Succ 1)
[debug] begin Evaluator.expand
[debug] expr=(Succ 1)
[debug] 1 times expand, expr=(^n.^f.^x.(f ((n f) x)) 1)
[debug] 2 times expand, expr=(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))
[debug] 3 times expand, expr=(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))
[debug] end Evaluator.expand
[debug] 1 times reduce, expr=^f.^x.(f ((^f.^x.(f x) f) x))
[debug] 2 times reduce, expr=^f.^x.(f (^x.(f x) x))
[debug] 3 times reduce, expr=^f.^x.(f (f x))
[debug] 4 times reduce, expr=^f.^x.(f (f x))
[debug] end Evaluator.reduce
^f.^x.(f (f x))
```


## Pair and List

ペアとリストの話。定義済の `Pair` は簡単な構造になっている。

```
> :expand
> Pair
^a.^b.^f.((f a) b)

> :reduce
> (Left ((Pair s) t))
s

> (Right ((Pair s) t))
t
```

毎回 Pair と書くのが面倒なので、`<` と `>` を使ってペアを表現する。
この記法は入れ子にできることを最後の例で示している。

```
> :reduce
> (Left <s t>)
s

> (Right <s t>)
t

> (Left (Left <<x y> x>))
x
```

ここからリストを作る。ペアの左が `True` の場合はリストの終端とする。
ペアの左が `False` の場合は、ペアの右にまたペアを作って、左が値、
右が残りのリストとする。ちょっと分かりにくい。

例を挙げると次のようになる。

```
<True True>                    => [] = Empty
<False <m Empty>>              => [m]
<False <m <False <n Empty>>>>  => [m, n]
```

空リストは定義済 `Empty` がある。
リストについては `[` と `]` を使った記法が用意されている。
次はペア記法でリストを書いた場合と、リスト記法を使った場合を交互に挙げている。

```
> :reduce
> <True True>
^f.((f ^x.^y.x) ^x.^y.x)

> []
^f.((f ^x.^y.x) ^x.^y.x)

> (First <False <m Empty>>)
m

> (First [m])
m

> (First (Rest <False <m <False <n Empty>>>>))
n

> (First (Rest [m n]))
n
```

ASCIIモードはリストの Church Number をASCIIコードとして印字する。
印字可能なASCIIコードの Church Number のリストを入力する記法として、
一般的なダブルクォートでの文字列入力が可能。
例えば次の通りとなる。

```
> :ascii
> ['A 'B 'C]
ABC

> "ABC"
ABC

> (Rest ['A 'B 'C])
BC

> (Rest "ABC")
BC
```

リストなので、例えば `Map` も使える。

```
> :ascii
> ((Map "ABC") (Add 1))
BCD
```


## FizzBuzz

note: some example maybe very slow. 
ここから先は実行速度がとても遅いことがある。

定義済の `DigitsToStr` を使うと、ASCIIモードで Church Number をそのまま印字できる。
例えば次の通り。

```
> :ascii
> (DigitToStr 3)
3
```

さて `FizzBuzz` はASCIIコードで改行を含む文字のリストを生成するように、
次のように定義している。ASCIIモードで評価することで、引数で与えられた数までの
いわゆる FizzBuzz を印字する。

```
FizzBuzz =
^x.((Map ((Range 1) x)) ^n.(((If (IsZero ((Mod n) 15))) "FizzBuzz")
                           (((If (IsZero ((Mod n) 3))) "Fizz")
                           (((If (IsZero ((Mod n) 5))) "Buzz")
                           ((Push (DigitToStr n)) 10)))))
```

試しに `15` まで実行してみるが、とても遅い。2010年の一般的なノートパソコンで
60秒かかる。

```
~~~ very slow ~~~

> :ascii
> (FizzBuzz 15)
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

その評価対象のlambda式はこの通り。

```
> :expand
> (FizzBuzz 15)
(^x.((^k.^f.((((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^l.^x.^g.(((^b.b (^p.(p ^a.^b.
a) l)) x) ((g (((f (^l.(^p.(p ^a.^b.b) (^p.(p ^a.^b.b) l)) l)) x) g)) (^l.(^p.(p
 ^a.^b.a) (^p.(p ^a.^b.b) l)) l)))) k) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x)) ^
l.^x.((^l.^x.((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) x) l)) l) (f x))
) (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x
.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y
)) m) n)) ((^l.^x.((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) x) l)) ((f 
(^n.^f.^x.(f ((n f) x)) m)) n)) m)) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x))) ^f.
^x.(f x)) x)) ^n.(((^b.b (^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) (((^f.(^x.(f (x x))
 ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x
.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y)) n) m)) ((f ((^x.^y
.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n)) m)) n) ^f.^x.(f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))) ((^a.^b.^f.((f a) b
) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f x)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
 ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))))))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
 ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))) ((^a.^b.^f.
((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f x)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))))))))))))))))))))))))))))))) ((^a.^b.^f.((f a) b) ^x
.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))
)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))) ((^a.^b.^f.((f a) b
) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) 
^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f 
(f (f (f (f x))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x))))))))))))))))))
)) (((^b.b (^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) (((^f.(^x.(f (x x)) ^x.(f (x x)))
 ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.
^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y)) n) m)) ((f ((^x.^y.((y ^n.^f.^x.
(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n)) m)) n) ^f.^x.(f (f (f x)))))) (
(^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))) ((^a.
^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (
f x))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x)))))))))))) (((^b.b (^x.((x
 (^x.^y.x ^x.^y.y)) ^x.^y.x) (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b 
((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (
g f))) ^u.x) ^u.u)) x) x) y)) n) m)) ((f ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f
))) ^u.x) ^u.u)) x) m) n)) n)) m)) n) ^f.^x.(f (f (f (f (f x)))))))) ((^a.^b.^f.
((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f x))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))) ((^a.^b
.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 x))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^
f.((f a) b) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) 
^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f 
(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))))))
))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
)))))))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) ^f.^x.(f (f 
(f (f (f (f (f (f (f (f x))))))))))) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x))))))
)))))) ((^l.^x.((((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^l.^x.^g.(((^b.b (^p.(p ^a.
^b.a) l)) x) ((g (((f (^l.(^p.(p ^a.^b.b) (^p.(p ^a.^b.b) l)) l)) x) g)) (^l.(^p
.(p ^a.^b.a) (^p.(p ^a.^b.b) l)) l)))) l) ((^l.^x.((^a.^b.^f.((f a) b) ^x.^y.y) 
((^a.^b.^f.((f a) b) x) l)) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x)) x)) ^l.^x.((
^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) x) l))) ((^f.(^x.(f (x x)) ^x.(
f (x x))) ^f.^n.((^l.^x.((((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^l.^x.^g.(((^b.b (
^p.(p ^a.^b.a) l)) x) ((g (((f (^l.(^p.(p ^a.^b.b) (^p.(p ^a.^b.b) l)) l)) x) g)
) (^l.(^p.(p ^a.^b.a) (^p.(p ^a.^b.b) l)) l)))) l) ((^l.^x.((^a.^b.^f.((f a) b) 
^x.^y.y) ((^a.^b.^f.((f a) b) x) l)) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x)) x))
 ^l.^x.((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) x) l))) (((^b.b ((^x.^
y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f)))
 ^u.x) ^u.u)) x) x) y)) n) ^f.^x.(f (f (f (f (f (f (f (f (f x))))))))))) ((^a.^b
.^f.((f a) b) ^x.^y.x) ^x.^y.x)) (f (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(
((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.
^h.(h (g f))) ^u.x) ^u.u)) x) x) y)) n) m)) (^n.^f.^x.(f ((n f) x)) ((f ((^x.^y.
((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n))) ^f.^x.x)) n) ^f.^
x.(f (f (f (f (f (f (f (f (f (f x)))))))))))))) ((^x.^y.((x ^n.^f.^x.(f ((n f) x
))) y) (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y
.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x)
 x) y)) n) m)) ((f ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m)
 n)) n)) m)) n) ^f.^x.(f (f (f (f (f (f (f (f (f (f x)))))))))))) ^f.^x.(f (f (f
 (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (
f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x)))))))))))))))))))))))
)))))))))))))))))))))))))))) n)) ^f.^x.(f (f (f (f (f (f (f (f (f (f x))))))))))
))))) ^f.^x.(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))))))))
```


