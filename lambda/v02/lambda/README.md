# lambda

ラムダ計算に必要な最低限のSECDマシンの実験。


## 構成

まずラムダ式を表現するASTとして、Expressionを定義する。

- interface Expression
  - Symbol
  - Function
  - Application

今回は簡素化のために、下記もExpressionに含める。
これはASTとしては必要ではない。VMのスタック要素として
Expressionを使う実装にしたため、必要になった。

- interface Expression
  - Closure
  - Dump

この Closure と Dump については、Environment という型をメンバーに持つ。
これは文字列（この実装では関数適用時の仮引数にあたるSymbol）をキーにして、
Expression を値としたmapになっている。

```
type Environment map[string]Expression
```

つぎにExpressionをVMで動作する命令にコンパイルするために
命令そのものにあたるStatementを定義する。

- interface Statement
  - Fetch
  - Close
  - Apply
  - Return

最後にVMを構成する。次の通りとなる。
codeはコンパイルしたStatementのリストで、これを順に処理しながら、
stackとenvに結果を記録していく。

```
type VM struct {
  stack []Expression
  env   Environment
  code  []Statement
}
```

実際にラムダ式、コンパイル結果、動作を順に見てみる。


## Case.1

ラムダ式が単純にシンボル一つの場合。

```
expression = x
```

コンパイルすると、Fetchだけになる。

```
Fetch x
```

Fetchは引数の文字列をEnvironmentで検索し、合致するキーがあれば、
そのExpressionをVMのスタックにPushする。
合致するキーがなければ、その文字列をSymbolとしてスタックにPushする。

このコードを実行する直前の状態は次の通り。

```
vm.code:
  Fetch x

vm.env
  (empty)

vm.stack
  (empty)
```

先頭のFetchが実行されると、stackにSymbolがPushされる。

```
vm.code:
  (empty)

vm.env
  (empty)

vm.stack
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
expression = (x y)
```

コンパイルすると、FetchとApplyになる。最も左から処理する場合
最初が `Fetch x` になる

```
Fetch x
Fetch y
Apply
```

Applyはスタックから二つ取り出す。
Pushした順序と逆順なので、一つ目が引数、二つ目が関数。
関数にあたるExpressionがClosureであった場合だけ、
特別な処理をする。それ以外の場合は、単にApplicationでラップして
スタックにPushするだけ。今回は後者。

このコードを実行する直前の状態は次の通り。

```
vm.code:
  Fetch x
  Fetch y
  Apply

vm.env
  (empty)

vm.stack
  (empty)
```

二つのFetchが実行されると、stackにSymbolがPushされた状態になる。

```
vm.code:
  Apply

vm.env
  (empty)

vm.stack
  +---------------------------+
  |        Symbol(y)          |
  +---------------------------+
  |        Symbol(x)          |
  +---------------------------+
```

今回はスタックから二つPopして、二つ目（関数側）がClosureではないので、
単純にApplicationとして再度Pushする。

```
vm.code:
  (empty)

vm.env
  (empty)

vm.stack
  +---------------------------+
  |    Application(x y)       |
  +---------------------------+
```

codeが空になったので、stackの先頭であるApplicationを取り出し、結果として返す。

```
result = (x y)
```



## Case.3

ラムダ式が単純な関数定義の場合。

```
expression = ^x.x
```

コンパイルすると、Closureの生成になる。
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
vm.code:
  Close x, [Fetch x; Return]

vm.env
  (empty)

vm.stack
  (empty)
```

Closeが実行されると、stackにClosureがPushされた状態になる。

```
vm.code:
  (empty)

vm.env
  (empty)

vm.stack
  +------------------------------+
  | Closure x, [Fetch x; Return] |
  +------------------------------+
```

codeが空になったので、stackの先頭であるClosureを取り出し、結果として返す。

```
result = <closure>
```

この実験では手を入れないが、結果がClosureであれば
Functionでラップして返すようにすれば、`^x.x` という
自然な形にすることもできる。



## Case.4

ラムダ式が関数適用の場合。

```
expression = (^x.x y)
```

コンパイルすると、これまでの説明から次の通りとなる。

```
Close x, [Fetch x; Return]
Fetch y
Apply
```

このコードを実行する直前の状態は次の通り。

```
vm.code:
  Close x, [Fetch x; Return]
  Fetch y
  Apply

