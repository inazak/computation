package eval

import (
  "fmt"
  "github.com/inazak/computation/lambda/v01/lambda/ast"
  "github.com/inazak/computation/lambda/v01/lambda/lexer"
  "github.com/inazak/computation/lambda/v01/lambda/parser"
)

var debug_on bool

type Evaluator struct {
  vt           *ast.VarTable
  output_mode  string
  debug_info   []string
}

type Output struct {
  Text      string
  IsError   bool
  IsNothing bool
}

const (
  OUTPUT_LAMBDA = "OUTPUT_LAMBDA"
  OUTPUT_INDEX  = "OUTPUT_INDEX"
  OUTPUT_EXPAND = "OUTPUT_EXPAND"
  OUTPUT_REDUCE = "OUTPUT_REDUCE"
  OUTPUT_ASCII  = "OUTPUT_ASCII"
  OUTPUT_ERROR  = "OUTPUT_ERROR"
)

func SetDebug(b bool) {
  debug_on = b
}

func NewEvaluator(m string) *Evaluator {
  return &Evaluator{
    vt: ast.NewVarTable(),
    output_mode: m,
  }
}

func (ev *Evaluator) AddDebugInfo(m string, arg ...interface{}) {
  if debug_on {
    ev.debug_info = append(ev.debug_info, fmt.Sprintf(m, arg...))
  }
}

func (ev *Evaluator) AddDebugInfoExprToString(m string, arg ...interface{}) {
  if debug_on {
    if len(arg) > 0 {
      expr, _ := arg[len(arg)-1].(ast.Expression)
      arg[len(arg)-1] = expr.String()
      ev.debug_info = append(ev.debug_info, fmt.Sprintf(m, arg...))
    } else {
      ev.AddDebugInfo(m)
    }
  }
}

func (ev *Evaluator) ClearDebugInfo() {
  ev.debug_info = nil
}

func (ev *Evaluator) GetDebugInfo() (bool, []string) {
  if ev.debug_info != nil {
    return true, ev.debug_info
  }
  return false, nil
}

func (ev *Evaluator) SetOutputMode(m string) {
  ev.output_mode = m
}

func (ev *Evaluator) eval(stmt ast.Statement) (expr ast.Expression, out Output) {

  astmt, ok := stmt.(*ast.AssignmentStatement)
  if ok {
    ev.vt.UpdateVariableWithName(astmt.Var.Expr, astmt.Var.Name)
    ev.vt.Set(astmt.Var.Name, astmt.Var.Expr)
    return &ast.Blank{}, Output{ Text: "", IsNothing: true }
  }

  estmt, ok := stmt.(*ast.ExpressionStatement)
  if ! ok {
    return expr, Output{ Text: "Error: unknown statement", IsError: true }
  }

  expr = estmt.Expr
  ev.vt.UpdateVariable(expr)
  expr.Indexing(ast.NewSymbolRef())

  switch ev.output_mode {

  case OUTPUT_LAMBDA:
    return expr, Output{ Text: expr.String() }

  case OUTPUT_INDEX:
    expr = ev.reduce(expr)
    return expr, Output{ Text: expr.StringByIndex() }

  case OUTPUT_EXPAND:
    expr = ev.expand(expr)
    return expr, Output{ Text: expr.String() }

  case OUTPUT_REDUCE:
    expr = ev.reduce(expr)
    return expr, Output{ Text: expr.String() }

  case OUTPUT_ASCII:
    expr := ev.reduce(expr)
    ok, s := ChurchNumberListToString(expr)
    if ok {
      return expr, Output{ Text: s }
    } else {
      return expr, Output{ Text: "Error: output is not church number list", IsError: true }
    }

  default:
    return expr, Output{ Text: "Error: unknown Evaluator.Mode", IsError: true }
  }
}


func (ev *Evaluator) Eval(stmt ast.Statement) Output {
  _, out := ev.eval(stmt)
  return out
}

func (ev *Evaluator) parseAndEval(text string) (perr []string, expr ast.Expression, o Output) {

  p    := parser.NewParser(lexer.NewLexer(text))
  stmt := p.Parse()
  perr = p.GetError()

  if perr != nil { //parse error
    return perr, nil, Output{ Text: "Error: parse error", IsError: true }
  }

  expr, o = ev.eval(stmt)
  return nil, expr, o
}

func (ev *Evaluator) ParseAndEval(text string) (perr []string, o Output) {
  perr, _, o = ev.parseAndEval(text)
  return perr, o
}


func (ev *Evaluator) expand(expr ast.Expression) (ast.Expression) {
  ev.AddDebugInfo("begin Evaluator.expand")
  ev.AddDebugInfoExprToString("expr=%v", expr)

  count := 0
  var ok = true
  for ok {
    ev.vt.UpdateVariable(expr)
    expr.Indexing(ast.NewSymbolRef())
    ok, expr = expr.Expand()

    count += 1
    ev.AddDebugInfoExprToString("%d times expand, expr=%v", count, expr)
  }

  ev.AddDebugInfo("end Evaluator.expand")
  return expr
}

func (ev *Evaluator) reduce(expr ast.Expression) (ast.Expression) {
  ev.AddDebugInfo("begin Evaluator.reduce")
  ev.AddDebugInfoExprToString("expr=%v", expr)

  expr = ev.expand(expr)

  count := 0
  var ok = true
  for ok {
    ok, expr = expr.Reduce()

    count += 1
    ev.AddDebugInfoExprToString("%d times reduce, expr=%v", count, expr)
  }

  ev.AddDebugInfo("end Evaluator.reduce")
  return expr
}


