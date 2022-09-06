package ast

import (
  "testing"
)

func TestIndexing1(t *testing.T) {

  ps := []struct{
    Input  string
    Expect string
    Data   Expression
  }{
    {
      Input:  "x",
      Expect: "1",
      Data:   &Symbol { Name: 'x' },
    },
    {
      Input:  "(x ^x.^y.^z.((x y) (z a)))",
      Expect: "(1 ^.^.^.((3 2) (1 4)))",
      Data: &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Function { Arg: 'x', Body:
            &Function { Arg: 'y', Body:
            &Function { Arg: 'z', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Symbol { Name: 'y' }, } , Right:
            &Application { Left:
            &Symbol { Name: 'z' } , Right:
            &Symbol { Name: 'a' }, }, } } } }, },
    },
    {
      Input:  "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))", //(Succ 1)
      Expect: "(^.^.^.(2 ((3 2) 1)) ^.^.(2 1))",
      Data:   &Application { Left:
              &Function { Arg: 'n', Body:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Symbol { Name: 'n' } , Right:
              &Symbol { Name: 'f' }, } , Right:
              &Symbol { Name: 'x' }, }, } } } } , Right:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } }, },
    },
    {
      Input:  "^f.^x.(f ((^f.^x.(f x) f) x))",
      Expect: "^.^.(2 ((^.^.(2 1) 2) 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } } , Right:
              &Symbol { Name: 'f' }, } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
    {
      Input:  "^f.^x.(f (^x.(f x) x))",
      Expect: "^.^.(2 (^.(3 1) 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
    {
      Input:  "^f.^x.(f (f x))",
      Expect: "^.^.(2 (2 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
  }

  for i, p := range ps {

    if p.Data.String() != p.Input {
      t.Fatalf("no.%d input='%s' but data got='%s'", i, p.Input, p.Data.String())
    }

    p.Data.Indexing(NewSymbolRef())

    if p.Data.StringByIndex() != p.Expect {
      t.Errorf("no.%d input='%s' expect='%s', but got='%s'",
        i, p.Data.String(), p.Expect, p.Data.StringByIndex())
    }
  }
}

func TestUpdateUnboundIndex1(t *testing.T) {

  ps := []struct{
    Input  string
    Before string
    After  string
    N      int
    Data   Expression
  }{
    {
      Input:  "(x ^x.^y.^z.((x y) (z a)))",
      Before: "(1 ^.^.^.((3 2) (1 4)))",
      After:  "(2 ^.^.^.((3 2) (1 5)))",
      N:      1,
      Data: &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Function { Arg: 'x', Body:
            &Function { Arg: 'y', Body:
            &Function { Arg: 'z', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Symbol { Name: 'y' }, } , Right:
            &Application { Left:
            &Symbol { Name: 'z' } , Right:
            &Symbol { Name: 'a' }, }, } } } }, },
    },
    {
      Input:  "(x ^x.^y.^z.((x y) (z a)))",
      Before: "(1 ^.^.^.((3 2) (1 4)))",
      After:  "(0 ^.^.^.((3 2) (1 3)))",
      N:      -1,
      Data: &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Function { Arg: 'x', Body:
            &Function { Arg: 'y', Body:
            &Function { Arg: 'z', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Symbol { Name: 'y' }, } , Right:
            &Application { Left:
            &Symbol { Name: 'z' } , Right:
            &Symbol { Name: 'a' }, }, } } } }, },
    },
    {
      Input:  "^x.(x ^x.^y.^z.((x y) (z a)))",
      Before: "^.(1 ^.^.^.((3 2) (1 5)))",
      After:  "^.(1 ^.^.^.((3 2) (1 6)))",
      N:      1,
      Data: &Function { Arg: 'x', Body:
            &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Function { Arg: 'x', Body:
            &Function { Arg: 'y', Body:
            &Function { Arg: 'z', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'x' } , Right:
            &Symbol { Name: 'y' }, } , Right:
            &Application { Left:
            &Symbol { Name: 'z' } , Right:
            &Symbol { Name: 'a' }, }, } } } }, } },
    },
  }

  for i, p := range ps {

    if p.Data.String() != p.Input {
      t.Fatalf("no.%d input='%s' but data got='%s'", i, p.Input, p.Data.String())
    }

    p.Data.Indexing(NewSymbolRef())

    if p.Data.StringByIndex() != p.Before {
      t.Fatalf("no.%d before='%s' but data got='%s'", i, p.Before, p.Data.StringByIndex())
    }

    p.Data.UpdateUnboundIndex(p.N, 0)

    if p.Data.StringByIndex() != p.After {
      t.Errorf("no.%d expect='%s', but got='%s'", i, p.After, p.Data.StringByIndex())
    }
  }
}


