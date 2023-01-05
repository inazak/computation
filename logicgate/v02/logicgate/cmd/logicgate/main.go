package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/reader"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/parser"
  "github.com/inazak/computation/logicgate/v02/logicgate/hdl/builder"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/view"
)

var usage=`
xxxxxxxxxxxxxxxxxxxxxxxx

Usage: 
    logicgate [OPTIONS] SOURCEFILE

Options:
    -h, -help         ... print imformation.
`

var optionsHelp   bool

func main() {

  flag.BoolVar(&optionsHelp, "help",   false, "print imformation.")
  flag.BoolVar(&optionsHelp, "h",      false, "print imformation.")

  flag.Parse()

  if optionsHelp {
    fmt.Printf("%s", usage)
    os.Exit(0)
  }

  if len(flag.Args()) != 1 {
    fmt.Printf("%s", usage)
    os.Exit(1)
  }
  filename := flag.Args()[0]

  sc, e1 := reader.ReadFromFile(filename)
  if e1 !=nil {
    fmt.Printf("error in reader: %s\n", e1.Message)
    os.Exit(2)
  }

  stmts, e2 := parser.Parse(sc)
  if e2 !=nil {
    fmt.Printf("error in parser: %s\n", e2.Message)
    os.Exit(2)
  }

  se, e3 := builder.Build(stmts)
  if e3 !=nil {
    fmt.Printf("error in builder: %s\n", e3.Message)
    os.Exit(3)
  }

  var ui view.UI
  ui = view.MakeCommandlineUI(se)

  ui.Start()
  os.Exit(0)
}

