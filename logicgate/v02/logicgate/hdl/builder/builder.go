package builder

import (
  "fmt"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/parser"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)

type BuilderError struct {
  Message string
}

func Build(stmts []parser.Statement) (se *control.Session, err *BuilderError) {

  se = control.MakeSession()
  connectstmt := []parser.Statement{}

  for _, stmt := range stmts {

    switch v := stmt.(type) {
    case parser.NandStatement:
      if e := se.AddNand(v.InputA, v.InputB, v.Output) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.LOStatement:
      if e := se.AddLO(v.Output) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.HIStatement:
      if e := se.AddHI(v.Output) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.CLOCKStatement:
      if e := se.AddCLOCK(v.Output) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.LEDStatement:
      if e := se.AddLED(v.Input) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.HALTStatement:
      if e := se.AddHALT(v.Input) ; e != nil {
        err = &BuilderError{
          Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
        }
        return nil, err
      }

    case parser.ConnectStatement:
      connectstmt = append(connectstmt, stmt)

    default:
      err = &BuilderError{
        Message: fmt.Sprintf("line %d, builder error `unknown statement`", stmt.GetSourceCodeLineNo()),
      }
      return nil, err
    }
  } //for

  prevsize := -1
  for len(connectstmt) > 0 {

    if len(connectstmt) == prevsize {
      return nil, err
    }
    prevsize = len(connectstmt)

    rest := []parser.Statement{}
    for _, stmt := range connectstmt {

      switch v := stmt.(type) {
      case parser.ConnectStatement:
        if e := se.Connect(v.Input, v.Output) ; e != nil {
          err = &BuilderError{
            Message: fmt.Sprintf("line %d, builder error `%s`", stmt.GetSourceCodeLineNo(), e.Message),
          }
          rest = append(rest, stmt)
        }

      default:
        panic("connect statement processing in builder")
      }
    }

    connectstmt = rest
  } // for len(connectstmt) > 0

  if list := se.GetUnconnectedPortName() ; len(list) != 0 {
    err = &BuilderError{
      Message: fmt.Sprintf("builder error, has unconnected port `%v`", list),
    }
    return nil, err
  }

  return se, nil
}

