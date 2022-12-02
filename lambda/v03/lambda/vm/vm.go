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
  Closure Closure
}

type Application struct {
  Left  Value
  Right Value
}

type Delay struct {
  Code []Instruction //left, right, Apply
}

type Closure struct {
  Arg  string
  Env  Environment
  Code []Instruction //code, Return
}

type Dump struct {
  Code []Instruction
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

func (d Delay) String() string {
  return "<delay>"
}

func (c Closure) String() string {
  return fmt.Sprintf("<closure e=%v>", c.Env)
}

func (d Dump) String() string {
  return fmt.Sprintf("<dump e=%v>", d.Env)
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


func (vm *VM) Run() Value {

  for {
    vm.debugPrint()

    statement := vm.Next()

    if statement == nil {

      rest := vm.PopStack()

      if rest != nil {
        switch v := rest.(type) {

        case Delay:

          codecp := make([]Instruction, len(v.Code))
          copy(codecp, v.Code)
          vm.code = codecp
          continue

        case Closure:

          // push dump
          envcp := make(Environment, len(vm.env))
          for k, v := range vm.env {
            envcp[k] = v
          }
          codecp := make([]Instruction, len(vm.code))
          copy(codecp, vm.code)
          codecp = append([]Instruction{ Wrap{ Closure: v }, }, codecp...)
          vm.PushStack( Dump { Env: envcp, Code: codecp } )

          // extend code
          vm.env  = v.Env
          vm.code = v.Code
          //delete(vm.env, v.Arg)
          continue

        }
      }

      vm.PushStack(rest)
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

      if function, ok := left.(Function) ; ok {
        left = function.Closure
      }

      if closure, ok := left.(Closure) ; ok {

        // push dump
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Instruction, len(vm.code))
        copy(codecp, vm.code)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend env
        vm.env  = closure.Env
        vm.code = closure.Code
        vm.env[closure.Arg] = right

      } else if delay, ok := left.(Delay) ; ok {
        vm.PushStack( right )
        vm.code = append( append( delay.Code, LRApply{} ), vm.code... )

      } else if delay, ok := right.(Delay) ; ok {
        vm.PushStack( left )
        vm.code = append( append( delay.Code, RLApply{} ), vm.code... )

      } else {
        vm.PushStack( Application{ Left: left, Right: right } )
      }

    case Call:
      vm.PushStack( Delay{ Code: v.Code } )

    case Close:
      envcp := make(Environment, len(vm.env))
      for k, v := range vm.env {
        envcp[k] = v
      }
      vm.PushStack( Closure{ Arg: v.Arg, Env: envcp, Code: v.Code } )

    case Return:
      result := vm.PopStack()
      d      := vm.PopStack()

      if delay, ok := result.(Delay) ; ok {

        vm.PushStack(d)

        vm.code = append( append(delay.Code, Return{}), vm.code...)

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

      if closure, ok := result.(Closure) ; ok {

        // push dump
        envcp := make(Environment, len(vm.env))
        for k, v := range vm.env {
          envcp[k] = v
        }
        codecp := make([]Instruction, len(vm.code))
        copy(codecp, vm.code)
        codecp = append([]Instruction{ Wrap{ Closure: closure }, Wrap{ Closure: v.Closure }, }, codecp...)
        vm.PushStack( Dump { Env: envcp, Code: codecp } )

        // extend code
        vm.env  = closure.Env
        vm.code = closure.Code
        //delete(vm.env, closure.Arg)

      } else {
        vm.PushStack(Function{ Arg: v.Closure.Arg, Body: result, Closure: v.Closure })
      }

    default:
      panic("vm.run: unknown statement")
    }

  }

  vm.debugPrint()

  return vm.PopStack()
}


