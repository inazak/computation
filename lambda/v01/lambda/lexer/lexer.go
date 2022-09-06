package lexer

import (
  "github.com/inazak/computation/lambda/v01/lambda/token"
)

type Lexer struct {
  text         string
  currPosition int
  nextPosition int
  c            byte //charactor of currPosition
  msg          string
}

func NewLexer(text string) *Lexer {
  l := &Lexer{ text: text }
  l.read()
  return l
}

func (l *Lexer) LookbackText() string {
  if l.currPosition < 2 {
    return ""
  } else {
    return l.text[:l.currPosition-2]
  }
}

func (l *Lexer) GetMsg() string {
  return l.msg
}

func (l *Lexer) read() {
  if l.nextPosition >= len(l.text) {
    l.c = 0
  } else {
    l.c = l.text[l.nextPosition]
  }
  l.currPosition = l.nextPosition
  l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
  var tk token.Token
  l.skipSpace()

  switch l.c {
    case '^':
      tk = token.NewToken(token.LAMBDA, "^")
    case '.':
      tk = token.NewToken(token.DOT, ".")
    case '(':
      tk = token.NewToken(token.LPAREN, "(")
    case ')':
      tk = token.NewToken(token.RPAREN, ")")
    case '[':
      tk = token.NewToken(token.LBRACKET, "[")
    case ']':
      tk = token.NewToken(token.RBRACKET, "]")
    case '<':
      tk = token.NewToken(token.LANGLE, "<")
    case '>':
      tk = token.NewToken(token.RANGLE, ">")
    case '\'':
      if l.nextIsChar() {
        l.read()
        tk = token.NewToken(token.CHAR, string(l.c))
      } else {
        tk = token.NewToken(token.UNKNOWN, string(l.c))
        l.msg = "found quote(') but follows unallowed character"
      }
    case '"':
      if ok, s := l.isStr() ; ok {
        tk = token.NewToken(token.STR, s)
      } else {
        tk = token.NewToken(token.UNKNOWN, "")
        l.msg = "found start str (\") but not found end of str (\")"
      }
    case '=':
      tk = token.NewToken(token.EQUAL, "=")
    case 0:
      tk = token.NewToken(token.EOL, "")
    default:
      if l.isSymbol() {
        tk = token.NewToken(token.SYMBOL, string(l.c))
      } else if ok, s := l.isVariable() ; ok {
        tk = token.NewToken(token.VARIABLE, s)
      } else if ok, s := l.isNumber() ; ok {
        tk = token.NewToken(token.NUMBER, s)
      } else {
        tk = token.NewToken(token.UNKNOWN, string(l.c))
        l.msg = "unallowed character -> " + string(l.c)
      }
  }

  l.read()
  return tk
}


func (l *Lexer) skipSpace() {
  for l.c == ' ' || l.c == '\t' {
    l.read()
  }
}

func (l *Lexer) isSymbol() bool {
  return 'a' <= l.c && l.c <= 'z'
}

func (l *Lexer) isVariable() (bool, string) {
  if 'A' <= l.c && l.c <= 'Z' {
    s := []byte{}
    s = append(s, l.c)
    for l.nextIsAlphabet() {
      l.read()
      s = append(s, l.c)
    }
    return true, string(s)
  } else {
    return false, ""
  }
}

func (l *Lexer) nextIsAlphabet() bool {
  if l.nextPosition < len(l.text) {
    c := l.text[l.nextPosition]
    if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
      return true
    }
  }
  return false
}

func (l *Lexer) isNumber() (bool, string) {
  if '0' <= l.c && l.c <= '9' {
    s := []byte{}
    s = append(s, l.c)
    for l.nextIsNumber() {
      l.read()
      s = append(s, l.c)
    }
    return true, string(s)
  } else {
    return false, ""
  }
}

func (l *Lexer) isStr() (bool, string) {
  s := []byte{}
  for l.nextIsStr() {
    l.read()
    if l.c == '"' {
      return true, string(s)
    } else if l.c == '\\' && l.nextIsStr() {
      l.read()
      s = append(s, l.c)
    } else {
      s = append(s, l.c)
    }
  }
  return false, string(s)
}

func (l *Lexer) nextIsNumber() bool {
  if l.nextPosition < len(l.text) {
    c := l.text[l.nextPosition]
    if '0' <= c && c <= '9' {
      return true
    }
  }
  return false
}

func (l *Lexer) nextIsChar() bool {
  if l.nextPosition < len(l.text) {
    c := l.text[l.nextPosition]
    if '!' <= c && c <= '~' {
      return true
    }
  }
  return false
}

func (l *Lexer) nextIsStr() bool {
  if l.nextPosition < len(l.text) {
    c := l.text[l.nextPosition]
    if ' ' <= c && c <= '~' {
      return true
    }
  }
  return false
}

