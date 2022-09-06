package parser

import (
  "fmt"
  "github.com/inazak/computation/lambda/v01/lambda/token"
  "github.com/inazak/computation/lambda/v01/lambda/lexer"
  "github.com/inazak/computation/lambda/v01/lambda/ast"
)

type Parser struct {
  l         *lexer.Lexer
  currToken token.Token
  nextToken token.Token
  errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
  p := &Parser{ l: l, errors: nil }
  p.readToken() // set p.nextToken
  p.readToken() // set p.nextToken and p.currToken
  return p
}

func (p *Parser) readToken() {
  p.currToken = p.nextToken
  p.nextToken = p.l.NextToken()

  if p.currToken.Type == token.UNKNOWN {
    p.errors = append(p.errors, fmt.Sprintf("Lexer got unknown token:"))
    p.errors = append(p.errors, fmt.Sprintf("  %s", p.l.GetMsg()))
  }
}

func (p *Parser) AddErrorMessage(process string, expect string) {
  p.errors = append(p.errors, fmt.Sprintf("Error at Parsing %s:", process))
  p.errors = append(p.errors, fmt.Sprintf("  reading text '%s' <-", shorten(p.l.LookbackText(), 10)))
  p.errors = append(p.errors, fmt.Sprintf("  parser expect %s,", expect))
  p.errors = append(p.errors, fmt.Sprintf("  but got %s '%s'", p.currToken.Type, p.currToken.Literal))
}

func shorten(s string, i int) string {
  if len(s) <= i {
    return s
  } else {
    return "..." + s[len(s)-i:len(s)]
  }
}

func (p *Parser) GetError() []string {
  return p.errors
}

func (p *Parser) IsTokenType(expect token.TokenType) bool {
  return p.currToken.Type == expect
}

func (p *Parser) PopToken() token.Token {
  tk := p.currToken
  p.readToken()
  return tk
}

func (p *Parser) ConsumeToken() {
  _ = p.PopToken()
}


// implementation of Parse

func (p *Parser) Parse() ast.Statement {
  stmt := p.ParseStatement()

  if p.GetError() != nil {
    return stmt
  }

  if ! p.IsTokenType(token.EOL) {
    p.AddErrorMessage("Statement", "end of line")
    return stmt
  }

  p.ConsumeToken()
  return stmt
}

func (p *Parser) ParseStatement() ast.Statement {
  var stmt ast.Statement

  switch p.currToken.Type {
  case token.VARIABLE:
    if p.nextToken.Type == token.EQUAL {
      stmt = p.ParseAssignmentStatement()
    } else {
      stmt = p.ParseExpressionStatement()
    }
  default:
    stmt = p.ParseExpressionStatement()
  }

  return stmt
}


func (p *Parser) ParseAssignmentStatement() ast.Statement {
  stmt := &ast.AssignmentStatement{}
  stmt.Var = &ast.Variable { Expr: &ast.Blank{} }

  if ! p.IsTokenType(token.VARIABLE) {
    p.AddErrorMessage("AssignmentStatement", "variable name")
    return stmt
  }

  tk := p.PopToken()
  stmt.Var.Name = tk.Literal

  if ! p.IsTokenType(token.EQUAL) {
    p.AddErrorMessage("AssignmentStatement", "equal '=' char")
    return stmt
  }
  p.ConsumeToken()

  stmt.Var.Expr = p.ParseExpression()
  return stmt
}


func (p *Parser) ParseExpressionStatement() ast.Statement {
  stmt := &ast.ExpressionStatement{}
  stmt.Expr = p.ParseExpression()
  return stmt
}


func (p *Parser) ParseExpression() ast.Expression {
  var expr ast.Expression

  switch p.currToken.Type {
  case token.LAMBDA:
    expr = p.ParseFunction()
  case token.LPAREN:
    expr = p.ParseApplication()
  case token.LANGLE:
    expr = p.ParsePair()
  case token.LBRACKET:
    expr = p.ParseList()
  case token.SYMBOL:
    expr = p.ParseSymbol()
  case token.VARIABLE:
    expr = p.ParseVariable()
  case token.NUMBER:
    expr = p.ParseNumber()
  case token.CHAR:
    expr = p.ParseChar()
  case token.STR:
    expr = p.ParseStr()
  default:
    p.AddErrorMessage("Expression", "one of the allowed char")
  }
  return expr
}


