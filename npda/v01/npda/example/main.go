package main

import (
  "fmt"
  . "github.com/inazak/computation/npda/v01/npda"
)

func main() {

  // 与えられた文字列が四則演算の構文通りかどうかをチェックする
  //
  // 構文は下記を利用する。数字は一桁の1,2,3のみ。
  // NPDAに展開する都合上、繰り返しを示す構文規則は使っていない
  //
  // <expr>   ::= <term> | <term> ('+'|'-') <expr>
  // <term>   ::= <factor> | <factor> ('*'|'/') <term>
  // <factor> ::= <number> | '(' <expr> ')'
  // <number> :== 1,2,3
  //
  // ルールに展開するときには次のようになる
  //
  // 0) 遷移状態は3つのみで、最後のs2が受理状態となる
  //    s0 -> s1 -> s2
  //
  //    s0 はスタックの底として '$' をセットしておく
  //
  // 1) 非終端記号は、代替の大文字一文字に変換する
  //
  //    <expr> => E, <term> => T, ...
  //
  // 2) 初期状態（s0）から次の状態（s1）への遷移として、
  //    スタックに非終端記号をPushするルールを追加する
  //    この場合、s0のスタック先頭が '$' であれば
  //    'E' <expr> をPushして s1 に進むルールを追加する
  //
  // 3) スタックのトップに非終端記号の文字が現れたら、
  //    対応する右辺に置き換えるようにスタックにプッシュするルールを追加する
  //    選択(|)がある場合は、単にルールを列挙する
  //
  //    例えばスタックの先頭が 'N' <nuber> であれば、
  //    'N' をPopして削除し、代わりに 1 をPushしたインスタンスを作るルール
  //    同様に、2 をPushしたインスタンス、3 をPushしたインスタンスを作るという
  //    複数のルールが作成される
  //
  //    例えばスタックの先頭が 'E' <expr> であれば、
  //    'E' をPopして削除し、代わりに 'T' <term> をPushしたインスタンスを作るルール、
  //    それから 'T' '+' 'E' を逆順にPushしたインスタンス、
  //    それから 'T' '-' 'E' を逆順にPushしたインスタンスを作るというルールになる
  //
  //    スタックのトップに非終端記号があるインスタンスは
  //    繰り返しこの 3) のルールが適用され、スタックのトップが
  //    全て終端記号になるまで続く。
  //
  // 4) スタックの先頭の文字と、入力文字が同じ場合
  //    スタックの先頭文字を取り除く（Pop）するルールを追加する
  //
  //    これにより、入力文字とスタック先頭の文字が合致したルールが選ばれ、
  //    入力文字とスタック先頭の文字が消費されながら、次の遷移先を探すことなる
  //
  // 
  // 5) スタックの先頭文字が、スタックの底を示す '$' になっている場合は
  //    s2 の受理状態に遷移する

  s0 := MakeNPDANode()
  s1 := MakeNPDANode()
  s2 := MakeNPDANode()
  s2.SetAcceptNode(true)

  s0.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s1,
                                    PushSymbol: []rune{'$', 'E'} })

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

  // <term>   ::= <factor> | <factor> ('*'|'/') <term>
  s1.AddTranRuleOfStack( TranRule { StackTop: 'T',
                                    NextNode: s1,
                                    PushSymbol: []rune{'F'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'T',
                                    NextNode: s1,
                                    PushSymbol: []rune{'T', '*', 'F'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'T',
                                    NextNode: s1,
                                    PushSymbol: []rune{'T', '/', 'F'} })

  // <factor> ::= <number> | '(' <expr> ')'
  s1.AddTranRuleOfStack( TranRule { StackTop: 'F',
                                    NextNode: s1,
                                    PushSymbol: []rune{'N'} })

  s1.AddTranRuleOfStack( TranRule { StackTop: 'F',
                                    NextNode: s1,
                                    PushSymbol: []rune{')', 'E', '('} })

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


  // rule for consuming the top of the stack
  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '+',
                                            StackTop: '+',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '-',
                                            StackTop: '-',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '*',
                                            StackTop: '*',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '/',
                                            StackTop: '/',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '(',
                                            StackTop: '(',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   ')',
                                            StackTop: ')',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '1',
                                            StackTop: '1',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '2',
                                            StackTop: '2',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '3',
                                            StackTop: '3',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  // rule for moving to accept node
  s1.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s2,
                                    PushSymbol: []rune{'$'} })

  ps := []string {
    "1*(2+3)-1/2",
  }

  for _, p := range ps {
    fmt.Printf(">> Test String: %v\n", p)
    ok, _ := s0.AcceptWithDebugPrint(p, true)
    fmt.Printf("\n>> Result: %v\n", ok)
  }
}


