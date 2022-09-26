package vm

import (
  "testing"
  "github.com/inazak/computation/lambda/v03/lambda/reader"
)

func ReadCompileAndRun(t *testing.T, text string) Value {
  l := reader.NewLexer(text)
  p := reader.NewParser(l)
  expr := p.Parse()
  code := Compile(expr)

  vm := NewVM(make(Environment), code)
  vm.EnableLogging()
  result := vm.Run()

  return result
}

func Test1(t *testing.T) {
  text   := "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))" //(Succ 1)
  expect := "^f.^x.(f (f x))" //2

  result := ReadCompileAndRun(t, text)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

