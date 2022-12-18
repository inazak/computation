package vm

import (
  "strings"
  "github.com/inazak/computation/lambda/v04/lambda/ast"
)


type Instruction interface {
  String() string
}

type Fetch struct {
  Name string
}

type Close struct {
  Arg  string
  Code []Instruction
}

type Call struct {
  Code []Instruction
}

type Apply struct {}
type RLApply struct {}
type LRApply struct {}

type Return struct {}

type Wrap struct {}

func (f Fetch) String() string {
  return "Fetch " + f.Name
}

func (c Close) String() string {
  list := []string{}
  for _, code := range c.Code {
    list = append(list, code.String())
  }
  return "Close " + c.Arg + ", [" + strings.Join(list, "; ") + "]"
}

func (c Call) String() string {
  list := []string{}
  for _, code := range c.Code {
    list = append(list, code.String())
  }
  return "Call [" + strings.Join(list, "; ") + "]"
}


func (a Apply) String() string {
  return "Apply"
}

func (a RLApply) String() string {
  return "RLApply"
}

func (a LRApply) String() string {
  return "LRApply"
}

func (r Return) String() string {
  return "Return"
}

func (w Wrap) String() string {
  return "Wrap"
}


// ***** Compile *****

func Compile(expr ast.Expression) []Instruction {

  switch v := expr.(type) {
  case ast.Symbol:
    return []Instruction{ Fetch{ Name: v.Name }, }

  case ast.Application:
    left  := Compile(v.Left)
    right := Compile(v.Right)
    return []Instruction{
      Call{
        Code: append(append(left, right...), Apply{}, Return{} ),
      },
    }

  case ast.Function:
    return []Instruction{
      Close{
        Arg: v.Arg,
        Code: append(Compile(v.Body), Return{} ),
      },
    }

  default:
    panic("compile: unknown expression")
  }
}

