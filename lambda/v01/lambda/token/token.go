package token

type TokenType string

const (
  SYMBOL   = "SYMBOL"
  VARIABLE = "VARIABLE"
  LAMBDA   = "LAMBDA"
  DOT      = "DOT"
  LPAREN   = "LPAREN"
  RPAREN   = "RPAREN"
  LBRACKET = "LBRACKET"
  RBRACKET = "RBRACKET"
  LANGLE   = "LANGLE"
  RANGLE   = "RANGLE"
  EQUAL    = "EQUAL"
  CHAR     = "CHAR"
  NUMBER   = "NUMBER"
  STR      = "STR"
  EOL      = "EOL"
  UNKNOWN  = "UNKNOWN"
)


type Token struct {
  Type    TokenType
  Literal string
}

func NewToken(t TokenType, s string) Token {
  return Token{ Type: t, Literal: s }
}

