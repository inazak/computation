package npda

import (
  "testing"
)

func TestBracketBalance(t *testing.T) {

  //      {$} => {$}
  //    +---------------------+    (,{b} => {b,b}
  //    |                     |    ),{b} => {}
  //    |                     |  +----------------+
  //    V   (,{$} => {b,$}    |  |                |
  // [s0*] -----------------> [s1] <--------------+

  s0 := MakeNPDANode()
  s1 := MakeNPDANode()

  s0.SetAcceptNode(true)
  s0.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '(',
                                            StackTop: '$',
                                            NextNode: s1,
                                            PushSymbol: []rune{'$', 'b'} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   '(',
                                            StackTop: 'b',
                                            NextNode: s1,
                                            PushSymbol: []rune{'b', 'b'} })

  s1.AddTranRuleOfSymbolAndStack( TranRule{ Symbol:   ')',
                                            StackTop: 'b',
                                            NextNode: s1,
                                            PushSymbol: []rune{} })

  s1.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s0,
                                    PushSymbol: []rune{'$'} })

  ps := []struct {
    String string
    Expect bool
  }{
    { String: "()",
      Expect: true, },
    { String: "(",
      Expect: false, },
    { String: ")",
      Expect: false, },
    { String: ")(",
      Expect: false, },
    { String: "(()",
      Expect: false, },
    { String: "())",
      Expect: false, },
    { String: "()(",
      Expect: false, },
    { String: "(())",
      Expect: true, },
    { String: "()()",
      Expect: true, },
    { String: "(()(())())",
      Expect: true, },
    { String: "(()(())()()",
      Expect: false, },
  }

  for i, p := range ps {
    ok, _ := s0.Accept(p.String)
    if ok != p.Expect {
      t.Errorf("No.%v string=%v, expect=%v, got=%v\n", i, p.String, p.Expect, ok)
    }
  }
}


func TestParseFourOperation(t *testing.T) {

  // Four arithmetic operations
  // Numbers are only 1,2,3.
  //
  // <expr>   ::= <term> | <term> ('+'|'-') <expr>
  // <term>   ::= <factor> | <factor> ('*'|'/') <term>
  // <factor> ::= <number> | '(' <expr> ')'
  // <number> :== 1,2,3

  s0 := MakeNPDANode()
  s1 := MakeNPDANode()
  s2 := MakeNPDANode()
  s2.SetAcceptNode(true)

  // start
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


  s1.AddTranRuleOfStack( TranRule { StackTop: '$',
                                    NextNode: s2,
                                    PushSymbol: []rune{'$'} })

  ps := []struct {
    String string
    Expect bool
  }{
    { String: "1",
      Expect: true, },
    { String: "1+2",
      Expect: true, },
    { String: "1+2-3",
      Expect: true, },
    { String: "1++2",
      Expect: false, },
    { String: "(1+2",
      Expect: false, },
    { String: "(1+2)*3",
      Expect: true, },
  }

  for i, p := range ps {
    ok, _ := s0.Accept(p.String)
    if ok != p.Expect {
      t.Errorf("No.%v string=%v, expect=%v, got=%v\n", i, p.String, p.Expect, ok)
    }
  }
}

