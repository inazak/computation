package ast

import (
  "testing"
)

func TestSymbolRef1(t *testing.T) {
  sr := NewSymbolRef()

  d := sr.GetDepth()
  _ = sr.AddIndex('a')
  _ = sr.AddIndex('b')
  _ = sr.AddIndex('c')

  if sr.GetIndex('a') != 3 {
    t.Errorf("'a' is index 3, but got %v", sr.GetIndex('a'))
  }

  if sr.GetIndex('b') != 2 {
    t.Errorf("'b' is index 2, but got %v", sr.GetIndex('b'))
  }

  if sr.GetIndex('c') != 1 {
    t.Errorf("'c' is index 1, but got %v", sr.GetIndex('c'))
  }

  if sr.GetIndex('z') != 4 {
    t.Errorf("'z' is index 4, but got %v", sr.GetIndex('z'))
  }

  sr.RestoreDepth(d)

  if sr.GetIndex('a') != 1 {
    t.Errorf("SymbolRef is restored, but 'a' got unexpected index")
  }

  if sr.GetIndex('z') != 1 {
    t.Errorf("SymbolRef is restored, but 'z' got unexpected index")
  }
}


