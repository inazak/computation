package bench

import (
  "github.com/inazak/computation/lambda/v01/lambda/ast"
  "testing"
)

// $ go test -bench .

func BenchmarkReduceMod20Normal(b *testing.B) {

  // use ((Mod 20) 1) -> ^.^.1
  //
  // (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^
  // y.y)) ^x.^y.x) ((^x.^y.((y ^n.(^p.(p ^a.^b.a) ((n ^p.((^a.^b.^f.((f a) b) (^p.(p
  //  ^a.^b.b) p)) (^n.^f.^x.(f ((n f) x)) (^p.(p ^a.^b.b) p)))) ((^a.^b.^f.((f a) b)
  //  ^f.^x.x) ^f.^x.x)))) x) x) y)) n) m)) ^x.(((f ((^x.^y.((y ^n.(^p.(p ^a.^b.a) ((
  // n ^p.((^a.^b.^f.((f a) b) (^p.(p ^a.^b.b) p)) (^n.^f.^x.(f ((n f) x)) (^p.(p ^a.
  // ^b.b) p)))) ((^a.^b.^f.((f a) b) ^f.^x.x) ^f.^x.x)))) x) m) n)) n) x)) m)) ^f.^x
  // .(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))))))))))
  // ))) ^f.^x.(f x))

  expect := "^.^.1"
  data := &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'m', Body:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'x'}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'y'}}}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'x'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'y'}, Right:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'a'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'n', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Symbol{Name:'f'}}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}}}}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'x'}}, Right:
          &ast.Symbol{Name:'y'}}}}}, Right:
          &ast.Symbol{Name:'n'}}, Right:
          &ast.Symbol{Name:'m'}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'y'}, Right:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'a'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'n', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Symbol{Name:'f'}}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}}}}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'m'}}, Right:
          &ast.Symbol{Name:'n'}}}, Right:
          &ast.Symbol{Name:'n'}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'m'}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'x'}}}}}}}}}}}}}}}}}}}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'x'}}}}}

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    sample := data.Copy() //FIXME

    sample.Indexing(ast.NewSymbolRef())
    ok, r := sample.Reduce()

    if ! ok {
      b.Fatalf("reduce return false")
    }

    for {
      if ok, r = r.Reduce() ; ! ok { break }
    }

    if r.StringByIndex() != expect {
      b.Fatalf("expect='%s', but got='%s'", expect, r.StringByIndex())
    }
  }
}

/* --- UseGoroutines method

import (
  ...
  "sync"
)

type Expression interface {
  ...
  ReduceUseGoroutines() (ok bool, expr Expression)
  ReplaceUseGoroutines(expr Expression, index int) Expression
}


func (s *ast.Symbol) ReduceUseGoroutines() (ok bool, expr Expression) {
  return false, s
}

func (f *ast.Function) ReduceUseGoroutines() (ok bool, expr Expression) {
  if ok, reduced := f.Body.ReduceUseGoroutines() ; ok {
    f.Body = reduced
    return true, f
  } else {
    return false, f
  }
}

func (a *ast.Application) ReduceUseGoroutines() (ok bool, expr Expression) {
  if f, ok := a.Left.(*ast.Function) ; ok {
    reduced := f.Body.ReplaceUseGoroutines(a.Right, 1)
    reduced.UpdateUnboundIndex(-1, 0)
    return true, reduced
  }
  if ok, reduced := a.Left.ReduceUseGoroutines() ; ok {
    a.Left = reduced
    return true, a
  }
  if ok, reduced := a.Right.ReduceUseGoroutines() ; ok {
    a.Right = reduced
    return true, a
  }

  // has no reducible expression
  return false, a
}


func (s *ast.Symbol) ReplaceUseGoroutines(expr Expression, index int) Expression {
  if s.Index == index {
    expr := expr.Copy()
    expr.UpdateUnboundIndex(index, 0)
    return expr
  } else {
    return s
  }
}

func (f *ast.Function) ReplaceUseGoroutines(expr Expression, index int) Expression {
  f.Body = f.Body.ReplaceUseGoroutines(expr, index + 1)
  return f
}

func (a *ast.Application) ReplaceUseGoroutines(expr Expression, index int) Expression {
  wg := sync.WaitGroup{}
  wg.Add(2)
  go func() {
    defer wg.Done()
    a.Left  = a.Left.ReplaceUseGoroutines(expr, index)
  }()
  go func() {
    defer wg.Done()
    a.Right = a.Right.ReplaceUseGoroutines(expr, index)
  }()
  wg.Wait()
  return a
}

*/

