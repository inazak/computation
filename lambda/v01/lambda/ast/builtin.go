package ast

var BUILTINS = []struct{
  Name string
  Expr Expression
}{
  { Name: "True",       Expr: VAR_TRUE },
  { Name: "False",      Expr: VAR_FALSE },
  { Name: "If",         Expr: VAR_IF },
  { Name: "IsEmpty",    Expr: VAR_ISEMPTY },
  { Name: "IsZero",     Expr: VAR_ISZERO },
  { Name: "LessOrEq",   Expr: VAR_LESSOREQ },
  { Name: "Succ",       Expr: VAR_SUCC },
  { Name: "Pred",       Expr: VAR_PRED },
  { Name: "Add",        Expr: VAR_ADD },
  { Name: "Sub",        Expr: VAR_SUB },
  { Name: "Div",        Expr: VAR_DIV },
  { Name: "Mod",        Expr: VAR_MOD },
  { Name: "Pair",       Expr: VAR_PAIR },
  { Name: "Left",       Expr: VAR_LEFT },
  { Name: "Right",      Expr: VAR_RIGHT },
  { Name: "Slide",      Expr: VAR_SLIDE },
  { Name: "Empty",      Expr: VAR_EMPTY },
  { Name: "First",      Expr: VAR_FIRST },
  { Name: "Rest",       Expr: VAR_REST },
  { Name: "Unshift",    Expr: VAR_UNSHIFT },
  { Name: "Range",      Expr: VAR_RANGE },
  { Name: "Fold",       Expr: VAR_FOLD },
  { Name: "Map",        Expr: VAR_MAP },
  { Name: "Push",       Expr: VAR_PUSH },
  { Name: "Y",          Expr: VAR_Y },
  { Name: "DigitToStr", Expr: VAR_DIGITTOSTR },
  { Name: "FizzBuzz",   Expr: VAR_FIZZBUZZ },
}


func MakeVariable(name string, expr Expression) *Variable {
  return &Variable { Name: name, Expr: expr }
}

var VAR_TRUE =
  &Function { Arg: 'x', Body:
  &Function { Arg: 'y', Body:
  &Symbol { Name: 'x' } } }

var VAR_FALSE =
  &Function { Arg: 'x', Body:
  &Function { Arg: 'y', Body:
  &Symbol { Name: 'y' } } }

var VAR_IF =
  &Function { Arg: 'b', Body:
  &Symbol { Name: 'b' } }

var VAR_ISEMPTY =
  VAR_LEFT

var VAR_ISZERO =
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'x' }, Right:
  &Application { Left: MakeVariable("True", VAR_TRUE),
                 Right: MakeVariable("False", VAR_FALSE), } },
  Right: MakeVariable("True", VAR_TRUE) } }

var VAR_LESSOREQ =
  &Function { Arg: 'x', Body:
  &Function { Arg: 'y', Body:
  &Application { Left: MakeVariable("IsZero", VAR_ISZERO), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Sub", VAR_SUB), Right:
  &Symbol { Name: 'x' } }, Right:
  &Symbol { Name: 'y' } } } } }

var VAR_SUCC =
  &Function { Arg: 'n', Body:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'n' }, Right:
  &Symbol { Name: 'f' } }, Right:
  &Symbol { Name: 'x' } } } } } }

var VAR_PRED =
  &Function { Arg: 'n', Body:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'n' }, Right:
  &Function { Arg: 'g', Body:
  &Function { Arg: 'h', Body:
  &Application { Left:
  &Symbol { Name: 'h' }, Right:
  &Application { Left:
  &Symbol { Name: 'g' }, Right:
  &Symbol { Name: 'f' } } } } } }, Right:
  &Function { Arg: 'u', Body:
  &Symbol { Name: 'x' } } }, Right:
  &Function { Arg: 'u', Body:
  &Symbol { Name: 'u' } } } } } }

var VAR_ADD =
  &Function { Arg: 'x', Body:
  &Function { Arg: 'y', Body:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'x' }, Right: MakeVariable("Succ", VAR_SUCC), }, Right:
  &Symbol { Name: 'y' } } } }

var VAR_SUB =
  &Function { Arg: 'x', Body:
  &Function { Arg: 'y', Body:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'y' }, Right: MakeVariable("Pred", VAR_PRED), }, Right:
  &Symbol { Name: 'x' } } } }

