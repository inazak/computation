package parser

import (
  "testing"
  "github.com/inazak/computation/lambda/v01/lambda/lexer"
  "github.com/inazak/computation/lambda/v01/lambda/ast"
)

func TestParser1(t *testing.T) {

  text := "x"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Errorf("Statement type unexpected")
  }

  symbol, ok := exprstmt.Expr.(*ast.Symbol)
  if ! ok {
    t.Errorf("ExpressionStatement.Expr type unexpected")
  }

  if symbol.Name != 'x' {
    t.Errorf("Symbol name unexpected")
  }
}

func TestParser2(t *testing.T) {

  text := "(^x.y z)"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Errorf("Statement type unexpected")
  }

  appl, ok := exprstmt.Expr.(*ast.Application)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type is not application")
  }

  left, ok := appl.Left.(*ast.Function)
  if ! ok {
    t.Fatalf("Left is not Function")
  }

  right, ok := appl.Right.(*ast.Symbol)
  if ! ok {
    t.Fatalf("Right is not Symbol")
  }

  if left.Arg != 'x' {
    t.Fatalf("Left Function arg is not x")
  }

  body, ok := left.Body.(*ast.Symbol)
  if ! ok {
    t.Fatalf("Left Function body is not Symbol")
  }

  if body.Name != 'y' {
    t.Fatalf("Left Function body is not y")
  }

  if right.Name != 'z' {
    t.Fatalf("Right name is not z")
  }

  if exprstmt.String() != text {
    t.Errorf("stmt.String() expected=%v, but got=%v", text, exprstmt.String())
  }
}

func TestParser3(t *testing.T) {

  text := "Zoo = x"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  assistmt, ok := expr.(*ast.AssignmentStatement)
  if ! ok {
    t.Errorf("Statement type unexpected")
  }

  vari := assistmt.Var
  if vari.Name != "Zoo" {
    t.Errorf("Variable name unexpected")
  }

  symbol, ok := vari.Expr.(*ast.Symbol)
  if ! ok {
    t.Errorf("Variable.Expr type unexpected")
  }

  if symbol.Name != 'x' {
    t.Errorf("Symbol name unexpected")
  }
}

func TestParser4(t *testing.T) {

  text := "(Increment 99)"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  appl, ok := exprstmt.Expr.(*ast.Application)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  _, ok = appl.Left.(*ast.Variable)
  if ! ok {
    t.Fatalf("Left is not Variable")
  }

  _, ok = appl.Right.(*ast.Number)
  if ! ok {
    t.Fatalf("Right is not Number")
  }

  if expr.String() != "(Increment 99)" {
    t.Errorf("expr is not restored, got=%v", expr.String())
  }
}

func TestParser5(t *testing.T) {

  text := "(Succ 'A)"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  appl, ok := exprstmt.Expr.(*ast.Application)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  _, ok = appl.Left.(*ast.Variable)
  if ! ok {
    t.Fatalf("Left is not Variable")
  }

  _, ok = appl.Right.(*ast.Char)
  if ! ok {
    t.Fatalf("Right is not Char")
  }

  if expr.String() != "(Succ 'A)" {
    t.Errorf("expr is not restored, got=%v", expr.String())
  }
}

func TestParser6(t *testing.T) {

  text := "<Zoo x>"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  pair, ok := exprstmt.Expr.(*ast.Pair)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  first, ok := pair.Left.(*ast.Variable)
  if ! ok {
    t.Fatalf("Pair.Left is not Variable")
  }

  if first.Name != "Zoo" {
    t.Fatalf("Pair.Left is not Zoo")
  }

  second, ok := pair.Right.(*ast.Symbol)
  if ! ok {
    t.Fatalf("Pair.Right is not Symbol")
  }

  if second.Name != 'x' {
    t.Fatalf("Pair.Right is not x")
  }

  if expr.String() != "<Zoo x>" {
    t.Errorf("expr is not restored, got=%v", expr.String())
  }
}

func TestParser7(t *testing.T) {

  text := "[]"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  list, ok := exprstmt.Expr.(*ast.List)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  if len(list.Data) != 0 {
    t.Fatalf("Lst.Data is not zero")
  }

  if expr.String() != "[]" {
    t.Errorf("expr is not restored, got=%v", expr.String())
  }
}

func TestParser8(t *testing.T) {

  text := "[1 Zoo x]"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  list, ok := exprstmt.Expr.(*ast.List)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  if len(list.Data) != 3 {
    t.Fatalf("List.Data is not 3")
  }

  if _, ok := list.Data[0].(*ast.Number) ; ! ok {
    t.Errorf("List.Data[0] is not Number")
  }

  if _, ok := list.Data[1].(*ast.Variable) ; ! ok {
    t.Errorf("List.Data[1] is not Variable")
  }

  if _, ok := list.Data[2].(*ast.Symbol) ; ! ok {
    t.Errorf("List.Data[2] is not Symbol")
  }

  if expr.String() != "[1 Zoo x]" {
    t.Errorf("expr is not restored, got=%v", expr.String())
  }
}

func TestParser9(t *testing.T) {

  text := "(\"hello, world\" x)"

  p := NewParser(lexer.NewLexer(text))
  expr := p.Parse()
  err  := p.GetError()

  if err != nil {
    t.Fatalf("got parser error=%v",  err)
  }

  exprstmt, ok := expr.(*ast.ExpressionStatement)
  if ! ok {
    t.Fatalf("Statement type unexpected")
  }

  a, ok := exprstmt.Expr.(*ast.Application)
  if ! ok {
    t.Fatalf("ExpressionStatement.Expr type unexpected")
  }

  str, ok := a.Left.(*ast.Str)
  if ! ok {
    t.Fatalf("a.Left is not Str")
  }

  sym, ok := a.Right.(*ast.Symbol)
  if ! ok {
    t.Fatalf("a.Right is not Symbol")
  }

  if str.String() != "\"hello, world\"" {
    t.Errorf("str is not restored, got=%v", str.String())
  }

  if sym.Name != 'x' {
    t.Errorf("sym is not restored, got=%v", sym.String())
  }
}

