package lexer

import (
  "testing"
  "github.com/inazak/computation/lambda/v01/lambda/token"
)

func TestNextToken1(t *testing.T) {

  text := "Add = ^x.^y.((x Succ) y)"
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.VARIABLE, Literal: "Add" },
    {Type: token.EQUAL,    Literal: "=" },
    {Type: token.LAMBDA,   Literal: "^" },
    {Type: token.SYMBOL,   Literal: "x" },
    {Type: token.DOT,      Literal: "." },
    {Type: token.LAMBDA,   Literal: "^" },
    {Type: token.SYMBOL,   Literal: "y" },
    {Type: token.DOT,      Literal: "." },
    {Type: token.LPAREN,   Literal: "(" },
    {Type: token.LPAREN,   Literal: "(" },
    {Type: token.SYMBOL,   Literal: "x" },
    {Type: token.VARIABLE, Literal: "Succ" },
    {Type: token.RPAREN,   Literal: ")" },
    {Type: token.SYMBOL,   Literal: "y" },
    {Type: token.RPAREN,   Literal: ")" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}

func TestNextToken2(t *testing.T) {

  text := "<>[]"
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.LANGLE,   Literal: "<" },
    {Type: token.RANGLE,   Literal: ">" },
    {Type: token.LBRACKET, Literal: "[" },
    {Type: token.RBRACKET, Literal: "]" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}

func TestNextToken3(t *testing.T) {

  text := "5 15 100"
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.NUMBER,   Literal: "5" },
    {Type: token.NUMBER,   Literal: "15" },
    {Type: token.NUMBER,   Literal: "100" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}


func TestNextToken4(t *testing.T) {

  text := "'A 'z '3 '= '( '! '~ ''"
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.CHAR,     Literal: "A" },
    {Type: token.CHAR,     Literal: "z" },
    {Type: token.CHAR,     Literal: "3" },
    {Type: token.CHAR,     Literal: "=" },
    {Type: token.CHAR,     Literal: "(" },
    {Type: token.CHAR,     Literal: "!" },
    {Type: token.CHAR,     Literal: "~" },
    {Type: token.CHAR,     Literal: "'" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}

func TestNextToken5(t *testing.T) {

  text := " \"hello, world!\" "
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.STR,      Literal: "hello, world!" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}

func TestNextToken6(t *testing.T) {

  text := " \"hello, \\\"our\\\" world!\" "
  l := NewLexer(text)

  expected := []struct{
    Type    token.TokenType
    Literal string
  }{
    {Type: token.STR,      Literal: "hello, \"our\" world!" },
    {Type: token.EOL,      Literal: ""  },
  }

  for i, e := range expected {
    tk := l.NextToken()
    if tk.Type != e.Type {
      t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
    }
    if tk.Literal != e.Literal {
      t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
    }
  }
}