func (p *Parser) ParseFunction() *ast.Function {
  f := &ast.Function{}

  if ! p.IsTokenType(token.LAMBDA) {
    p.AddErrorMessage("Function", "lambda '^' char")
    return f
  }
  p.ConsumeToken()

  tk := p.PopToken()
  if len(tk.Literal) != 1 {
    p.AddErrorMessage("Function", "symbol character")
    return f
  }
  f.Arg = byte(tk.Literal[0])

  if ! p.IsTokenType(token.DOT) {
    p.AddErrorMessage("Function", "dot '.' char")
    return f
  }
  p.ConsumeToken()

  f.Body = p.ParseExpression()
  return f
}

func (p *Parser) ParseApplication() *ast.Application {
  a := &ast.Application{}

  if ! p.IsTokenType(token.LPAREN) {
    p.AddErrorMessage("Application", "left paren '(' char")
    return a
  }
  p.ConsumeToken()

  a.Left  = p.ParseExpression()
  a.Right = p.ParseExpression()

  if ! p.IsTokenType(token.RPAREN) {
    p.AddErrorMessage("Application", "right paren ')' char")
    return a
  }
  p.ConsumeToken()

  return a
}

func (p *Parser) ParsePair() *ast.Pair {
  a := &ast.Pair{}

  if ! p.IsTokenType(token.LANGLE) {
    p.AddErrorMessage("Pair", "left angle bracket '<' char")
    return a
  }
  p.ConsumeToken()

  a.Left  = p.ParseExpression()
  a.Right = p.ParseExpression()

  if ! p.IsTokenType(token.RANGLE) {
    p.AddErrorMessage("Pair", "right angle bracket '>' char")
    return a
  }
  p.ConsumeToken()

  return a
}

func (p *Parser) ParseList() *ast.List {
  l := &ast.List{ Data: []ast.Expression{} }

  if ! p.IsTokenType(token.LBRACKET) {
    p.AddErrorMessage("List", "left bracket '[' char")
    return l
  }
  p.ConsumeToken()

  for ! p.IsTokenType(token.RBRACKET) {
    if p.IsTokenType(token.EOL) {
      p.AddErrorMessage("List", "right bracket ']' char")
      return l
    }

    l.Data = append(l.Data, p.ParseExpression())
  }

  p.ConsumeToken() //token.RBRACKET
  return l
}

func (p *Parser) ParseSymbol() *ast.Symbol {
  s := &ast.Symbol{}

  if ! p.IsTokenType(token.SYMBOL) {
    p.AddErrorMessage("Symbol", "symbol character")
    return s
  }

  tk := p.PopToken()
  if len(tk.Literal) != 1 {
    p.AddErrorMessage("Symbol", "symbol character")
    return s
  }
  s.Name = byte(tk.Literal[0])

  return s
}

func (p *Parser) ParseVariable() *ast.Variable {
  v := &ast.Variable{ Expr: &ast.Blank{} }

  if ! p.IsTokenType(token.VARIABLE) {
    p.AddErrorMessage("Variable", "variable name")
    return v
  }

  tk := p.PopToken()
  v.Name = tk.Literal

  return v
}

func (p *Parser) ParseNumber() *ast.Number {
  n := &ast.Number{}

  if ! p.IsTokenType(token.NUMBER) {
    p.AddErrorMessage("Number", "number character")
    return n
  }

  tk := p.PopToken()
  n.Name = tk.Literal

  return n
}

func (p *Parser) ParseChar() *ast.Char {
  c := &ast.Char{}

  if ! p.IsTokenType(token.CHAR) {
    p.AddErrorMessage("Char", "one byte character")
    return c
  }

  tk := p.PopToken()
  c.C = tk.Literal[0]

  return c
}

func (p *Parser) ParseStr() *ast.Str {
  s := &ast.Str{}

  if ! p.IsTokenType(token.STR) {
    p.AddErrorMessage("Str", "any printable character")
    return s
  }

  tk := p.PopToken()
  s.S = tk.Literal

  return s
}

