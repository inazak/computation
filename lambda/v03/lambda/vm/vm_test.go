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

  vm     := NewVM(make(Environment), code)
  result := vm.Run()

  t.Logf("----- expression")
  t.Logf(expr.String())

  t.Logf("----- code")
  for _, c := range code { t.Logf(c.String()) }

  t.Logf("----- result")
  t.Logf(result.String())

  return result
}

func Test1(t *testing.T) {
  text   := "((^a.^b.(a b) x) y)"
  expect := "(x y)"

  result := ReadCompileAndRun(t, text)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}

func Test2(t *testing.T) {
  text   := "^a.(a ^x.(^y.y z))"
  expect := "^a.(a ^x.z)"

  result := ReadCompileAndRun(t, text)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}