var VAR_DIV =
  &Application { Left: MakeVariable("Y", VAR_Y), Right:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'm', Body:
  &Function { Arg: 'n', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left:
  &Application { Left: MakeVariable("LessOrEq", VAR_LESSOREQ), Right:
  &Symbol { Name: 'n' } }, Right:
  &Symbol { Name: 'm' } } }, Right:
  &Application { Left: MakeVariable("Succ", VAR_SUCC), Right:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Sub", VAR_SUB), Right:
  &Symbol { Name: 'm' } }, Right:
  &Symbol { Name: 'n' } } }, Right:
  &Symbol { Name: 'n' } } } }, Right:
  &Number { Name: "0" } } } } } }

var VAR_MOD =
  &Application { Left: MakeVariable("Y", VAR_Y), Right:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'm', Body:
  &Function { Arg: 'n', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left:
  &Application { Left: MakeVariable("LessOrEq", VAR_LESSOREQ), Right:
  &Symbol { Name: 'n' } }, Right:
  &Symbol { Name: 'm' } } }, Right:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Sub", VAR_SUB), Right:
  &Symbol { Name: 'm' } }, Right:
  &Symbol { Name: 'n' } } }, Right:
  &Symbol { Name: 'n' } } }, Right:
  &Symbol { Name: 'm' } } } } } }

var VAR_PAIR =
  &Function { Arg: 'a', Body:
  &Function { Arg: 'b', Body:
  &Function { Arg: 'f', Body:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'f' } , Right:
  &Symbol { Name: 'a' }, } , Right:
  &Symbol { Name: 'b' }, } } } }

var VAR_LEFT =
  &Function { Arg: 'p', Body:
  &Application { Left:
  &Symbol { Name: 'p' } , Right:
  &Function { Arg: 'a', Body:
  &Function { Arg: 'b', Body:
  &Symbol { Name: 'a' } } } }, }

var VAR_RIGHT =
  &Function { Arg: 'p', Body:
  &Application { Left:
  &Symbol { Name: 'p' } , Right:
  &Function { Arg: 'a', Body:
  &Function { Arg: 'b', Body:
  &Symbol { Name: 'b' } } } }, }

var VAR_SLIDE =
  &Function { Arg: 'p', Body:
  &Pair { Left:
  &Application { Left: MakeVariable("Right", VAR_RIGHT), Right:
  &Symbol { Name: 'p' } }, Right:
  &Application { Left: MakeVariable("Succ", VAR_SUCC), Right:
  &Application { Left: MakeVariable("Right", VAR_RIGHT), Right:
  &Symbol { Name: 'p' } } } } }

var VAR_EMPTY =
  &Pair { Left:  MakeVariable("True", VAR_TRUE),
          Right: MakeVariable("True", VAR_TRUE) }

var VAR_FIRST =
  &Function { Arg: 'l', Body:
  &Application { Left: MakeVariable("Left", VAR_LEFT), Right:
  &Application { Left: MakeVariable("Right", VAR_RIGHT), Right:
  &Symbol { Name: 'l' } } } }

var VAR_REST =
  &Function { Arg: 'l', Body:
  &Application { Left: MakeVariable("Right", VAR_RIGHT), Right:
  &Application { Left: MakeVariable("Right", VAR_RIGHT), Right:
  &Symbol { Name: 'l' } } } }

var VAR_UNSHIFT =
  &Function { Arg: 'l', Body:
  &Function { Arg: 'x', Body:
  &Pair { Left: MakeVariable("False", VAR_FALSE), Right:
  &Pair { Left:
  &Symbol { Name: 'x' }, Right:
  &Symbol { Name: 'l' } } } } }

var VAR_RANGE =
  &Application { Left: MakeVariable("Y", VAR_Y), Right:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'm', Body:
  &Function { Arg: 'n', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left:
  &Application { Left: MakeVariable("LessOrEq", VAR_LESSOREQ), Right:
  &Symbol { Name: 'm' } }, Right:
  &Symbol { Name: 'n' } } }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Unshift", VAR_UNSHIFT), Right:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left: MakeVariable("Succ", VAR_SUCC), Right:
  &Symbol { Name: 'm' } } }, Right:
  &Symbol { Name: 'n' } } }, Right:
  &Symbol { Name: 'm' } } }, Right: MakeVariable("Empty", VAR_EMPTY), } } } } }

