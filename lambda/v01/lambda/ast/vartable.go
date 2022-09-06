package ast

type VarTable struct {
  table map[string]Expression
}

func NewVarTable() *VarTable {
  vt := &VarTable{ table: make(map[string]Expression) }

  for _, b := range BUILTINS {
    vt.Set(b.Name, b.Expr)
  }

  return vt
}

func (vt VarTable) Contain (name string) bool {
  _, ok := vt.table[name]
  return ok
}

func (vt VarTable) Set (name string, expr Expression) {
  vt.table[name] = expr
}

func (vt VarTable) Get (name string) Expression {
  return vt.table[name]
}


func (vt *VarTable) UpdateVariable(expr Expression) {

  switch expr.(type) {

  case *Variable:
    v, _ := expr.(*Variable)
    if vt.Contain(v.Name) {
      v.Expr = vt.Get(v.Name)
    }

  case *Function:
    f, _ := expr.(*Function)
    vt.UpdateVariable(f.Body)

  case *Application:
    a, _ := expr.(*Application)
    vt.UpdateVariable(a.Left)
    vt.UpdateVariable(a.Right)

  case *Pair:
    p, _ := expr.(*Pair)
    vt.UpdateVariable(p.Left)
    vt.UpdateVariable(p.Right)

  case *List:
    l, _ := expr.(*List)
    for i, _ := range l.Data {
      vt.UpdateVariable(l.Data[i])
    }
  }
}

func (vt *VarTable) UpdateVariableWithName(expr Expression, name string) {

  switch expr.(type) {

  case *Variable:
    v, _ := expr.(*Variable)
    if v.Name == name && vt.Contain(v.Name) {
      v.Expr = vt.Get(v.Name)
    }

  case *Function:
    f, _ := expr.(*Function)
    vt.UpdateVariableWithName(f.Body, name)

  case *Application:
    a, _ := expr.(*Application)
    vt.UpdateVariableWithName(a.Left, name)
    vt.UpdateVariableWithName(a.Right, name)

  case *Pair:
    p, _ := expr.(*Pair)
    vt.UpdateVariableWithName(p.Left, name)
    vt.UpdateVariableWithName(p.Right, name)

  case *List:
    l, _ := expr.(*List)
    for i, _ := range l.Data {
      vt.UpdateVariableWithName(l.Data[i], name)
    }
  }
}