/* --- UseGoroutines Benchmark

func BenchmarkReduceMod20Parallel(b *testing.B) {

  // use ((Mod 20) 1) -> ^.^.1
  //
  // (((^f.(^x.(f (x x)) ^x.(f (x x))) ^f.^m.^n.(((^b.b ((^x.^y.(^x.((x (^x.^y.x ^x.^
  // y.y)) ^x.^y.x) ((^x.^y.((y ^n.(^p.(p ^a.^b.a) ((n ^p.((^a.^b.^f.((f a) b) (^p.(p
  //  ^a.^b.b) p)) (^n.^f.^x.(f ((n f) x)) (^p.(p ^a.^b.b) p)))) ((^a.^b.^f.((f a) b)
  //  ^f.^x.x) ^f.^x.x)))) x) x) y)) n) m)) ^x.(((f ((^x.^y.((y ^n.(^p.(p ^a.^b.a) ((
  // n ^p.((^a.^b.^f.((f a) b) (^p.(p ^a.^b.b) p)) (^n.^f.^x.(f ((n f) x)) (^p.(p ^a.
  // ^b.b) p)))) ((^a.^b.^f.((f a) b) ^f.^x.x) ^f.^x.x)))) x) m) n)) n) x)) m)) ^f.^x
  // .(f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f (f x))))))))))))))))))
  // ))) ^f.^x.(f x))

  expect := "^.^.1"
  data := &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'m', Body:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'x'}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'x'}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'y'}}}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Symbol{Name:'x'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'y'}, Right:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'a'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'n', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Symbol{Name:'f'}}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}}}}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'x'}}, Right:
          &ast.Symbol{Name:'y'}}}}}, Right:
          &ast.Symbol{Name:'n'}}, Right:
          &ast.Symbol{Name:'m'}}}, Right:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'x', Body:
          &ast.Function{Arg:'y', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'y'}, Right:
          &ast.Function{Arg:'n', Body:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'a'}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'n', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'n'}, Right:
          &ast.Symbol{Name:'f'}}, Right:
          &ast.Symbol{Name:'x'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Function{Arg:'p', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'p'}, Right:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Symbol{Name:'p'}}}}}}, Right:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Function{Arg:'a', Body:
          &ast.Function{Arg:'b', Body:
          &ast.Function{Arg:'f', Body:
          &ast.Application{Left:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'a'}}, Right:
          &ast.Symbol{Name:'b'}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Symbol{Name:'x'}}}}}}}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'m'}}, Right:
          &ast.Symbol{Name:'n'}}}, Right:
          &ast.Symbol{Name:'n'}}, Right:
          &ast.Symbol{Name:'x'}}}}, Right:
          &ast.Symbol{Name:'m'}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'x'}}}}}}}}}}}}}}}}}}}}}}}}, Right:
          &ast.Function{Arg:'f', Body:
          &ast.Function{Arg:'x', Body:
          &ast.Application{Left:
          &ast.Symbol{Name:'f'}, Right:
          &ast.Symbol{Name:'x'}}}}}

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    sample := data.Copy() //FIXME

    sample.Indexing(ast.NewSymbolRef())
    ok, r := sample.ReduceUseGoroutines()

    if ! ok {
      b.Fatalf("reduce return false")
    }

    for {
      if ok, r = r.ReduceUseGoroutines() ; ! ok { break }
    }

    if r.StringByIndex() != expect {
      b.Fatalf("expect='%s', but got='%s'", expect, r.StringByIndex())
    }
  }
}
*/

