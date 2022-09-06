# npda

Nondeterministic Pushdown Automaton.

非決定性プッシュダウンオートマトンのシミュレーション


## Example - simulation NPDA

`example/main.go` で何をやっているかというと、
与えられた文字列が四則演算の構文通りかどうかをチェックするNPDAを作っている。

四則演算の構文は下記を利用している。数字は一桁の1,2,3のみ。
NPDAに展開する都合上、繰り返しを示す構文規則は使っていない

```
<expr>   ::= <term> | <term> ('+'|'-') <expr>
<term>   ::= <factor> | <factor> ('*'|'/') <term>
<factor> ::= <number> | '(' <expr> ')'
<number> :== 1,2,3
```

### ルールに展開するときには次のようになる

0. 遷移状態は3つのみで、最後のs2が受理状態となる。
`s0` はスタックの底として `'$'` をセットしておく。
```
s0 -> s1 -> s2
```


1. 非終端記号は、代替の大文字一文字に変換する。
```
 <expr> => E, <term> => T, ...
```


2. 初期状態 `s0` から次の状態 `s1` への遷移として、
スタックに非終端記号をPushするルールを追加する。

この場合、`s0` のスタック先頭が `'$'` であれば、
`'E' <expr>` をPushして `s1` に進むルールを追加する。

```
  s0.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s1,
                                    PushSymbol: []rune{'$', 'E'} })
```


3. スタックのトップに非終端記号の文字が現れたら、
対応する右辺に置き換えるようにスタックにプッシュするルールを追加する。
選択(|)がある場合は、単にルールを列挙する。

例えばスタックの先頭が `'N' <nuber>` であれば、
`'N'` をPopして削除し、代わりに `1` をPushしたインスタンスを作るルールとなる。
同様に、`2` をPushしたインスタンス、`3` をPushしたインスタンスを作るという
複数のルールが作成される。

```
  // <number> :== N
  s1.AddTranRuleOfStack( TranRule { StackTop: 'N',
                                    NextNode: s1,
                                    PushSymbol: []rune{'1'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'N',
                                    NextNode: s1,
                                    PushSymbol: []rune{'2'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'N',
                                    NextNode: s1,
                                    PushSymbol: []rune{'3'} })
```

例えばスタックの先頭が `'E' <expr>` であれば、
`'E'` をPopして削除し、代わりに `'T' <term>` をPushしたインスタンスを作るルール、
それから `'T' '+' 'E'` を逆順にPushしたインスタンス、
それから `'T' '-' 'E'` を逆順にPushしたインスタンスを作るルールが必要となる。

```
  // <expr>   ::= <term> | <term> ('+'|'-') <expr>
  s1.AddTranRuleOfStack( TranRule { StackTop: 'E',
                                    NextNode: s1,
                                    PushSymbol: []rune{'T'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'E',
                                    NextNode: s1,
                                    PushSymbol: []rune{'E', '+', 'T'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'E',
                                    NextNode: s1,
                                    PushSymbol: []rune{'E', '-', 'T'} })
```

スタックのトップに非終端記号があるインスタンスは、
繰り返しこの 3. のルールが適用され、スタックのトップが
全て終端記号になるまで続く。


4. スタックの先頭の文字と、入力文字が同じ場合は、
スタックの先頭文字を取り除く（Pop）するルールを追加する。

これにより、入力文字とスタック先頭の文字が合致したルールが選ばれ、
入力文字とスタック先頭の文字が消費されながら、次の遷移先を探すことなる。

```
  // rule for consuming the top of the stack
  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '+',
                                            StackTop: '+',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '-',
                                            StackTop: '-',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })
```


5. スタックの先頭文字が、スタックの底を示す `'$'` になっている場合は
`s2` の受理状態に遷移する

```
  // rule for moving to accept node
  s1.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s2,
                                    PushSymbol: []rune{'$'} })
```


### 実行すると次のようになる

これを `1*(2+3)-1/2` に対して実行する。

まず入力文字を読む前の段階で、ルールに従い Expr が展開され、
Expr から、Term、Factor、Number、が展開されていく。
最終的にスタックの先頭には終端文字だけが並ぶ状態の集合が出来上がる。
下記Nodeの`c0c0`は状態の`s1`にあたる。


