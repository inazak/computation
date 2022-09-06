package ast

import (
  "testing"
)

func TestMakeChurchNumber(t *testing.T) {

  var e Expression

  e = MakeChurchNumber(0)
  e.Indexing(NewSymbolRef())

  if e.String() != "^f.^x.x" {
    t.Errorf("MakeChurchNumber(0).String() return unexpected value = %v", e)
  }
  if e.StringByIndex() != "^.^.1" {
    t.Errorf("MakeChurchNumber(0).StringByIndex() return unexpected value = %v", e)
  }

  e = MakeChurchNumber(1)
  e.Indexing(NewSymbolRef())

  if e.String() != "^f.^x.(f x)" {
    t.Errorf("MakeChurchNumber(1).String() return unexpected value = %v", e)
  }
  if e.StringByIndex() != "^.^.(2 1)" {
    t.Errorf("MakeChurchNumber(1).StringByIndex() return unexpected value = %v", e)
  }

  e = MakeChurchNumber(2)
  e.Indexing(NewSymbolRef())

  if e.String() != "^f.^x.(f (f x))" {
    t.Errorf("MakeChurchNumber(2).String() return unexpected value = %v", e)
  }
  if e.StringByIndex() != "^.^.(2 (2 1))" {
    t.Errorf("MakeChurchNumber(2).StringByIndex() return unexpected value = %v", e)
  }

  e = MakeChurchNumber(3)
  e.Indexing(NewSymbolRef())

  if e.String() != "^f.^x.(f (f (f x)))" {
    t.Errorf("MakeChurchNumber(3).String() return unexpected value = %v", e)
  }
  if e.StringByIndex() != "^.^.(2 (2 (2 1)))" {
    t.Errorf("MakeChurchNumber(3).StringByIndex() return unexpected value = %v", e)
  }
}

func TestIsChurchNumber(t *testing.T) {
  for i := 0; i < 9; i++ {
    e := MakeChurchNumber(i)
    e.Indexing(NewSymbolRef())
    if ok, n := IsChurchNumber(e) ; ! ok || n != i {
      t.Errorf("expected=%v, got= %v", i, n)
    }
  }
}

