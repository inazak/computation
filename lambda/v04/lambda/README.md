# lambda

## キーワード

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