```
Node(c0c0) Stack[1,$]
Node(c0c0) Stack[2,$]
Node(c0c0) Stack[3,$]
Node(c0c0) Stack[(,E,),$]
Node(c0c0) Stack[1,*,T,$]
Node(c0c0) Stack[2,*,T,$]
Node(c0c0) Stack[3,*,T,$]
Node(c0c0) Stack[(,E,),*,T,$]
Node(c0c0) Stack[1,/,T,$]
Node(c0c0) Stack[2,/,T,$]
Node(c0c0) Stack[3,/,T,$]
Node(c0c0) Stack[(,E,),/,T,$]
Node(c0c0) Stack[1,+,E,$]
Node(c0c0) Stack[2,+,E,$]
Node(c0c0) Stack[3,+,E,$]
Node(c0c0) Stack[(,E,),+,E,$]
Node(c0c0) Stack[1,*,T,+,E,$]
Node(c0c0) Stack[2,*,T,+,E,$]
Node(c0c0) Stack[3,*,T,+,E,$]
Node(c0c0) Stack[(,E,),*,T,+,E,$]
Node(c0c0) Stack[1,/,T,+,E,$]
Node(c0c0) Stack[2,/,T,+,E,$]
Node(c0c0) Stack[3,/,T,+,E,$]
Node(c0c0) Stack[(,E,),/,T,+,E,$]
Node(c0c0) Stack[1,-,E,$]
Node(c0c0) Stack[2,-,E,$]
Node(c0c0) Stack[3,-,E,$]
Node(c0c0) Stack[(,E,),-,E,$]
Node(c0c0) Stack[1,*,T,-,E,$]
Node(c0c0) Stack[2,*,T,-,E,$]
Node(c0c0) Stack[3,*,T,-,E,$]
Node(c0c0) Stack[(,E,),*,T,-,E,$]
Node(c0c0) Stack[1,/,T,-,E,$]
Node(c0c0) Stack[2,/,T,-,E,$]
Node(c0c0) Stack[3,/,T,-,E,$]
Node(c0c0) Stack[(,E,),/,T,-,E,$]
```

そして1文字目の`1`を読むと次の通りになる。
スタックの先頭に`1`があったもののみ選択され、スタックの`1`は消費される。
下記の残った状態の先頭が「あるべき次の入力文字」となる。
この状態でNodeの`c100`が存在している。これは受理状態である`s2`となっている。
数字の`1`のみの文字列であれば、正しい構文であるため。

```
Node(c100) Stack[$]
Node(c0c0) Stack[*,T,$]
Node(c0c0) Stack[/,T,$]
Node(c0c0) Stack[+,E,$]
Node(c0c0) Stack[*,T,+,E,$]
Node(c0c0) Stack[/,T,+,E,$]
Node(c0c0) Stack[-,E,$]
Node(c0c0) Stack[*,T,-,E,$]
Node(c0c0) Stack[/,T,-,E,$]
```

そして2文字目の`*`を読むと次の通りになる。
スタックの先頭に`*`があったもののみ選択され、スタックの`*`は消費される。
下記の残った状態の先頭が「あるべき次の入力文字」となる。
受理状態のNodeである`c100`は無くなる。

```
Node(c0c0) Stack[1,$]
Node(c0c0) Stack[2,$]
Node(c0c0) Stack[3,$]
Node(c0c0) Stack[(,E,),$]
Node(c0c0) Stack[1,*,T,$]
Node(c0c0) Stack[2,*,T,$]
Node(c0c0) Stack[3,*,T,$]
Node(c0c0) Stack[(,E,),*,T,$]
Node(c0c0) Stack[1,/,T,$]
Node(c0c0) Stack[2,/,T,$]
Node(c0c0) Stack[3,/,T,$]
Node(c0c0) Stack[(,E,),/,T,$]
Node(c0c0) Stack[1,+,E,$]
Node(c0c0) Stack[2,+,E,$]
Node(c0c0) Stack[3,+,E,$]
Node(c0c0) Stack[(,E,),+,E,$]
Node(c0c0) Stack[1,*,T,+,E,$]
Node(c0c0) Stack[2,*,T,+,E,$]
Node(c0c0) Stack[3,*,T,+,E,$]
Node(c0c0) Stack[(,E,),*,T,+,E,$]
Node(c0c0) Stack[1,/,T,+,E,$]
Node(c0c0) Stack[2,/,T,+,E,$]
Node(c0c0) Stack[3,/,T,+,E,$]
Node(c0c0) Stack[(,E,),/,T,+,E,$]
Node(c0c0) Stack[1,-,E,$]
Node(c0c0) Stack[2,-,E,$]
Node(c0c0) Stack[3,-,E,$]
Node(c0c0) Stack[(,E,),-,E,$]
Node(c0c0) Stack[1,*,T,-,E,$]
Node(c0c0) Stack[2,*,T,-,E,$]
Node(c0c0) Stack[3,*,T,-,E,$]
Node(c0c0) Stack[(,E,),*,T,-,E,$]
Node(c0c0) Stack[1,/,T,-,E,$]
Node(c0c0) Stack[2,/,T,-,E,$]
Node(c0c0) Stack[3,/,T,-,E,$]
Node(c0c0) Stack[(,E,),/,T,-,E,$]
```

この繰り返しを経て最終的に、最後の文字`2`を読み終わると、
下記の状態の集合となり、受理状態であることが分かる。

```
Node(c100) Stack[$]
Node(c0c0) Stack[*,T,$]
Node(c0c0) Stack[/,T,$]
Node(c0c0) Stack[+,E,$]
Node(c0c0) Stack[*,T,+,E,$]
Node(c0c0) Stack[/,T,+,E,$]
Node(c0c0) Stack[-,E,$]
Node(c0c0) Stack[*,T,-,E,$]
Node(c0c0) Stack[/,T,-,E,$]

>> Result: true
```

