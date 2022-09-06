package eval

import (
  "testing"
)

func TestIsPair(t *testing.T) {

  ps := [][]string {
    { "<x y>", "x", "y" },
    { "((Pair x) y)", "x", "y" },
    { "^f.((f x) y)", "x", "y" },
    { "((^a.^b.^f.((f a) b) x) y)", "x", "y" },
    { "<(x y) (Zoo Zoo)>", "(x y)", "(Zoo Zoo)" },
    { "<<x y> z>", "^f.((f x) y)", "z" }, //depend ast.VAR_PAIR
    { "[]", "^x.^y.x", "^x.^y.x" },
  }

  ev  := NewEvaluator(OUTPUT_LAMBDA)

  for i, p := range ps {
    input  := p[0]
    left   := p[1]
    right  := p[2]

    perr, expr, _ := ev.parseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    ok, l, r := IsPair(expr)

    if ! ok {
      t.Fatalf("no=%d IsPair return false", i)
    }

    if l.String() != left {
      t.Errorf("no=%d left is expected=%v, but got=%v", i, left, l.String())
    }

    if r.String() != right {
      t.Errorf("no=%d right is expected=%v, but got=%v", i, right, r.String())
    }
  }
}

func TestIsNotPair(t *testing.T) {

  ps := []string {
    "x",
    "(x y)",
    "^f.(f x)",
    "1",
    "Zoo",
  }

  ev  := NewEvaluator(OUTPUT_LAMBDA)

  for i, input := range ps {

    perr, expr, _ := ev.parseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    ok, _, _ := IsPair(expr)

    if ok {
      t.Errorf("no=%d IsPair expect false, but return true", i)
    }
  }
}

func TestIsEmpty(t *testing.T) {

  ps := []string {
    "Empty",
    "<True True>",
    "<True z>",
    "((^a.^b.^f.((f a) b) ^x.^y.x) z)",
  }

  ev  := NewEvaluator(OUTPUT_LAMBDA)

  for i, input := range ps {

    perr, expr, _ := ev.parseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    ok := IsEmpty(expr)

    if ! ok {
      t.Errorf("no=%d IsEmpty expect true, but return false", i)
    }
  }
}

func TestChurchNumberListToString(t *testing.T) {

  ps := [][]string {
    { "[]", "" },
    { "\"\"", "" },
    { "['x 'y 'z]", "xyz" },
    { "\"xyz\"", "xyz" },
    { "[65 66 67]", "ABC" },
    { "[[65 66] [67 68] [69 70]]", "ABCDEF" },
    { "[65 [66 [67 68] 69] 70]", "ABCDEF" },
    { "[ [] [65] [] ]", "A" },
  }

  ev  := NewEvaluator(OUTPUT_LAMBDA)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, expr, _ := ev.parseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    ok, output := ChurchNumberListToString(expr)

    if ! ok {
      t.Fatalf("no=%d ChurchNumberListToString return false", i)
    }

    if output != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, output)
    }
  }
}
