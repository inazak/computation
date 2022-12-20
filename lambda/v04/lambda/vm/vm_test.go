package vm

import (
  "testing"
  "github.com/inazak/computation/lambda/v04/lambda/reader"
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

func TestVMRunSucc(t *testing.T) {
  text   := "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))" //(Succ 1)
  expect := "^f.^x.(f (f x))" //2

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

func TestVMRunPred(t *testing.T) {
  text   := "(^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u) ^f.^x.(f (f x)))" //(Pred 2)
  expect := "^f.^x.(f x)" //1

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

func TestVMRunSub(t *testing.T) {
  text   := "((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) ^f.^x.(f (f x))) ^f.^x.(f x))" //((Sub 2) 1)
  expect := "^f.^x.(f x)" //1

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

func TestVMRunIsZeroT(t *testing.T) {
  text   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.x)" //(IsZero 0)
  expect := "^x.^y.x" //True

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

func TestVMRunIsZeroF(t *testing.T) {
  text   := "(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ^f.^x.(f x))" //(IsZero 1)
  expect := "^x.^y.y" //False

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

func TestVMMod(t *testing.T) {
  text   := "(((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^y.y)) ^x.^y.x) ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) x) y)) n) m)) ((f ((^x.^y.((y ^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)) x) m) n)) n)) m)) ^f.^x.(f (f (f (f (f x)))))) ^f.^x.(f (f x)))" //((Mod 5) 2)
  expect := "^f.^x.(f x)" //1

  result := ReadCompileAndRun(t, text, false)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", expect, result.String())
  }
}