func TestReduce1(t *testing.T) {

  ps := []struct{
    Input  string
    Before string
    After  string
    Data   Expression
  }{
    {
      Input:  "(^x.x y)",
      Before: "(^.1 1)",
      After:  "1",
      Data: &Application { Left:
            &Function { Arg: 'x', Body:
            &Symbol { Name: 'x' } }, Right:
            &Symbol { Name: 'y' } },
    },
    {
      Input:  "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))",
      Before: "(^.^.^.(2 ((3 2) 1)) ^.^.(2 1))",
      After:  "^.^.(2 ((^.^.(2 1) 2) 1))",
      Data:   &Application { Left:
              &Function { Arg: 'n', Body:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Symbol { Name: 'n' } , Right:
              &Symbol { Name: 'f' }, } , Right:
              &Symbol { Name: 'x' }, }, } } } } , Right:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } }, },
    },
    {
      Input:  "^f.^x.(f ((^f.^x.(f x) f) x))",
      Before: "^.^.(2 ((^.^.(2 1) 2) 1))",
      After:  "^.^.(2 (^.(3 1) 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } } , Right:
              &Symbol { Name: 'f' }, } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
    {
      Input:  "^f.^x.(f (^x.(f x) x))",
      Before: "^.^.(2 (^.(3 1) 1))",
      After:  "^.^.(2 (2 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
    {
      Input:  "(^n.^f.^x.(f ((n a) x)) ^f.^x.(f b))",
      Before: "(^.^.^.(2 ((3 4) 1)) ^.^.(2 3))",
      After:  "^.^.(2 ((^.^.(2 5) 3) 1))",
      Data:   &Application { Left:
              &Function { Arg: 'n', Body:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Symbol { Name: 'n' } , Right:
              &Symbol { Name: 'a' }, } , Right:
              &Symbol { Name: 'x' }, }, } } } } , Right:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'b' }, } } }, },
    },
    {
      Input:  "^f.^x.(f ((^f.^x.(f b) a) x))",
      Before: "^.^.(2 ((^.^.(2 5) 3) 1))",
      After:  "^.^.(2 (^.(4 4) 1))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'b' }, } } } , Right:
              &Symbol { Name: 'a' }, } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
    {
      Input:  "^f.^x.(f (^x.(a b) x))",
      Before: "^.^.(2 (^.(4 4) 1))",
      After:  "^.^.(2 (3 3))",
      Data:   &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'a' } , Right:
              &Symbol { Name: 'b' }, } } , Right:
              &Symbol { Name: 'x' }, }, } } },
    },
  }

  for i, p := range ps {

    if p.Data.String() != p.Input {
      t.Fatalf("no.%d input='%s' but data got='%s'", i, p.Input, p.Data.String())
    }

    p.Data.Indexing(NewSymbolRef())

    if p.Data.StringByIndex() != p.Before {
      t.Fatalf("no.%d before='%s' but data got='%s'", i, p.Before, p.Data.StringByIndex())
    }

    ok, r := p.Data.Reduce()

    if ! ok {
      t.Errorf("no.%d reduce return false", i, )
      continue
    }

    if r.StringByIndex() != p.After {
      t.Errorf("no.%d expect='%s', but got='%s'", i, p.After, r.StringByIndex())
    }
  }
}

