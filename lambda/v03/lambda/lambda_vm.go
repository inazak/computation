package lambda

import (
  "strings"
)


type Environment map[string]Expression

// ***** AST and StackItem *****

type Expression interface {
  String() string
}

type Symbol struct {
  Name string
}

type Function struct {
  Arg  string
  Body Expression
}

type Application struct {
  Left  Expression
  Right Expression
}

type Closure struct {
  Arg  string
  Env  Environment
  Code []Statement
}

type Dump struct {
  Code []Statement
  Env Environment
}

func (s Symbol) String() string {
  return s.Name
}

func (f Function) String() string {
  return "^" + f.Arg + "." + f.Body.String()
}

func (a Application) String() string {
  return "(" + a.Left.String() + " " + a.Right.String() + ")"
}

func (c Closure) String() string {
  return "<closure>"
}

func (d Dump) String() string {
  return "<dump>"
}


// ***** Statement *****

type Statement interface {
  String() string
}

type Fetch struct {
  Name string
}

type Close struct {
  Arg  string
  Code []Statement
}

type Apply struct {}
type Return struct {}

type Wrap struct {
  Arg string
}

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

func (a Apply) String() string {
  return "Apply"
}

func (r Return) String() string {
  return "Return"
}

func (w Wrap) String() string {
  return "Wrap " + w.Arg
}


// ***** Compile *****

func Compile(expr Expression) []Statement {

  switch v := expr.(type) {
  case Symbol:
    return []Statement{ Fetch{ Name: v.Name }, }

  case Application:
    left  := Compile(v.Left)
    right := Compile(v.Right)
    return append(append(left, right...), Apply{})

  case Function:
    return []Statement{
      Close{
        Arg: v.Arg,
        Code: append(Compile(v.Body), Return{} ),
      },
    }

  default:
    panic("compile: unknown expression")
  }

}


// ***** Machine *****

type VM struct {
  stack []Expression
  env   Environment
  code  []Statement
}

func NewVM(env Environment, code []Statement) *VM {
  return &VM {
    stack: []Expression{},
    env:   env,
    code:  code,
  }
}

func (vm *VM) PushStack(item Expression) {
  vm.stack = append(vm.stack, item)
}

func (vm *VM) PopStack() Expression {
  if len(vm.stack) == 0 {
    return nil
  }
  item := vm.stack[len(vm.stack)-1]
  vm.stack = vm.stack[:len(vm.stack)-1]
  return item
}

func (vm *VM) Next() Statement {
  if len(vm.code) == 0 {
    return nil
  }
  statement := vm.code[0]
  vm.code = vm.code[1:]
  return statement
}

func (vm *VM) GetEnv(name string) (expr Expression, ok bool) {
  expr, ok = vm.env[name]
  return expr, ok
}

func (vm *VM) SetEnv(name string, expr Expression) {
  vm.env[name] = expr
}


func (vm *VM) Run() Expression {

  for {
    statement := vm.Next()
    if statement == nil {

      result := vm.PopStack()
      if closure, ok := result.(Closure) ; ok {

        // push dump
        // but env is not used and codecp is empty
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Statement, len(vm.code))
        copy(codecp, vm.code)
        codecp = append([]Statement{ Wrap{ Arg: closure.Arg }, }, codecp...)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend code
        vm.code = closure.Code

        statement = vm.Next()

      } else {
        vm.PushStack(result)
        break
      }
    }

    switch v := statement.(type) {
    case Fetch:
      expr, ok := vm.GetEnv(v.Name)
      if ok {
        vm.PushStack( expr )
      } else {
        vm.PushStack( Symbol { Name: v.Name } )
      }

    case Apply:
      right := vm.PopStack()
      left  := vm.PopStack()

      if closure, ok := left.(Closure) ; ok {

        // push dump
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Statement, len(vm.code))
        copy(codecp, vm.code)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend env
        vm.env  = closure.Env
        vm.code = closure.Code
        vm.env[closure.Arg] = right

      } else if closure, ok := right.(Closure) ; ok {

        vm.PushStack(left)

        // push dump
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Statement, len(vm.code))
        copy(codecp, vm.code)
        codecp = append([]Statement{ Wrap{ Arg: closure.Arg }, Apply{} }, codecp...)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend code
        vm.code = closure.Code

      } else {
        vm.PushStack( Application{ Left: left, Right: right } )
      }

    case Close:
      envcp := make(Environment, len(vm.env))
      for k, v := range vm.env {
        envcp[k] = v
      }
      vm.PushStack( Closure{ Arg: v.Arg, Env: envcp, Code: v.Code } )

    case Return:
      result := vm.PopStack()
      d      := vm.PopStack()
      if dump, ok := d.(Dump); ok {
        vm.code = dump.Code
        vm.env  = dump.Env
        vm.PushStack(result)
      } else {
        panic("vm.run: lost dump in return statement")
      }

    case Wrap:
      result := vm.PopStack()
      vm.PushStack(Function{ Arg: v.Arg, Body: result })

    default:
      panic("vm.run: unknown statement")
    }

  }

  return vm.PopStack()
}


