package builder

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/reader"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/parser"
)

func TestBuilder(t *testing.T) {

  script := `
    NAND 0001 0002 0003 //comment 
    NAND 0004 0005 0006

    // comment comment
    CONNECT 0006 9999
  `

  sc, rerr := reader.ReadFromString(script)
  if rerr != nil {
    t.Fatalf(rerr.Message)
  }

  stmts, perr := parser.Parse(sc)
  if perr != nil {
    t.Fatalf(perr.Message)
  }

  session, berr := Build(stmts)
  if berr != nil {
    t.Fatalf(berr.Message)
  }

  if session == nil {
    t.Fatalf("session is nil")
  }
}

