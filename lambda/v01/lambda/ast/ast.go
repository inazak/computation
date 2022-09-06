package ast

import (
  "fmt"
  "strconv"
)

type Statement interface {
  String() string
}

type AssignmentStatement struct {
  Var *Variable
}

type ExpressionStatement struct {
  Expr Expression
}

func (as *AssignmentStatement) String() string {
  return as.Var.String()
}

func (es *ExpressionStatement) String() string {
  return es.Expr.String()
}


type Expression interface {
  Indexing(sr *SymbolRef)
  String()        string
  StringByIndex() string
  Copy()          Expression
  UpdateUnboundIndex(v, depth int)
  Reduce() (ok bool, expr Expression)
  Replace(expr Expression, index int) Expression
  Expand() (ok bool, expr Expression)
}

type Symbol struct {
  Name  byte
  Index int
}

type Function struct {
  Arg  byte
  Body Expression
}

type Application struct {
  Left  Expression
  Right Expression
}

type Pair struct {
  Left  Expression
  Right Expression
}

type List struct {
  Data []Expression
}

type Variable struct {
  Name string
  Expr Expression
}

type Number struct {
  Name  string
}

type Char struct {
  C byte
}

type Str struct {
  S string
}

type Blank struct {
}

// implementation of String

func (s *Symbol) String() string {
  return string(s.Name)
}

func (f *Function) String() string {
  return "^" + string(f.Arg) + "." + f.Body.String()
}

func (a *Application) String() string {
  return "(" + a.Left.String() + " " + a.Right.String() + ")"
}

func (p *Pair) String() string {
  return "<" + p.Left.String() + " " + p.Right.String() + ">"
}

func (l *List) String() string {
  if len(l.Data) == 0 {
    return "[]"
  } else if len(l.Data) == 1 {
    return "[" + l.Data[0].String() + "]"
  } else {
    s := "[" + l.Data[0].String()
    for i := 1; i < len(l.Data); i++ {
      s += " " + l.Data[i].String()
    }
    s += "]"
    return s
  }
}

func (v *Variable) String() string {
  return v.Name
}

func (n *Number) String() string {
  return n.Name
}

func (c *Char) String() string {
  return "'" + string(c.C)
}

func (s *Str) String() string {
  return "\"" + s.S + "\""
}

func (b *Blank) String() string {
  return ""
}

// implementation of StringByIndex

func (s *Symbol) StringByIndex() string {
  return fmt.Sprintf("%d", s.Index)
}

func (f *Function) StringByIndex() string {
  return "^." + f.Body.StringByIndex()
}

func (a *Application) StringByIndex() string {
  return "(" + a.Left.StringByIndex() + " " + a.Right.StringByIndex() + ")"
}

func (p *Pair) StringByIndex() string {
  return "<" + p.Left.StringByIndex() + " " + p.Right.StringByIndex() + ">"
}

func (l *List) StringByIndex() string {
  if len(l.Data) == 0 {
    return "[]"
  } else if len(l.Data) == 1 {
    return "[" + l.Data[0].StringByIndex() + "]"
  } else {
    s := "[" + l.Data[0].StringByIndex()
    for i := 1; i < len(l.Data); i++ {
      s += " " + l.Data[i].StringByIndex()
    }
    s += "]"
    return s
  }
}

func (v *Variable) StringByIndex() string {
  return v.Name
}

func (n *Number) StringByIndex() string {
  return n.Name
}

func (c *Char) StringByIndex() string {
  return "'" + string(c.C)
}

func (s *Str) StringByIndex() string {
  return "\"" + s.S + "\""
}

func (b *Blank) StringByIndex() string {
  return ""
}

// implementation of Indexing

func (s *Symbol) Indexing(sr *SymbolRef) {
  s.Index = sr.GetIndex(s.Name)
}

func (f *Function) Indexing(sr *SymbolRef) {
  depth := sr.GetDepth()
  sr.AddIndex(f.Arg)
  f.Body.Indexing(sr)
  sr.RestoreDepth(depth)
}

func (a *Application) Indexing(sr *SymbolRef) {
  a.Left.Indexing(sr)
  a.Right.Indexing(sr)
}

func (p *Pair) Indexing(sr *SymbolRef) {
  //do nothing
}

func (l *List) Indexing(sr *SymbolRef) {
  //do nothing
}

func (v *Variable) Indexing(sr *SymbolRef) {
  //do nothing
}

func (n *Number) Indexing(sr *SymbolRef) {
  //do nothing
}

func (c *Char) Indexing(sr *SymbolRef) {
  //do nothing
}

func (s *Str) Indexing(sr *SymbolRef) {
  //do nothing
}

func (b *Blank) Indexing(sr *SymbolRef) {
  //do nothing
}

// implementation of Copy

func (s *Symbol) Copy() Expression {
  return &Symbol{
    Name:  s.Name,
    Index: s.Index,
  }
}

func (f *Function) Copy() Expression {
  return &Function{
    Arg:  f.Arg,
    Body: f.Body.Copy(),
  }
}

func (a *Application) Copy() Expression {
  return &Application{
    Left:  a.Left.Copy(),
    Right: a.Right.Copy(),
  }
}

func (p *Pair) Copy() Expression {
  return &Pair{
    Left:  p.Left.Copy(),
    Right: p.Right.Copy(),
  }
}

func (l *List) Copy() Expression {
  data := []Expression{}
  for _, d := range l.Data {
    data = append(data, d.Copy())
  }
  return &List{ Data: data }
}

func (v *Variable) Copy() Expression {
  return &Variable {
    Name: v.Name,
    Expr: v.Expr.Copy(),
  }
}

func (n *Number) Copy() Expression {
  return &Number {
    Name: n.Name,
  }
}

