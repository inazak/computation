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


func TestIsZero0(t *testing.T) {
  code   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.x)" //(IsZero 0)
  expect := "^x.^y.x" //True

  t.Logf("code = " + code)
  result := ReadCompileAndRun(t, code, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

