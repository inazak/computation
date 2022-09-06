package lambda

import (
  "testing"
)

func CompileAndRun(t *testing.T, expr Expression) Expression {
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
  test   := "x"
  expr   := Symbol{ Name: "x" }
  expect := "x"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}


func Test2(t *testing.T) {
  test   := "(x y)"
  expr   := Application { Left: Symbol{ Name: "x" }, Right: Symbol{ Name: "y" } }
  expect := "(x y)"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}


func Test3(t *testing.T) {
  test   := "^x.x"
  expr   := Function { Arg: "x", Body: Symbol{ Name: "x" } }
  expect := "<closure>"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}


func Test4(t *testing.T) {
  test   := "(^x.x y)"
  expr   := Application{ Left: Function { Arg: "x", Body: Symbol{ Name: "x" } },
                         Right: Symbol{ Name: "y" } }
  expect := "y"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}


func Test5(t *testing.T) {
  test := "((^a.^b.(a b) x) y)"
  expr := Application{
            Left: Application{
              Left: Function { Arg: "a", Body:
                Function{ Arg: "b", Body:
                  Application{ Left: Symbol{ Name: "a" },
                               Right: Symbol{ Name: "b" } } } },
              Right: Symbol{ Name: "x" } },
            Right: Symbol{ Name: "y" } }

  expect := "(x y)"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}


func Test6(t *testing.T) {
  test := "(x (^a.a y))"
  expr := Application{ Left: Symbol{ Name: "x" },
                       Right: Application{ Left: Function{ Arg: "a", Body:
                                             Symbol{ Name: "a"} },
                                           Right: Symbol{ Name: "y" } } }

  expect := "(x y)"

  if expr.String() != test {
    t.Fatalf("expr does not match test string")
  }

  result := CompileAndRun(t, expr)

  if result.String() != expect {
    t.Errorf("expected=%s, but got=%s", result.String(), expect)
  }
}

