package test

import (
  "fmt"
  "testing"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/reader"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/parser"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/builder"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)


type Pair struct {
  PortName string
  State    int
}

type TestCase struct {
  Script string
  Expect [][]Pair
}

func MakeSession(script string) (*control.Session, error) {

  sc, e1 := reader.ReadFromString(script)
  if e1 != nil {
    return nil, fmt.Errorf("failure in reader.ReadFromString")
  }

  stmts, e2 := parser.Parse(sc)
  if e2 != nil {
    return nil, fmt.Errorf("failure in parser.Parse: %s", e2.Message)
  }

  sess, e3 := builder.Build(stmts)
  if e3 != nil {
    return nil, fmt.Errorf("failure in builder.Build: %s", e3.Message)
  }

  return sess, nil
}


func DoTest(t *testing.T, tc TestCase) {

  se, err := MakeSession(tc.Script)
  if err != nil {
    t.Fatalf("in MakeSession: %s", err.Error())
  }

  se.PreUpdate() //initialize

  for i, exp := range tc.Expect {
    se.Cycle()
    for _, p := range exp {
      if se.GetState(p.PortName) != p.State {
        t.Errorf("unexpected in times=%d, portname=%s, expect=%d, but got=%d",
          i, p.PortName, p.State, se.GetState(p.PortName))
      }
    }
  }

}


