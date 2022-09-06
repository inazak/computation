package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/inazak/computation/lambda/v01/lambda/eval"
  "github.com/inazak/computation/lambda/v01/lambda/repl"
  //"github.com/pkg/profile"
)

var usage=`
lambda is simple lambda interpreter.

Usage: 
    lambda [OPTIONS]

Options:
    -d, -debug          ... run with debug print.
    -m, -mode MODENAME  ... select eval & output mode. default is 'reduce'.

      MODENAME: lambda  ... print lambda expression.
                expand  ... expand variable and print lambda expression.
                reduce  ... reduce expression and print.
                 index  ... reduce expression and print with symbol index.
                 ascii  ... reduce expression and print ASCII charactor.

    -f, -file FILEPATH  ... read file from FILEPATH, parse, eval, and print.

in REPL:
    :q, :quit          ... quit REPL.
    :l, :lambda        ... switch to lambda mode.
    :e, :expand        ... switch to expand mode.
    :i, :index         ... switch to index mode.
    :r, :reduce        ... switch to reduce mode.
    :a, :ascii         ... switch to ascii mode.
    :f, :file FILEPATH ... read file and parse, eval.
`

var optionsDebug bool
var optionsMode  string
var optionsFile  string
var optionsHelp  bool

func main() {
  //defer profile.Start(profile.ProfilePath(".")).Stop()

  // options parse
  flag.BoolVar(&optionsDebug,  "debug", false, "run with debug print.")
  flag.BoolVar(&optionsDebug,  "d",     false, "run with debug print.")
  flag.StringVar(&optionsMode, "mode",  "reduce", "change output mode.")
  flag.StringVar(&optionsMode, "m",     "reduce", "change output mode.")
  flag.StringVar(&optionsFile, "file",  "", "input source file.")
  flag.StringVar(&optionsFile, "f",     "", "input source file.")
  flag.BoolVar(&optionsHelp,   "help",  false, "print help message.")
  flag.BoolVar(&optionsHelp,   "h",     false, "print help message.")
  flag.Parse()

  if optionsHelp {
    fmt.Printf("%s", usage)
    os.Exit(0)
  }

  if optionsDebug {
    eval.SetDebug(true)
  }

  var ev *eval.Evaluator
  switch optionsMode {
  case "":
  case "lambda", "LAMBDA":
    ev = eval.NewEvaluator(eval.OUTPUT_LAMBDA)
  case "index", "INDEX":
    ev = eval.NewEvaluator(eval.OUTPUT_INDEX)
  case "expand", "EXPAND":
    ev = eval.NewEvaluator(eval.OUTPUT_EXPAND)
  case "reduce", "REDUCE":
    ev = eval.NewEvaluator(eval.OUTPUT_REDUCE)
  case "ascii", "ASCII":
    ev = eval.NewEvaluator(eval.OUTPUT_ASCII)
  default:
    fmt.Printf("%s", usage)
    os.Exit(1)
  }

  // parse file and eval, exit
  if optionsFile != "" {
    repl.LoadAndEvalAndPrint(ev, optionsFile)
    return
  }

  if len(flag.Args()) != 0 {
    fmt.Printf("%s", usage)
    os.Exit(1)
  }

  repl.Start(ev)
  return
}

