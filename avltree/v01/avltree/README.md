# computation_avltree

AVL木の話。


## 木の回転（単回転）

例えば右回転するとしたら、

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig01.png)

切って

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig02.png)

ずらして

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig03.png)

つなぐイメージ。左回転は対称で考える。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig04.png)


## 木の回転（二重回転）

例えば左回転して右回転するとしたら、

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig11.png)

`A` と `B` の間で左回転のために、切って

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig12.png)

ずらして

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig13.png)

つなぐ。左回転が終わり。つぎに

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig14.png)

`B` と `P` の間で右回転のために、切って

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig15.png)

ずらして

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig16.png)

つなぐ。二重回転終わり。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig17.png)


## 回転した時の高さ（単回転）

単回転の右回転を考えてみる。回転によって根（root）からの深さ（根から葉までの高さ）が
変わったのは、赤色の部分木の部分。`X`は浅く（低く）、`Z`は深く（高く）なっている。

それぞれのノードが高さ（height）というパラメータを持っていたとして、
`X`や`Z`の部分木の「内部のノード」では高さは変わっていない。高さは葉がゼロで
親に向かって加算されるから。

```
if node.left == nil && node.right == nil
  node.height = 0
else
  node.height = max(node.left.height, node.right.height) + 1
```

そのため回転によって高さ（height）の再計算が必要なのは、子となる部分木が
変わってしまった`A`と`P`だけになる。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig21.png)


## 回転した時の高さ（二重回転）

二重回転の左回転-右回転を考えてみる。回転によって根（root）からの深さ（根から葉までの高さ）が
変わったのは、赤色の部分木の部分。`X`と`Y`は浅く（低く）、`Z`は深く（高く）なっている。

浅く（低く）なった`X`と`Y`が、どちらも元は`B`の子の部分木であることに注目。
二重回転が必要なのは、この位置にある部分木の高さを変える場合となる。

単回転と同じで`X`や`Y`や`Z`の部分木の「内部のノード」では高さは変わっていない。
高さは葉がゼロで親に向かって加算されるから。

回転によって高さ（height）の再計算が必要なのは、子となる部分木が
変わってしまった`A`と`B`と`P`になる。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig22.png)


## 削除するノードの左の子がない場合、または右の子がない場合

削除するノードに左の子がない場合は単に右の子を持ち上げるだけとなる。
削除後は再帰を戻りながら、高さの再計算を行う。
同じように右の子がない場合、という条件も追加する。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig31.png)


## 削除するノードの左の子の先の、右の子がない場合

次に考慮するのが「左の子はあるが、その先の右の子がない」場合。これは左の子を持ち上げればよい。
削除後は再帰を戻りながら、高さの再計算を行う。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig32.png)


## 削除するノードがそれ以外の場合

上記の2つ以外の場合は、「左の子の先にある、一番右の子を、削除ノードと入れ替える」という方法になる。
どうやって入れ替えるかというと、まず左の子の先にある、一番右の子を探すところから始める。
図の場合は `E` となる。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig33.png)

ここが肝心なところだが、見つけた E を削除するために、削除するノードの左の子をルートにして、
削除する関数をもう一つ呼び出す。つまり削除関数の中で、もう一つ別の削除関数を呼ぶことになる。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig34.png)

呼び出された `E` を削除する関数は、右の子が無いので、左の子を持ち上げて終わりとなる。
この時、削除関数が再帰的呼び出されているので、削除するノード `A` の左の子までは、
高さの再計算が再帰的に行われることになる。

そして削除された `E` を、本来削除したいノードである `A` と入れ替える。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig35.png)

最後に再帰がルートまで戻りながら、高さの再計算を行うことで削除が完了する。

![](https://raw.githubusercontent.com/inazak/computation/master/avltree/v01/avltree/misc/fig36.png)

