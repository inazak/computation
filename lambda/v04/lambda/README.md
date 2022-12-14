

## v04 キーワード

整理すると下記の通り。

| AST          | Code                | Stack Value                      |
| ----         | ----                | ----                             |
| Symbol       | Fetch               | Symbol                           |
| Application  | Call, Apply, Return | Callure {Env, Code}, Application |
| Function     | Close, Return       | Closure {Env, Code}, Function    |


## Case.1

ラムダ式が単純にシンボル一つの場合。

```
x
```

ASTは次の通り。

```
Symbol { Name: "x" }
```

コンパイルすると、Fetchだけになる。

```
Fetch x
```

このコードを実行する直前の状態は次の通り。

```
code:
  Fetch x

env:
  (empty)

stack:
  (empty)
```

先頭のFetchが実行されると、スタックにSymbolがPushされる。

```
code:
  (empty)

env:
  (empty)

stack:
  +---------------------------+
  |        Symbol(x)          |
  +---------------------------+
```

codeが空になったので、スタックの先頭であるSymbolを取り出し、結果として返す。

```
x
```


## Case.2

ラムダ式が単純な適用の場合。

```
(x y)
```

ASTは次の通り。

```
Application { Left:  Symbol { Name: "x" },
              Right: Symbol { Name: "y" }, }
```

コンパイルすると、Callになる。
```
Call [ Fetch x; Fetch y; Apply; Return ]
```

このコードを実行する直前の状態は次の通り。

```
code:
  Call [ Fetch x; Fetch y; Apply; Return ]

env:
  (empty)

stack:
  (empty)
```

Callが実行されると、現在のenvをキャプチャした
CallureがスタックにPushされる。

```
code:
  (empty)

env:
  (empty)

stack:
  +---------------------------+
  |    Callure [env,code]     |
  +---------------------------+
```

codeが空になって、かつスタックの先頭がCallureなので
スタックからCallureを取り出して、次の操作を行う。

1. 現在のenvとcodeの内容を、Dumpに詰め込んで、スタックにPushする
2. envとcodeをCallureの保持している内容で上書きする


次のようになる。

```
code:
  Fetch x   <-- Callure.Code で上書きした
  Fetch y   <-- Callure.Code で上書きした
  Apply     <-- Callure.Code で上書きした
  Return    <-- Callure.Code で上書きした

env:
  (empty)   <-- Callure.Env で上書き

stack:
  +------------------------------+
  |      Dump(code,env)          |
  +------------------------------+
```

二つのFetchが実行されると、スタックにSymbolがPushされた状態になる。

```
code:
  Apply
  Return

env:
  (empty)

stack:
  +------------------------------+
  |          Symbol(y)           |
  +------------------------------+
  |          Symbol(x)           |
  +------------------------------+
```

Applyはスタックから二つPopする。二つ目（関数側）がClosureではないので、
単純にApplicationとして再度Pushする。

```
code:
  Return

env:
  (empty)

stack:
  +------------------------------+
  |     Application(x y)         |
  +------------------------------+
  |      Dump(code,env)          |
  +------------------------------+
```

次にReturnが実行される。スタックから二つPopする。
一つ目を処理結果として返すために、もう一度スタックにPushする。
二つ目にPopした値はDumpなので、このDumpに保持している値で
envとcodeを上書きする、つまり以前の状態に戻す。

Returnを実行した後が次の通り。

```
code:
  (empty)   <-- Dump.Code で上書きした

env:
  (empty)   <-- Dump.Env で上書きした

stack:
  +------------------------------+
  |     Application(x y)         |
  +------------------------------+
```

codeが空になったので、スタックの先頭である
Applicationを取り出し、結果として返す。


## Case.3

ラムダ抽象の場合。

```
^x.x
```

ASTは次の通り。

```
Function { Arg:  "x", Body: Symbol { Name: "x" } }
```

コンパイルすると、Closeになる。
パラメータとして、ラムダ変数の文字と、本体部分をコンパイルした値を持つ。
さらに、本体をコンパイルしたコード部分には、末尾に Return を付加する。
今回の場合、次の通り。

```
Close x, [Fetch x; Return]
```

コンパイルされたCloseが実行されると、そのラムダ変数と本体、
実行時のenvをまとめたClosureを作って、スタックにPushする。

このコードを実行する直前の状態は次の通り。

```
code:
  Close x, [Fetch x; Return]

env:
  (empty)

stack:
  (empty)
```

Closeが実行されると、スタックにClosureがPushされた状態になる。

```
code:
  (empty)

env:
  (empty)

stack:
  +------------------------------+
  | Closure x, [Fetch x; Return] |
  +------------------------------+
```

codeが空になったので、スタックの先頭であるClosureを取り出し、結果として返す。

