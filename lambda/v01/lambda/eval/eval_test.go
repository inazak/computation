package eval

import (
  "testing"
)

func TestEval1(t *testing.T) {

  ps := [][]string {
    { "x", "x", },
    { "^x.x", "^x.x", },
    { "Var = a", "", },
  }

  ev  := NewEvaluator(OUTPUT_LAMBDA)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval2(t *testing.T) {

  ps := [][]string {
    { "Var = a", "", },
    { "Bar = (x Var)", "", },
    { "Var", "a", },
    { "Bar", "(x a)", },
  }

  ev  := NewEvaluator(OUTPUT_EXPAND)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval3(t *testing.T) {

  ps := [][]string {
    { "Var = ((^f.^x.^y.((y ^x.(y x)) f) w) z)", "", },
    { "Var", "^y.((y ^x.(y x)) w)", },
  }

  ev  := NewEvaluator(OUTPUT_REDUCE)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval4(t *testing.T) {

  ps := [][]string {
    { "0", "^f.^x.x", },
    { "1", "^f.^x.(f x)", },
    { "2", "^f.^x.(f (f x))", },
    { "(1 2)", "(^f.^x.(f x) ^f.^x.(f (f x)))", },
  }

  ev  := NewEvaluator(OUTPUT_EXPAND)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval5(t *testing.T) {

  ps := [][]string {
    { "'A", "^f.^x.(f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f (f (f (f (f (f" +
                 " (f (f (f (f (f x)))))))))))))" +
                 "))))))))))))))))))))))))))))))" +
                 "))))))))))))))))))))))", },
  }

  ev  := NewEvaluator(OUTPUT_EXPAND)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval6(t *testing.T) {

  ps := [][]string {
    { "<x y>", "((^a.^b.^f.((f a) b) x) y)", },
    { "[x y]", "((^a.^b.^f.((f a) b) ^x.^y.y) ((^a.^b.^f.((f a) b) x) ((^a.^b.^f.((f a) b) ^x.^y.y) " +
               "((^a.^b.^f.((f a) b) y) ((^a.^b.^f.((f a) b) ^x.^y.x) ^x.^y.x)))))", },
  }

  ev  := NewEvaluator(OUTPUT_EXPAND)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval7(t *testing.T) {

  ps := [][]string {
    { "(Left <x y>)", "x", },
    { "(Right <x y>)", "y", },
    { "(First [x y])", "x", },
    { "(Rest [x y])", "^f.((f ^x.^y.y) ^f.((f y) ^f.((f ^x.^y.x) ^x.^y.x)))", },
  }

  ev  := NewEvaluator(OUTPUT_REDUCE)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval8(t *testing.T) {

  ps := [][]string {
    { "(Succ 0)", "^.^.(2 1)", },
    { "(Succ 1)", "^.^.(2 (2 1))", },
    { "(Succ 2)", "^.^.(2 (2 (2 1)))", },
    { "(Pred 1)", "^.^.1", },
    { "(Pred 2)", "^.^.(2 1)", },
    { "(Pred 3)", "^.^.(2 (2 1))", },
    { "(Pred (First (Rest [1 2 3])))", "^.^.(2 1)", },
  }

  ev  := NewEvaluator(OUTPUT_INDEX)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEval9(t *testing.T) {

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

  ev  := NewEvaluator(OUTPUT_ASCII)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

func TestEvalMod(t *testing.T) {

  ps := [][]string {
    { "((Mod 1) 1)", "^.^.1" },
    { "((Mod 2) 1)", "^.^.1" },
    { "((Mod 2) 2)", "^.^.1" },
    { "((Mod 3) 2)", "^.^.(2 1)" },
    { "((Mod 4) 2)", "^.^.1" },
    { "((Mod 5) 2)", "^.^.(2 1)" },
    { "((Mod 1) 3)", "^.^.(2 1)" },
    { "((Mod 2) 3)", "^.^.(2 (2 1))" },
    { "((Mod 3) 3)", "^.^.1" },
    { "((Mod 4) 3)", "^.^.(2 1)" },
    { "((Mod 5) 3)", "^.^.(2 (2 1))" },
    { "((Mod 6) 3)", "^.^.1" },
    { "((Mod 7) 3)", "^.^.(2 1)" },
    { "((Mod 8) 3)", "^.^.(2 (2 1))" },
  }

  ev  := NewEvaluator(OUTPUT_INDEX)

  for i, p := range ps {
    input  := p[0]
    expect := p[1]

    perr, o := ev.ParseAndEval(input)

    if perr != nil { //parse error
      t.Fatalf("parse error: no=%d, %v", i, perr)
    }

    if o.Text != expect {
      t.Errorf("no=%d expected=%v, but got=%v", i, expect, o.Text)
    }
  }
}

