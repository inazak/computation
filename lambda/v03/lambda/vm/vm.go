package vm

import (
  "log"
  "os"
  "strings"
  "github.com/inazak/computation/lambda/v03/lambda/ast"
)


type Environment map[string]Value

// ***** Value *****

type Value interface {
  String() string
}

type Symbol struct {
  Name string
}

type Function struct {
  Arg  string
  Body Value
  Closure Closure
}

type Application struct {
  Left  Value
  Right Value
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
  Closure Closure
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

func Compile(expr ast.Expression) []Statement {

  switch v := expr.(type) {
  case ast.Symbol:
    return []Statement{ Fetch{ Name: v.Name }, }

  case ast.Application:
    left  := Compile(v.Left)
    right := Compile(v.Right)
    return append(append(left, right...), Apply{})

  case ast.Function:
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
  stack []Value
  env   Environment
  code  []Statement
  logger *log.Logger
}

func NewVM(env Environment, code []Statement) *VM {
  return &VM {
    stack: []Value{},
    env:   env,
    code:  code,
  }
}

func (vm *VM) logf(format string, v ...interface{}) {
  if vm.logger != nil {
    vm.logger.Printf(format, v...)
  }
}

func (vm *VM) EnableLogging() {
  vm.logger = log.New(os.Stderr, "", log.LstdFlags)
}

func (vm *VM) DisableLogging() {
  vm.logger = nil
}


func (vm *VM) PushStack(item Value) {
  vm.stack = append(vm.stack, item)
}

func (vm *VM) PopStack() Value {
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

func (vm *VM) GetEnv(name string) (expr Value, ok bool) {
  expr, ok = vm.env[name]
  return expr, ok
}

func (vm *VM) SetEnv(name string, expr Value) {
  vm.env[name] = expr
}


func (vm *VM) Run() Value {

  for {
    statement := vm.Next()

    if statement == nil {

      vm.logf("[debug] code is empty, exit vm.Run()\n")

      break
    }

    vm.logf("[debug] code :%s\n", statement)

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

      if function, ok := left.(Function) ; ok {
        left = function.Closure
      }

      if closure, ok := left.(Closure) ; ok {

        vm.logf("[debug] apply left is closure\n")

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

        vm.logf("[debug] apply right is closure\n")

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

      //add Wrap
      if closure, ok := result.(Closure) ; ok {

        vm.logf("[debug] stack top is closure\n")
        vm.logf("[debug] add wrap, dump, and expand closure\n")

        vm.PushStack(d)

        // push dump
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Statement, len(vm.code))
        copy(codecp, vm.code)
        codecp = append([]Statement{ Wrap{ Arg: closure.Arg, Closure: closure }, Return{}, }, codecp...)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend code
        vm.code = closure.Code

      } else {

        vm.logf("[debug] stack top is NOT closure\n")

        vm.PushStack(result)

        if dump, ok := d.(Dump); ok {
          vm.code = dump.Code
          vm.env  = dump.Env
        } else {
          panic("vm.run: lost dump in return statement")
        }

      }

    case Wrap:
      result := vm.PopStack()
      vm.PushStack(Function{ Arg: v.Arg, Body: result, Closure: v.Closure })


    default:
      panic("vm.run: unknown statement")
    }

    vm.logf("[debug] env  :%s\n", vm.env)
    vm.logf("[debug] stack:%s\n", vm.stack)
    vm.logf("[debug] ---\n")

  }

  return vm.PopStack()
}


