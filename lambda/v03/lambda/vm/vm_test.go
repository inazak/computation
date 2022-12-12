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


func Test(t *testing.T) {
  //code   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.x)"
  //expect := "^x.^y.x" //->NG
  //code   := "((^f.^x.x (^x.^y.x ^x.^y.y)) ^x.^y.x)"
  //expect := "^x.^y.x" //->OK
  //code   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^a.z)"
  //expect := "(z ^x.^y.x)" //->OK
  //code   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^a.a)"
  //expect := "^x.^y.y" //->OK
  //code   := "(^x.(^x.^y.x ^x.^y.y) ^f.^x.x)"
  //expect := "^y.^x.^y.y" //->OK

  t.Logf("code = " + code)
  result := ReadCompileAndRun(t, code, true)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

