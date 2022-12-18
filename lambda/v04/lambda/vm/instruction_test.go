package vm

import (
  "testing"
  "github.com/inazak/computation/lambda/v04/lambda/reader"
)


func ReadCompile(t *testing.T, text string) []Instruction {
  l := reader.NewLexer(text)
  p := reader.NewParser(l)
  expr := p.Parse()
  code := Compile(expr)

  return code
}

func TestCompile1(t *testing.T) {

  text   := "x"
  expect := []Instruction{ Fetch { Name: "x" }, }

  code := ReadCompile(t, text)

  for i, inst := range code {
    if inst.String() != expect[i].String() {
      t.Errorf("expected=%s, but got=%s", expect[i], inst)
    }
  }
}

func TestCompile2(t *testing.T) {

  text   := "(^x.x y)"
  expect := []Instruction{
    Call{
      Code: []Instruction{
        Close { Arg: "x", Code: []Instruction{ Fetch{ Name: "x" }, Return{}, } },
        Fetch { Name: "y" },
        Apply {},
        Return {},
      },
    },
  }

  code := ReadCompile(t, text)

  for i, inst := range code {
    if inst.String() != expect[i].String() {
      t.Errorf("expected=%s, but got=%s", expect[i], inst)
    }
  }
}