var VAR_FOLD =
  &Application { Left: MakeVariable("Y", VAR_Y), Right:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'l', Body:
  &Function { Arg: 'x', Body:
  &Function { Arg: 'g', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left: MakeVariable("IsEmpty", VAR_ISEMPTY), Right:
  &Symbol { Name: 'l' } } }, Right:
  &Symbol { Name: 'x' } }, Right:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'g' }, Right:
  &Application { Left:
  &Application { Left:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left: MakeVariable("Rest", VAR_REST), Right:
  &Symbol { Name: 'l' } } }, Right:
  &Symbol { Name: 'x' } }, Right:
  &Symbol { Name: 'g' } } }, Right:
  &Application { Left: MakeVariable("First", VAR_FIRST), Right:
  &Symbol { Name: 'l' } } } } } } } } }

var VAR_MAP =
  &Function { Arg: 'k', Body:
  &Function { Arg: 'f', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("Fold", VAR_FOLD), Right:
  &Symbol { Name: 'k' } }, Right: MakeVariable("Empty", VAR_EMPTY) }, Right:
  &Function { Arg: 'l', Body:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Application { Left: MakeVariable("Unshift", VAR_UNSHIFT), Right:
  &Symbol { Name: 'l' } }, Right:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Symbol { Name: 'x' } } } } } } } }

var VAR_PUSH =
  &Function { Arg: 'l', Body:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("Fold", VAR_FOLD), Right:
  &Symbol { Name: 'l' } }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Unshift", VAR_UNSHIFT),
                 Right: MakeVariable("Empty", VAR_EMPTY) }, Right:
  &Symbol { Name: 'x' } } }, Right: MakeVariable("Unshift", VAR_UNSHIFT) } } }

var VAR_Y =
  &Function { Arg: 'f', Body:
  &Application { Left:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Symbol { Name: 'x' }, Right:
  &Symbol { Name: 'x' } } } }, Right:
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Symbol { Name: 'x' }, Right:
  &Symbol { Name: 'x' } } } } } }

var VAR_DIGITTOSTR =
  &Application { Left: MakeVariable("Y", VAR_Y), Right:
  &Function { Arg: 'f', Body:
  &Function { Arg: 'n', Body:
  &Application { Left:
  &Application { Left: MakeVariable("Push", VAR_PUSH), Right:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left:
  &Application { Left: MakeVariable("LessOrEq", VAR_LESSOREQ), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "9" } } }, Right: MakeVariable("Empty", VAR_EMPTY) }, Right:
  &Application { Left:
  &Symbol { Name: 'f' }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Div", VAR_DIV), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "10" } } } } }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Add", VAR_ADD), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Mod", VAR_MOD), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "10" } } }, Right:
  &Number { Name: "48" } } } } } }

var VAR_FIZZBUZZ =
  &Function { Arg: 'x', Body:
  &Application { Left:
  &Application { Left: MakeVariable("Map", VAR_MAP), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Range", VAR_RANGE), Right:
  &Number { Name: "1" } }, Right:
  &Symbol { Name: 'x' } } }, Right:
  &Function { Arg: 'n', Body:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left: MakeVariable("IsZero", VAR_ISZERO), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Mod", VAR_MOD), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "15" } } } }, Right:
  &Str { S: "FizzBuzz\n" } }, Right:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left: MakeVariable("IsZero", VAR_ISZERO), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Mod", VAR_MOD), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "3" } } } }, Right:
  &Str { S: "Fizz\n" } }, Right:
  &Application { Left:
  &Application { Left:
  &Application { Left: MakeVariable("If", VAR_IF), Right:
  &Application { Left: MakeVariable("IsZero", VAR_ISZERO), Right:
  &Application { Left:
  &Application { Left: MakeVariable("Mod", VAR_MOD), Right:
  &Symbol { Name: 'n' } }, Right:
  &Number { Name: "5" } } } }, Right:
  &Str { S: "Buzz\n" } }, Right:
  &Application { Left:
  &Application { Left: MakeVariable("Push", VAR_PUSH), Right:
  &Application { Left: MakeVariable("DigitToStr", VAR_DIGITTOSTR), Right:
  &Symbol { Name: 'n' } } }, Right:
  &Number { Name: "10" } } } } } } } }

