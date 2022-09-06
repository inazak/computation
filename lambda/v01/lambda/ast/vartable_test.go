package ast

import (
  "testing"
)

func TestVarTable1(t *testing.T) {

  vt := NewVarTable()

  vt.Set("Var1", &Symbol{ Name: 'a' })
  vt.Set("Var2", &Symbol{ Name: 'b' })
  vt.Set("Var3", &Symbol{ Name: 'c' })

  vt.Set("Var2", &Symbol{ Name: 'x' }) //Update

  if ! vt.Contain("Var1") {
    t.Errorf("VarTable not contain Var1")
  }

  if ! vt.Contain("Var2") {
    t.Errorf("VarTable not contain Var2")
  }

  if ! vt.Contain("Var3") {
    t.Errorf("VarTable not contain Var3")
  }

  if vt.Contain("Var9") {
    t.Errorf("VarTable contain unknown Var9")
  }

  v1 := vt.Get("Var1")
  if s1, ok := v1.(*Symbol) ; ! ok || s1.Name != 'a' {
    t.Errorf("Var1 is incorrect")
  }

  v2 := vt.Get("Var2")
  if s2, ok := v2.(*Symbol) ; ! ok || s2.Name != 'x' {
    t.Errorf("Var2 is incorrect")
  }

  v3 := vt.Get("Var3")
  if s3, ok := v3.(*Symbol) ; ! ok || s3.Name != 'c' {
    t.Errorf("Var3 is incorrect")
  }
}


func TestVarTable2(t *testing.T) {

  vt := NewVarTable()

  vt.Set("Var1", &Symbol{ Name: 'a' })
  vt.Set("Var2", &Symbol{ Name: 'b' })
  vt.Set("Var3", &Symbol{ Name: 'c' })

  expr := &Application { Left:
          &Variable { Name: "Var3", Expr: &Blank {} }, Right:
          &Application { Left:
          &Variable { Name: "Var1", Expr: &Blank {} }, Right:
          &Variable { Name: "Var9", Expr: &Blank {} }, }, }

  if expr.String() != "(Var3 (Var1 Var9))" {
    t.Errorf("expr is unexpected")
  }

  vt.UpdateVariable(expr)

  ok, expanded := expr.Expand()
  for ok {
    ok, expanded = expanded.Expand()
  }

  if expanded.String() != "(c (a Var9))" {
    t.Errorf("updated result is unexpected %s", expanded.String())
  }
}

func TestVarTable3(t *testing.T) {

  vt := NewVarTable()

  vt.Set("Var1", &Symbol{ Name: 'a' })
  vt.Set("Var2", &Symbol{ Name: 'b' })
  vt.Set("Var3", &Symbol{ Name: 'c' })

  expr := &Application { Left:
          &Variable { Name: "Var3", Expr: &Blank {} }, Right:
          &Application { Left:
          &Variable { Name: "Var1", Expr: &Blank {} }, Right:
          &Variable { Name: "Var9", Expr: &Blank {} }, }, }

  if expr.String() != "(Var3 (Var1 Var9))" {
    t.Errorf("expr is unexpected")
  }

  vt.UpdateVariableWithName(expr, "Var1")

  ok, expanded := expr.Expand()
  for ok {
    ok, expanded = expanded.Expand()
  }

  if expanded.String() != "(Var3 (a Var9))" {
    t.Errorf("updated result is unexpected %s", expanded.String())
  }
}

