package eval

import (
  "fmt"
  "strings"
  "github.com/inazak/computation/lambda/v01/lambda/ast"
)

func IsPair(expr ast.Expression) (ok bool, left ast.Expression, right ast.Expression)  {

  ev  := NewEvaluator(OUTPUT_REDUCE)
  expr = ev.reduce(expr.Copy())

  expr.Indexing(ast.NewSymbolRef())
  s := expr.StringByIndex()

  if ! strings.HasPrefix(s, "^.((1 ") {
    return false, nil, nil
  }

  f, _  := expr.(*ast.Function)
  a1, _ := f.Body.(*ast.Application)
  a2, _ := a1.Left.(*ast.Application)

  return true, a2.Right.Copy(), a1.Right.Copy()
}

func IsEmpty(expr ast.Expression) (ok bool)  {

  if ok, left, _ := IsPair(expr) ; ok {
    left.Indexing(ast.NewSymbolRef())
    s := left.StringByIndex()
    if s == "^.^.2" {
      return true
    }
  }
  return false
}

func ChurchNumberListToString(expr ast.Expression) (bool, string) {

  if ok, n := ast.IsChurchNumber(expr) ; ok {
    return true, fmt.Sprintf("%c", n)
  }

  s := ""

  for {

    ok, _, r1 := IsPair(expr)
    if ! ok {
      return false, s
    }

    if ok && IsEmpty(expr) {
      break
    }

    ok, l2, r2 := IsPair(r1)
    if ! ok {
      return false, s
    }

    ok, n := ast.IsChurchNumber(l2)
    if ok {
      s = fmt.Sprintf("%s%c", s, n)
    } else {
      if ok, rs := ChurchNumberListToString(l2) ; ok {
        s = s + rs
      } else {
        return false, s + rs
      }
    }

    expr = r2
  }

  return true, s
}