```
code:
  (empty)

env:
  (empty)

stack:
  (empty)

result:
  Closure x, [Fetch x; Return]
```

しかしClosureはまだ内部に簡約できる可能性が残っている。そのため次の操作を行う。

1. ClosureのArgを、スタックにSymbolとしてPushする
2. codeにWrapを挿入する
3. この状態でcodeとenvの内容を、Dumpに詰め込んで、スタックにPushする
4. envとcodeを、Closureの保持している内容で上書きする
5. ClosureのArgを、envから削除する

まず 1. と 2. を実行すると次のようになる。

```
code:
  Wrap

env:
  (empty)

stack:
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+

result:
  Closure x, [Fetch x; Return]
```

次に 3. のDumpを実行する。

```
code:
  (empty)

env:
  (empty)

stack:
  +------------------------------+
  | Dump(env=empty, code=[Wrap]) |
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+

result:
  Closure x, [Fetch x; Return]
```

そして 4. でenvとcodeを上書きする。

```
code:
  Fetch x    <-- Closure.Code で上書きした
  Return     <-- Closure.Code で上書きした

env:
  ?=xxxx     <-- Closure.Code で上書きした

stack:
  +------------------------------+
  | Dump(env=empty, code=[Wrap]) |
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+

result:
  Closure x, [Fetch x; Return]
```

最後に 5. でClosureのArgと同じ名前のキーがあれば削除する。
なぜなら、この操作ではラムダ変数を無視して本体を評価しているため、
ラムダ変数が何かに展開されることはあってはならない。
言い換えると、ラムダ変数は束縛されていない状態で、本体を評価する。
この状態で実行を続ける。

```
code:
  Fetch x
  Return

env:
  (empty)

stack:
  +------------------------------+
  | Dump(env=empty, code=[Wrap]) |
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+
```

最初のFetchはxをenvから探すが、これは先の 5. で削除されているので見つからない。
SymbolがPushされて、Return待ちとなる。

```
code:
  Return

env:
  ?=xxxx

stack:
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+
  | Dump(env=empty, code=[Wrap]) |
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+
```

結果としてSymbol(x)がスタックトップに残り、codeとenvはDumpで書き戻しされる。

```
code:
  Wrap      <-- Dump.Code で上書きした

env:
  (empty)   <-- Dump.Env で上書きした

stack:
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+
```

codeにあるWrapはスタックの先頭を本体とし、
スタックの次のSymbolをラムダ変数とした
Functionを生成して、スタックにPushする。

```
code:
  (empty)

env:
  (empty)

stack:
  +-----------------------------------+
  | Function{Arg="x", Body=Symbol(x)} |
  +-----------------------------------+
```

最終的にFunctionの文字列表現として下記となる。

```
^x.x
```


## Case.4

ラムダ式が関数への適用の場合。

```
(^x.^y.x (x z))
```

ASTは次の通り。

```
Application { Left: Function { Arg:  "x", Body: 
                      Function { Arg:  "y", Body: 
                        Symbol { Name: "x" } } },
              Right: Application { Left: Symbol { Name: "x" },
                                   Right: Symbol { Name: "z" } } }
```

コンパイルすると、Call になる。

```
Call [ Close x, [ Close y, [ Fetch x ; Return] ; Return ] ;
       Call [ Fetch x ; Fetch z ; Apply ; Return ] ;
       Apply ;
       Return ]
```

まずCallが実行されてCallureがスタックにPushされる。
codeが空になるので、Callureの保持しているEnvとCodeが展開される。
ここまでは Case.2 と同じ。

```
code:
  Close x, [ Close y [ Fetch x ; Return ] ; Return ]   <-- Callure.Code で上書きした
  Call [ Fetch x ; Fetch z ; Apply ; Return ]          <-- Callure.Code で上書きした
  Apply                                                <-- Callure.Code で上書きした
  Return                                               <-- Callure.Code で上書きした

env:
  (empty)    <-- Callure.Env で上書きした

stack:
  +---------------------------------+
  | Dump env=(empty), code=(empty ) |
  +---------------------------------+
```

順次codeを実行していくと、Applyになる。
Callureは先に評価（展開）されることはなく、Callureという値のまま、関数に適用される。

```
code:
  Apply
  Return

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Callure [ Fetch x ; Fetch z ; Apply ; Return ]                      |
  +---------------------------------------------------------------------+
  | Closure x, env=(empty), code=[Close y, [Fetch x ; Return] ; Return] |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=(empty)                                      |
  +---------------------------------------------------------------------+
```

codeとenvはClosureの保持している値で上書きされる。
envはemptyで上書きされるが、Closureのラムダ変数 x に Callure が束縛されるので、
この組み合わせが保持される。


