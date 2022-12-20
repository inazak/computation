

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

先頭のFetchが実行されると、stackにSymbolがPushされる。

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

codeが空になったので、stackの先頭であるSymbolを取り出し、結果として返す。

```
result = x
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

codeが空になって、かつstackの先頭がCallureなので
stackからCallureを取り出して、次の操作を行う。

1. 現在のcodeとenvの内容を、Dumpに詰め込んで、スタックにPushする
2. Callureの保持しているcodeとenvで、codeとenvに上書きする


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

二つのFetchが実行されると、stackにSymbolがPushされた状態になる。

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
code と env を上書きする、つまり以前の状態に戻す。

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

codeが空になったので、stackの先頭であるApplicationを取り出し、結果として返す。


## Case.3

ラムダ式が単純な関数定義の場合。

```
^x.x
```

ASTは次の通り。

```
Function { Arg:  "x", Body: Symbol { Name: "x" } }
```

コンパイルすると、Closeになる。
パラメータとして、仮引数の文字と、関数本体部分をコンパイルした値を持つ。
さらに、関数本体をコンパイルしたコード部分には、末尾に `Return` を付加する。
今回の場合、次の通り。

```
Close x, [Fetch x; Return]
```

コンパイルされたCloseが実行されると、そのパラメータと
実行時のEnvironmentをまとめて、Closureを作ってスタックにPushする。

このコードを実行する直前の状態は次の通り。

```
code:
  Close x, [Fetch x; Return]

env:
  (empty)

stack:
  (empty)
```

Closeが実行されると、stackにClosureがPushされた状態になる。

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

codeが空になったので、stackの先頭であるClosureを取り出し、結果として返す。

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
4. Closureの保持しているcodeとenvで、codeとenvに上書きする
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
なぜなら、この操作では仮引数を無視して本体を評価しているため。
仮引数が何かに展開されることはあってはならない。
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

結果としてSymbol(x)が返り、codeとenvはDumpで書き戻しされる。

```
code:
  Wrap      <-- Dump.Code で上書きした

env:
  (empty)   <-- Dump.Env で上書きした

stack:
  +------------------------------+
  |         Symbol(x)            |
  +------------------------------+

result:
  Symbol(x)
```

codeにあるWrapはスタックの先頭のSymbolをArg、取得済の結果をBodyとした
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



