package ast

const SYMBOLREF_SIZE = 1024

type SymbolRef struct {
  ref       [SYMBOLREF_SIZE]byte
  currDepth int
}

func NewSymbolRef() *SymbolRef {
  return &SymbolRef{
    ref:       [SYMBOLREF_SIZE]byte{},
    currDepth: 0,
  }
}

func (sr *SymbolRef) GetIndex(c byte) int {
  for i := 1; sr.currDepth - i >= 0; i++ {
    if sr.ref[sr.currDepth - i] == c {
      return i
    }
  }
  return sr.currDepth + 1
}

func (sr *SymbolRef) AddIndex(c byte) int {
  sr.ref[sr.currDepth] = c
  sr.currDepth += 1
  if sr.currDepth >= SYMBOLREF_SIZE {
    panic("SymbolRef array overflow")
  }
  return sr.currDepth
}

func (sr *SymbolRef) GetDepth() int {
  return sr.currDepth
}

func (sr *SymbolRef) RestoreDepth(d int) {
  sr.currDepth = d
}