```
code:
  Close y, [Fetch x ; Return]
  Return

env:
  x = Callure [ Fetch x ; Fetch z ; Apply ; Return ]

stack:
  +---------------------------------------------------------------------+
  | Dump env=(empty), code= [Return]                                    |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=(empty)                                      |
  +---------------------------------------------------------------------+
```

CloseがClosureをPushして、Returnになる。

```
code:
  Return

env:
  x = Callure [ Fetch x ; Fetch z ; Apply ; Return ]

stack:
  +---------------------------------------------------------------------+
  | Closure y, env=x:<callure>, code=[Fetch x ; Return]                 |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code= [Return]                                    |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=(empty)                                      |
  +---------------------------------------------------------------------+
```

スタックからClosureとDumpをPopして、Dumpのenvとcodeを展開する。
スタックにはClosureを戻す。

```
code:
  Return   <-- Dump から上書き

env:
  (empty)  <-- Dump から上書き

stack:
  +---------------------------------------------------------------------+
  | Closure y, env=x:<callure>, code=[Fetch x ; Return]                 |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=(empty)                                      |
  +---------------------------------------------------------------------+
```

もう一度Returnになっているので、同じ処理を行う。

```
code:
  (empty)  <-- Dump から上書き

env:
  (empty)  <-- Dump から上書き

stack:
  +---------------------------------------------------------------------+
  | Closure y, env=x:<callure>, code=[Fetch x ; Return]                 |
  +---------------------------------------------------------------------+
```

コードがなくなったので、Case.3 と同じくClosureの内部の簡約になる。
スタックにはClosureのArgをSymbolとしてPushし、
codeにWrapを挿入してDumpをPushする。

```
code:
  Fetch x
  Return

env:
  x = Callure [ Fetch x ; Fetch z ; Apply ; Return ]

stack:
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

実行される Fetch でenvを参照し、スタックにCallureをPushする。

```
code:
  Return

env:
  x = Callure [ Fetch x ; Fetch z ; Apply ; Return ]

stack:
  +---------------------------------------------------------------------+
  | Callure [ Fetch x ; Fetch z ; Apply ; Return ]                      |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

Returnはスタックの先頭がCallureの場合、Callureの展開を行う。
この処理は少しややこしい。

0. codeがReturnで、スタックの先頭がCallureの場合
1. codeにReturnを挿入する（上の0.のReturnは消費してしまったため、代わりに追加）
2. この状態でcodeとenvの内容を、Dumpに詰め込んで、スタックにPushする
3. envとcodeを、Callureの保持している内容で上書きする

その結果、今度はenvに参照がない Fetch x となるので、単純にSymbolとして
スタックにPushする。Fetch z も同じ。

```
code:
  Fetch x
  Fetch z
  Apply
  Return

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Dump env=x:<callure> code=[Return]                                  |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

次のApplyは単純にApplicationを作成する。

```
code:
  Apply
  Return

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Symbol(z)                                                           |
  +---------------------------------------------------------------------+
  | Symbol(x)                                                           |
  +---------------------------------------------------------------------+
  | Dump env=x:<callure> code=[Return]                                  |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

codeの先頭がReturnになるので、一つ目のDumpを取り出す。

```
code:
  Return

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Application(x z)
  +---------------------------------------------------------------------+
  | Dump env=x:<callure> code=[Return]                                  |
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

envとcodeを上書きすると、またReturnが残る。次のDumpを取り出す。

```
code:
  Return

env:
  x = Callure [ Fetch x ; Fetch z ; Apply ; Return ]

stack:
  +---------------------------------------------------------------------+
  | Application(x z)
  +---------------------------------------------------------------------+
  | Dump env=(empty), code=[Wrap]                                       |
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

最後にWrapが残るので、Case.3 と同じ動作になる。

```
code:
  Wrap

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Application(x z)
  +---------------------------------------------------------------------+
  | Symbol(y)                                                           |
  +---------------------------------------------------------------------+
```

スタックの先頭を本体とし、スタックの次のSymbolをラムダ変数とした
Functionを生成して、スタックにPushする。

```
code:
  (empty)

env:
  (empty)

stack:
  +---------------------------------------------------------------------+
  | Function{ Arg=y, Body=(x z) }                                       |
  +---------------------------------------------------------------------+
```

これにより次の結果が得られる。

```
元のラムダ式
(^x.^y.x (x z))

結果
^y.(x z)
```


## この時点の問題

簡約後に残る文字列から、元の変数の束縛が分からない。
簡約後の文字列は「正確ではない」ということになる。

例を挙げる。下記の場合、yの下に振った数字の通り、y1 と y2 は
異なる変数であるが、

```
(^x.^y.(x y) y)
     2    2  1
```

簡約後に下記のようになった文字列を見ても判断ができない。

```
^y.(y y)
 2  1 2
```

回避する方法としては、変数はすべて関数のラムダ変数として束縛し、
何も束縛がない自由変数を許さない、ということになる。


