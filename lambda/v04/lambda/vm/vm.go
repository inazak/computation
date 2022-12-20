package vm

import (
  "fmt"
  "log"
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
}

type Application struct {
  Left  Value
  Right Value
}

type Callure struct {
  Env  Environment
  Code []Instruction //left, right, Apply, Return
}

type Closure struct {
  Arg  string
  Env  Environment
  Code []Instruction //code, Return
}

type Dump struct {
  Env Environment
  Code []Instruction
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

func (c Callure) String() string {
  return fmt.Sprintf("<callure e=%v>",c.Env)
}

func (c Closure) String() string {
  return fmt.Sprintf("<closure e=%v>",c.Env)
}

func (d Dump) String() string {
  return fmt.Sprintf("<dump e=%v c=%v>", d.Env, d.Code)
}


// ***** Machine *****

type VM struct {
  stack []Value
  env   Environment
  code  []Instruction
  logger *log.Logger
}

func NewVM(env Environment, code []Instruction) *VM {
  return &VM {
    stack: []Value{},
    env:   env,
    code:  code,
  }
}

func (vm *VM) IsStackEmpty() bool {
  return len(vm.stack) == 0
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

func (vm *VM) Next() Instruction {
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

func (vm *VM) PushDump() {
  //copy environment
  envcp := make(Environment, len(vm.env))
  for k, v := range vm.env {
    envcp[k] = v
  }

  //copy code
  codecp := make([]Instruction, len(vm.code))
  copy(codecp, vm.code)

  //push Dump
  vm.PushStack( Dump { Env: envcp, Code: codecp } )
}

func (vm *VM) InsertInstruction(i Instruction) {
  vm.code = append([]Instruction{ i }, vm.code...)
}

func (vm *VM) InsertInstructions(is []Instruction) {
  vm.code = append(is, vm.code...)
}

func (vm *VM) Run() Value {

  LOOP:
  for {

    vm.debugPrint()

    statement := vm.Next()
    if statement == nil {
      break
    }

    switch v := statement.(type) {
    case Fetch:
      expr, ok := vm.GetEnv(v.Name)
      if ok {
        vm.PushStack( expr )
      } else {
        vm.PushStack( Symbol { Name: v.Name } )
      }

    case Apply, RLApply, LRApply:

      var right, left Value

      switch v.(type) {
      case Apply, RLApply:
        right = vm.PopStack()
        left  = vm.PopStack()
      case LRApply:
        left  = vm.PopStack()
        right = vm.PopStack()
      }

      if closure, ok := left.(Closure) ; ok {
        vm.PushDump()
        vm.env  = closure.Env
        vm.code = closure.Code
        vm.env[closure.Arg] = right

      } else if callure, ok := left.(Callure) ; ok {
        vm.PushStack( right )
        vm.InsertInstruction( LRApply{} )
        vm.PushDump()
        vm.env  = callure.Env
        vm.code = callure.Code

      } else if closure, ok := right.(Closure) ; ok {
        vm.PushStack( left )
        vm.InsertInstruction( RLApply{} )
        vm.PushStack( Symbol { Name: closure.Arg } )
        vm.InsertInstruction( Wrap{} )
        vm.PushDump()
        vm.env  = closure.Env
        vm.code = closure.Code
        delete(vm.env, closure.Arg) //for internal reduction

      } else if callure, ok := right.(Callure) ; ok {
        vm.PushStack( left )
        vm.InsertInstruction( RLApply{} )
        vm.PushDump()
        vm.env  = callure.Env
        vm.code = callure.Code

      } else {
        vm.PushStack( Application{ Left: left, Right: right } )
      }

    case Call:
      //copy env
      envcp := make(Environment, len(vm.env))
      for k, v := range vm.env {
        envcp[k] = v
      }
      vm.PushStack( Callure{ Env: envcp, Code: v.Code } )

    case Close:
      //copy env
      envcp := make(Environment, len(vm.env))
      for k, v := range vm.env {
        envcp[k] = v
      }
      vm.PushStack( Closure{ Arg: v.Arg, Env: envcp, Code: v.Code } )

    case Return:
      result := vm.PopStack()
      d      := vm.PopStack()

      if callure, ok := result.(Callure) ; ok {

        vm.PushStack(d)
        vm.InsertInstruction( Return{} )
        vm.PushDump()
        vm.env  = callure.Env
        vm.code = callure.Code

      } else {

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
      s      := vm.PopStack()
      symbol, _ := s.(Symbol) //for internal reduction

      if closure, ok := result.(Closure) ; ok {

        vm.PushStack( Symbol { Name: symbol.Name } )
        vm.InsertInstruction( Wrap{} )

        vm.PushStack( Symbol { Name: closure.Arg } )
        vm.InsertInstruction( Wrap{} )

        vm.PushDump()
        vm.env  = closure.Env
        vm.code = closure.Code
        delete(vm.env, closure.Arg) //for internal reduction

      } else {
        vm.PushStack(Function{ Arg: symbol.Name, Body: result })
      }


    default:
      panic("vm.run: unknown statement")
    }

  } //for


  if vm.IsStackEmpty() {
    return nil
  }

  rest := vm.PopStack()

  switch v := rest.(type) {

  case Callure:

    vm.PushDump()
    vm.env = v.Env
    vm.code = v.Code
    goto LOOP

  case Closure:

    vm.PushStack( Symbol { Name: v.Arg } )
    vm.InsertInstruction( Wrap{} )
    vm.PushDump()
    vm.env = v.Env
    vm.code = v.Code
    delete(vm.env, v.Arg) //for internal reduction
    goto LOOP

  } //switch

  vm.PushStack(rest)

  vm.debugPrint()

  return vm.PopStack()
}