func TestReduce2(t *testing.T) {

  ps := []struct{
    Input  string
    Begin  string
    End    string
    Data   Expression
  }{
    {
      Input:  "(^x.x y)",
      Begin: "(^.1 1)",
      End:   "1",
      Data: &Application { Left:
            &Function { Arg: 'x', Body:
            &Symbol { Name: 'x' } }, Right:
            &Symbol { Name: 'y' } },
    },
    {
      Input:  "(^n.^f.^x.(f ((n f) x)) ^f.^x.(f x))", // (Succ 1)
      Begin:  "(^.^.^.(2 ((3 2) 1)) ^.^.(2 1))",
      End:    "^.^.(2 (2 1))",
      Data:   &Application { Left:
              &Function { Arg: 'n', Body:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Application { Left:
              &Application { Left:
              &Symbol { Name: 'n' } , Right:
              &Symbol { Name: 'f' }, } , Right:
              &Symbol { Name: 'x' }, }, } } } } , Right:
              &Function { Arg: 'f', Body:
              &Function { Arg: 'x', Body:
              &Application { Left:
              &Symbol { Name: 'f' } , Right:
              &Symbol { Name: 'x' }, } } }, },
    },
    {
      Input:  "(^n.(^p.(p ^a.^b.a) ((n ^p.((^a.^b.^f.((f a) b) (^p.(p ^a.^b.b) p)) " +
              "(^n.^f.^x.(f ((n f) x)) (^p.(p ^a.^b.b) p)))) ((^a.^b.^f.((f a) b) ^f.^x.x) ^f.^x.x))) " +
              "^f.^x.(f (f (f x))))", // (Pred 3)
      Begin:  "(^.(^.(1 ^.^.2) ((1 ^.((^.^.^.((1 3) 2) (^.(1 ^.^.1) 1)) " +
              "(^.^.^.(2 ((3 2) 1)) (^.(1 ^.^.1) 1)))) ((^.^.^.((1 3) 2) ^.^.1) ^.^.1))) " +
              "^.^.(2 (2 (2 1))))",
      End:    "^.^.(2 (2 1))",
      Data: &Application { Left:
            &Function { Arg: 'n', Body:
            &Application { Left:
            &Function { Arg: 'p', Body:
            &Application { Left:
            &Symbol { Name: 'p' } , Right:
            &Function { Arg: 'a', Body:
            &Function { Arg: 'b', Body:
            &Symbol { Name: 'a' } } }, } } , Right:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'n' } , Right:
            &Function { Arg: 'p', Body:
            &Application { Left:
            &Application { Left:
            &Function { Arg: 'a', Body:
            &Function { Arg: 'b', Body:
            &Function { Arg: 'f', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Symbol { Name: 'a' }, } , Right:
            &Symbol { Name: 'b' }, } } } } , Right:
            &Application { Left:
            &Function { Arg: 'p', Body:
            &Application { Left:
            &Symbol { Name: 'p' } , Right:
            &Function { Arg: 'a', Body:
            &Function { Arg: 'b', Body:
            &Symbol { Name: 'b' } } }, } } , Right:
            &Symbol { Name: 'p' }, }, } , Right:
            &Application { Left:
            &Function { Arg: 'n', Body:
            &Function { Arg: 'f', Body:
            &Function { Arg: 'x', Body:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'n' } , Right:
            &Symbol { Name: 'f' }, } , Right:
            &Symbol { Name: 'x' }, }, } } } } , Right:
            &Application { Left:
            &Function { Arg: 'p', Body:
            &Application { Left:
            &Symbol { Name: 'p' } , Right:
            &Function { Arg: 'a', Body:
            &Function { Arg: 'b', Body:
            &Symbol { Name: 'b' } } }, } } , Right:
            &Symbol { Name: 'p' }, }, }, } }, } , Right:
            &Application { Left:
            &Application { Left:
            &Function { Arg: 'a', Body:
            &Function { Arg: 'b', Body:
            &Function { Arg: 'f', Body:
            &Application { Left:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Symbol { Name: 'a' }, } , Right:
            &Symbol { Name: 'b' }, } } } } , Right:
            &Function { Arg: 'f', Body:
            &Function { Arg: 'x', Body:
            &Symbol { Name: 'x' } } }, } , Right:
            &Function { Arg: 'f', Body:
            &Function { Arg: 'x', Body:
            &Symbol { Name: 'x' } } }, }, }, } } , Right:
            &Function { Arg: 'f', Body:
            &Function { Arg: 'x', Body:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Application { Left:
            &Symbol { Name: 'f' } , Right:
            &Symbol { Name: 'x' }, }, } } } }, },
    },
  }

  for i, p := range ps {

    if p.Data.String() != p.Input {
      t.Fatalf("no.%d input='%s' but data got='%s'", i, p.Input, p.Data.String())
    }

    p.Data.Indexing(NewSymbolRef())

    if p.Data.StringByIndex() != p.Begin {
      t.Fatalf("no.%d begin='%s' but data got='%s'", i, p.Begin, p.Data.StringByIndex())
    }

    ok, r := p.Data.Reduce()

    if ! ok {
      t.Errorf("no.%d reduce return false", i, )
      continue
    }

    for {
      if ok, r = r.Reduce() ; ! ok { break }
    }

    if r.StringByIndex() != p.End {
      t.Errorf("no.%d expect='%s', but got='%s'", i, p.End, r.StringByIndex())
    }
  }
}

func TestExpandPair(t *testing.T) {
  p := &Pair {
    Left:  &Symbol { Name: 'x' },
    Right: &Symbol { Name: 'y' },
  }

  _, expanded := p.Expand()
  if expanded.String() != "((Pair x) y)" {
    t.Errorf("unexpected result, got=%s", expanded.String())
  }
}

func TestExpandList(t *testing.T) {
  var d = []Expression{}
  d = append(d, &Symbol { Name: 'x' })
  d = append(d, &Symbol { Name: 'y' })
  d = append(d, &Symbol { Name: 'z' })
  p := &List { Data: d }

  _, expanded := p.Expand()
  if expanded.String() != "<False <x <False <y <False <z Empty>>>>>>" {
    t.Errorf("unexpected result, got=%s", expanded.String())
  }
}

func TestExpandStr(t *testing.T) {
  p := &Str{ S: "xyz" }

  if p.String() != "\"xyz\"" {
    t.Errorf("p.String() unexpected result, got=%s", p.String())
  }

  _, expand1 := p.Expand()
  if expand1.String() != "['x 'y 'z]" {
    t.Errorf("expand1.String() unexpected result, got=%s", expand1.String())
  }

  _, expand2 := expand1.Expand()
  if expand2.String() != "<False <'x <False <'y <False <'z Empty>>>>>>" {
    t.Errorf("expand2.String() unexpected result, got=%s", expand2.String())
  }
}