func (c *Char) Copy() Expression {
  return &Char {
    C: c.C,
  }
}

func (s *Str) Copy() Expression {
  return &Str {
    S: s.S,
  }
}

func (b *Blank) Copy() Expression {
  return &Blank {}
}

// implementation of UpdateUnboundIndex

func (s *Symbol) UpdateUnboundIndex(i, depth int) {
  if s.Index > depth {
    s.Index += i
  }
}

func (f *Function) UpdateUnboundIndex(i, depth int) {
  f.Body.UpdateUnboundIndex(i, depth + 1)
}

func (a *Application) UpdateUnboundIndex(i, depth int) {
  a.Left.UpdateUnboundIndex(i, depth)
  a.Right.UpdateUnboundIndex(i, depth)
}

func (p *Pair) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (l *List) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (v *Variable) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (n *Number) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (c *Char) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (s *Str) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

func (b *Blank) UpdateUnboundIndex(i, depth int) {
  //do nothing
}

// implementation of Reduce

func (s *Symbol) Reduce() (ok bool, expr Expression) {
  return false, s
}

func (f *Function) Reduce() (ok bool, expr Expression) {
  if ok, reduced := f.Body.Reduce() ; ok {
    f.Body = reduced
    return true, f
  } else {
    return false, f
  }
}

func (a *Application) Reduce() (ok bool, expr Expression) {
  if f, ok := a.Left.(*Function) ; ok {
    reduced := f.Body.Replace(a.Right, 1)
    reduced.UpdateUnboundIndex(-1, 0)
    return true, reduced
  }
  if ok, reduced := a.Left.Reduce() ; ok {
    a.Left = reduced
    return true, a
  }
  if ok, reduced := a.Right.Reduce() ; ok {
    a.Right = reduced
    return true, a
  }

  // has no reducible expression
  return false, a
}

func (p *Pair) Reduce() (ok bool, expr Expression) {
  return false, p
}

func (l *List) Reduce() (ok bool, expr Expression) {
  return false, l
}

func (v *Variable) Reduce() (ok bool, expr Expression) {
  return false, v
}

func (n *Number) Reduce() (ok bool, expr Expression) {
  return false, n
}

func (c *Char) Reduce() (ok bool, expr Expression) {
  return false, c
}

func (s *Str) Reduce() (ok bool, expr Expression) {
  return false, s
}

func (b *Blank) Reduce() (ok bool, expr Expression) {
  return false, b
}

// implementation of Replace

func (s *Symbol) Replace(expr Expression, index int) Expression {
  if s.Index == index {
    expr := expr.Copy()//FIXME
    expr.UpdateUnboundIndex(index, 0)
    return expr
  } else {
    return s
  }
}

func (f *Function) Replace(expr Expression, index int) Expression {
  f.Body = f.Body.Replace(expr, index + 1)
  return f
}

func (a *Application) Replace(expr Expression, index int) Expression {
  a.Left  = a.Left.Replace(expr, index)
  a.Right = a.Right.Replace(expr, index)
  return a
}

func (p *Pair) Replace(expr Expression, index int) Expression {
  return p
}

func (l *List) Replace(expr Expression, index int) Expression {
  return l
}

func (v *Variable) Replace(expr Expression, index int) Expression {
  return v
}

func (n *Number) Replace(expr Expression, index int) Expression {
  return n
}

func (c *Char) Replace(expr Expression, index int) Expression {
  return c
}

func (s *Str) Replace(expr Expression, index int) Expression {
  return s
}

func (b *Blank) Replace(expr Expression, index int) Expression {
  return b
}

// implementation of Expand

func (s *Symbol) Expand() (ok bool, expr Expression) {
  return false, s
}

func (f *Function) Expand() (ok bool, expr Expression) {
  if ok, body := f.Body.Expand() ; ok {
    return true, &Function {
      Arg:  f.Arg,
      Body: body,
    }
  } else {
    return false, f
  }
}

func (a *Application) Expand() (ok bool, expr Expression) {
  if ok, left := a.Left.Expand() ; ok {
    return true, &Application {
      Left: left,
      Right: a.Right,
    }
  }
  if ok, right := a.Right.Expand() ; ok {
    return true, &Application {
      Left: a.Left,
      Right: right,
    }
  }
  return false, a
}

func (p *Pair) Expand() (ok bool, expr Expression) {
  return true, &Application { Left: &Application { Left:  MakeVariable("Pair", VAR_PAIR),
                                                   Right: p.Left },
                              Right: p.Right }

}

func (l *List) Expand() (ok bool, expr Expression) {
  var e Expression
  e = MakeVariable("Empty", VAR_EMPTY)
  for i := len(l.Data) -1; i >= 0; i-- {
    e = &Pair { Left: MakeVariable("False", VAR_FALSE), Right:
        &Pair { Left: l.Data[i], Right: e } }
  }
  return true, e
}

func (v *Variable) Expand() (ok bool, expr Expression) {
  if _, ok := v.Expr.(*Blank) ; ok {
    return false, v
  } else {
    return true, v.Expr.Copy()
  }
}

func (n *Number) Expand() (ok bool, expr Expression) {
  i, err := strconv.Atoi(n.Name)
  if err != nil {
    return false, n //cant expand
  }
  return true, MakeChurchNumber(i)
}

func (c *Char) Expand() (ok bool, expr Expression) {
  return true, MakeChurchNumber(int(c.C))
}

func (s *Str) Expand() (ok bool, expr Expression) {
  data := []Expression{}
  for i := 0; i < len(s.S); i++ {
    data = append(data, &Char{ C: s.S[i] })
  }
  return true, &List{ Data: data }
}

func (b *Blank) Expand() (ok bool, expr Expression) {
  return false, b
}