vm.env
  (empty)

vm.stack
  (empty)
```

Closeまで実行されると、stackにClosureとSymbolがPushされた状態になる。

```
vm.code:
  Apply

vm.env
  (empty)

vm.stack
  +------------------------------+
  |           Symbol(y)          |
  +------------------------------+
  | Closure x, [Fetch x; Return] |
  +------------------------------+
```

Applyはスタックから二つ取り出す。一つ目が引数、二つ目が関数。
関数にあたるExpressionがClosureなので、次の操作を行う。

1. 現在のvm.codeとvm.envの内容を、Dumpに詰め込んで、スタックにPushする
2. Closureの保持しているCodeとEnvで、vm.codeとvm.envに上書きする
3. Closureの仮引数を、スタックからPopした二つ目の引数に紐づけてvm.envに設定する

次のようになる。

```
vm.code:
  Fetch x  <-- Closure.Code
  Return   <-- Closure.Code

vm.env
  [x] = y  <-- Closure.Env & Closure.Arg = Symbol(y)

vm.stack
  +------------------------------+
  |      Dump(code,env)          |
  +------------------------------+
```

ここで次のコード `Fetch x` が実行されると、vm.env に
登録されている内容から `y` が返却され、スタックにPushされる。

```
vm.code:
  Return

vm.env
  [x] = y

vm.stack
  +------------------------------+
  |         Symbol(y)            |
  +------------------------------+
  |      Dump(code,env)          |
  +------------------------------+
```

次にReturnが実行される。スタックから二つPopする。
一つ目を処理結果として返すために、もう一度スタックにPushする。
二つ目にPopした値はDumpなので、このDumpに保持している値で
vm.code と vm.env を上書きする、つまり以前の状態に戻す。

Returnを実行した後が次の通り。

```
vm.code:
  (empty)

vm.env
  (empty)

vm.stack
  +------------------------------+
  |         Symbol(y)            |
  +------------------------------+
```

codeが空になったので、stackの先頭であるSymbolを取り出し、結果として返す。



## 関数内部の簡約の実装案

次の3ステップの追加で関数内部の簡約は実装できる。あまり美しくはない。

### 1. ExpressionをFunctionに変換する命令の追加

Statement に `Wrap` を追加する。これはArgという仮引数の文字を持っている。
スタックから一つPopして、その値を関数本体としたFunctionを生成し、
もう一度スタックにPushする。

### 2. vm.codeが空で、スタックトップがClosureの場合に、実行を継続する

vm.codeがゼロの場合はループを終了していたが、スタックトップが
Closureの場合だけ、下記で継続する。

1. 新規に命令WrapをCodeとし、現在のvm.envをEnvとしたDumpを、スタックにPushする
2. vm.env はそのまま
3. vm.code をClosureのBodyで上書き
4. この状態で実行を継続する

最終的に、ClosureのBodyの最後にあったReturnが実行されて、1. で
PushしたDumpが呼び戻され、Wrapが実行されて、Functionが結果になる

### 3. Applyの動作の変更

Applyの引数側（右側）がClosureの場合、関数本体を簡約する必要がある。
Apply自体を次のように書き換える。

1. 関数側（左側）がClosureの場合は、これまで通り
2. 関数側も引数側も（左も右も）Closureではない場合は、これまで通り
3. 引数側（右側）がClosureの場合は、

この 3. の場合の処理を下記とする。

1. 関数側（左側）の値はスタックにPushする
2. 現在のvm.codeの先頭に `Wrap` と `Apply` をこの順で追加する
3. 新規にvm.codeをCodeとし、現在のvm.envをEnvとしたDumpを、スタックにPushする
4. vm.env はそのまま
5. vm.code をClosureのBodyで上書き
6. この状態で実行を継続する
 
最終的に、ClosureのBodyの最後にあったReturnが実行されて、3. で
PushしたDumpが呼び戻され、Wrapが実行されてから、もう一度改めてApplyすることになる
 


