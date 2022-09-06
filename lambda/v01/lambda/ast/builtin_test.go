package ast

import (
  "testing"
)

func GetBuiltinVariable(name string) *Variable {
  for _, b := range BUILTINS {
    if b.Name == name {
      return &Variable{ Name: name, Expr: b.Expr }
    }
  }
  panic("GetBuiltinVariable: target not found")
}

func TestTrue(t *testing.T) {
  a := GetBuiltinVariable("True")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.^y.x" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestFalse(t *testing.T) {
  a := GetBuiltinVariable("False")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.^y.y" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestIf(t *testing.T) {
  a := GetBuiltinVariable("If")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^b.b" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestIsZero(t *testing.T) {
  a := GetBuiltinVariable("IsZero")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.((x (True False)) True)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestLessOrEq(t *testing.T) {
  a := GetBuiltinVariable("LessOrEq")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.^y.(IsZero ((Sub x) y))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestSucc(t *testing.T) {
  a := GetBuiltinVariable("Succ")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^n.^f.^x.(f ((n f) x))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestPred(t *testing.T) {
  a := GetBuiltinVariable("Pred")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^n.^f.^x.(((n ^g.^h.(h (g f))) ^u.x) ^u.u)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestAdd(t *testing.T) {
  a := GetBuiltinVariable("Add")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.^y.((x Succ) y)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestSub(t *testing.T) {
  a := GetBuiltinVariable("Sub")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^x.^y.((y Pred) x)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestDiv(t *testing.T) {
  a := GetBuiltinVariable("Div")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "(Y ^f.^m.^n.(((If ((LessOrEq n) m)) (Succ ((f ((Sub m) n)) n))) 0))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestMod(t *testing.T) {
  a := GetBuiltinVariable("Mod")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "(Y ^f.^m.^n.(((If ((LessOrEq n) m)) ((f ((Sub m) n)) n)) m))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestPair(t *testing.T) {
  a := GetBuiltinVariable("Pair")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^a.^b.^f.((f a) b)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestLeft(t *testing.T) {
  a := GetBuiltinVariable("Left")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^p.(p ^a.^b.a)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestRight(t *testing.T) {
  a := GetBuiltinVariable("Right")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^p.(p ^a.^b.b)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestSlide(t *testing.T) {
  a := GetBuiltinVariable("Slide")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^p.<(Right p) (Succ (Right p))>" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestEmpty(t *testing.T) {
  a := GetBuiltinVariable("Empty")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "<True True>" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestFirst(t *testing.T) {
  a := GetBuiltinVariable("First")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^l.(Left (Right l))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestRest(t *testing.T) {
  a := GetBuiltinVariable("Rest")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^l.(Right (Right l))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestUnshift(t *testing.T) {
  a := GetBuiltinVariable("Unshift")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^l.^x.<False <x l>>" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestRange(t *testing.T) {
  a := GetBuiltinVariable("Range")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "(Y ^f.^m.^n.(((If ((LessOrEq m) n)) ((Unshift ((f (Succ m)) n)) m)) Empty))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestFold(t *testing.T) {
  a := GetBuiltinVariable("Fold")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "(Y ^f.^l.^x.^g.(((If (IsEmpty l)) x) ((g (((f (Rest l)) x) g)) (First l))))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestMap(t *testing.T) {
  a := GetBuiltinVariable("Map")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^k.^f.(((Fold k) Empty) ^l.^x.((Unshift l) (f x)))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestPush(t *testing.T) {
  a := GetBuiltinVariable("Push")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^l.^x.(((Fold l) ((Unshift Empty) x)) Unshift)" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestY(t *testing.T) {
  a := GetBuiltinVariable("Y")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "^f.(^x.(f (x x)) ^x.(f (x x)))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

func TestDigitToStr(t *testing.T) {
  a := GetBuiltinVariable("DigitToStr")

  ok, b := a.Expand()
  if ! ok {
    t.Fatalf("expression is not expandable")
  }

  if b.String() != "(Y ^f.^n.((Push (((If ((LessOrEq n) 9)) Empty) (f ((Div n) 10)))) " +
                   "((Add ((Mod n) 10)) 48)))" {
    t.Errorf("unexpected expand result: %s", b.String())
  }
}

