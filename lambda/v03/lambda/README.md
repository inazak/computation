# lambda

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
 

