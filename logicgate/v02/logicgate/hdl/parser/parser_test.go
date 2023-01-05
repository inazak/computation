package parser

import (
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/reader"
)

func TestParser(t *testing.T) {

  script := `
    LO 0001
    HI 0004
    NAND 0001 0002 0003 //comment 
    NAND 0004 0005 0006

    // comment comment
    CONNECT 0006 9999
  `

  sc, rerr := reader.ReadFromString(script)
  if rerr != nil {
    t.Fatalf(rerr.Message)
  }

  stmts, perr := Parse(sc)
  if perr != nil {
    t.Fatalf(perr.Message)
  }

  if len(stmts) != 5 {
    t.Fatalf("statements is not expected length")
  }

  if s, ok := stmts[0].(LOStatement) ; ok {
    if s.Output != "0001" {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 2 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[1].(HIStatement) ; ok {
    if s.Output != "0004" {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 3 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[2].(NandStatement) ; ok {
    if (s.InputA != "0001") || (s.InputB != "0002") || (s.Output != "0003") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 4 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[3].(NandStatement) ; ok {
    if (s.InputA != "0004") || (s.InputB != "0005") || (s.Output != "0006") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 5 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[4].(ConnectStatement) ; ok {
    if (s.Input != "0006") || (s.Output != "9999") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 8 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }
}

