package ast

import (
  "testing"
)

func TestExpressionString(t *testing.T) {

  ps := []struct{
    Expect string
    Data   Expression
  }{
    {
      Expect: "x",
      Data:   Symbol { Name: "x" },
    },
    {
      Expect: "(x ^x.^y.^z.((x y) (z a)))",
      Data: Application { Left:
            Symbol { Name: "x" } , Right:
            Function { Arg: "x", Body:
            Function { Arg: "y", Body:
            Function { Arg: "z", Body:
            Application { Left:
            Application { Left:
            Symbol { Name: "x" } , Right:
            Symbol { Name: "y" }, } , Right:
            Application { Left:
            Symbol { Name: "z" } , Right:
            Symbol { Name: "a" }, }, } } } }, },
    },
  }

  for i, p := range ps {
    if p.Data.String() != p.Expect {
      t.Errorf("no.%d expect='%s' but got='%s'", i, p.Expect, p.Data.String())
    }
  }
}

