package vm

import (
  "testing"
  "github.com/inazak/computation/lambda/v03/lambda/reader"
)

func ReadCompileAndRun(t *testing.T, text string, logging bool) (result Value) {
  l := reader.NewLexer(text)
  p := reader.NewParser(l)
  expr := p.Parse()
  code := Compile(expr)

  vm := NewVM(make(Environment), code)

  if logging {
    vm.EnableLogging()
    result = vm.Run()
    vm.DisableLogging()
  } else {
    result = vm.Run()
  }

  return result
}


func Test1(t *testing.T) {
  text   := "(^x.y z)"
  expect := "y"

  result := ReadCompileAndRun(t, text, true)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

