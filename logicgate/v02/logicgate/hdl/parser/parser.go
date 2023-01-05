package parser

import (
  "fmt"
  "regexp"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/reader"
)

type ParserError struct {
  Message string
}

type Statement interface {
  GetStatementString()  string
  GetSourceCodeLineNo() int
}

type NandStatement struct {
  InputA string
  InputB string
  Output string
  LineNo int
}

type ConnectStatement struct {
  Input  string
  Output string
  LineNo int
}

type LOStatement struct {
  Output string
  LineNo int
}

type HIStatement struct {
  Output string
  LineNo int
}

type CLOCKStatement struct {
  Output string
  LineNo int
}

type HALTStatement struct {
  Input  string
  LineNo int
}

type LEDStatement struct {
  Input  string
  LineNo int
}

func (s NandStatement) GetStatementString() string {
  return fmt.Sprintf("NAND %s %s %s", s.InputA, s.InputB, s.Output)
}

func (s NandStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s ConnectStatement) GetStatementString() string {
  return fmt.Sprintf("CONNECT %s %s", s.Input, s.Output)
}

func (s ConnectStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s LOStatement) GetStatementString() string {
  return fmt.Sprintf("LO %s", s.Output)
}

func (s LOStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s HIStatement) GetStatementString() string {
  return fmt.Sprintf("HI %s", s.Output)
}

func (s HIStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s CLOCKStatement) GetStatementString() string {
  return fmt.Sprintf("CLOCK %s", s.Output)
}

func (s CLOCKStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s HALTStatement) GetStatementString() string {
  return fmt.Sprintf("HALT %s", s.Input)
}

func (s HALTStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}

func (s LEDStatement) GetStatementString() string {
  return fmt.Sprintf("LED %s", s.Input)
}

func (s LEDStatement) GetSourceCodeLineNo() int {
  return s.LineNo
}


func Parse(sc *reader.SourceCode) (stmts []Statement, err *ParserError) {

  RE_ARG3  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s+(\d+)\s+(\d+)\s*(//.*)?$`)
  RE_ARG2  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s+(\d+)\s*(//.*)?$`)
  RE_ARG1  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s*(//.*)?$`)
  RE_BLANK := regexp.MustCompile(`^\s*(//.*)?$`)

  for i, s := range sc.Line {
    lineno := i + 1

    var stmt Statement
    switch {
    case RE_ARG3.MatchString(s):
      ma := RE_ARG3.FindStringSubmatch(s)
      switch ma[1] {
      case "NAND":
        stmt = NandStatement{
          InputA: ma[2],
          InputB: ma[3],
          Output: ma[4],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown keyword `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_ARG2.MatchString(s):
      ma := RE_ARG2.FindStringSubmatch(s)
      switch ma[1] {
      case "CONNECT":
        stmt = ConnectStatement{
          Input:  ma[2],
          Output: ma[3],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown keyword `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_ARG1.MatchString(s):
      ma := RE_ARG1.FindStringSubmatch(s)
      switch ma[1] {
      case "LO":
        stmt = LOStatement{
          Output: ma[2],
          LineNo: lineno,
        }
      case "HI":
        stmt = HIStatement{
          Output: ma[2],
          LineNo: lineno,
        }
      case "CLOCK":
        stmt = CLOCKStatement{
          Output: ma[2],
          LineNo: lineno,
        }
      case "HALT":
        stmt = HALTStatement{
          Input:  ma[2],
          LineNo: lineno,
        }
      case "LED":
        stmt = LEDStatement{
          Input:  ma[2],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown define type `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_BLANK.MatchString(s):
      continue

    default:
      err = &ParserError{
        Message: fmt.Sprintf("line %d, unknown statement `%s`", lineno, s),
      }
      return stmts, err
    }

    stmts = append(stmts, stmt)
  }

  return stmts, nil
}


