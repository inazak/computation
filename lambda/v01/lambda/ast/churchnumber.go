package ast

import (
  "strings"
)

func MakeChurchNumber(i int) *Function {
  var expr Expression

  expr = &Symbol { Name: 'x', Index: 0 }
  for ; i > 0; i -= 1 {
    expr = &Application {
      Left:  &Symbol { Name: 'f', Index: 0 },
      Right: expr,
    }
  }

  expr = &Function {
    Arg:  'f',
    Body: &Function {
      Arg:  'x',
      Body: expr,
    },
  }

  f, _ := expr.(*Function)

  return f
}


func IsChurchNumber(expr Expression) (bool, int) {

  s := expr.StringByIndex()

  if ! strings.HasPrefix(s, "^.^.") {
    return false, 0
  }

  s = s[4:] //delete "^.^."

  if s == "1" {
    return true, 0
  }

  count := 1
  for {
    if s == "(2 1)" {
      return true, count
    } else if strings.HasPrefix(s, "(2 ") && strings.HasSuffix(s, ")") {
      s = s[3:]
      s = s[:len(s)-1]
      count += 1
    } else {
      break
    }
  }

  return false, 0
}

