package repl

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "github.com/inazak/computation/lambda/v01/lambda/eval"
)

var header_message = `
-- this is lambda repl -- 
in REPL:
    :q, :quit          ... quit REPL.
    :l, :lambda        ... switch to lambda mode.
    :e, :expand        ... switch to expand mode.
    :i, :index         ... switch to index mode.
    :r, :reduce        ... switch to reduce mode.
    :a, :ascii         ... switch to ascii mode.
    :f, :file FILEPATH ... read file and parse, eval.
`

func Start(ev *eval.Evaluator) {

  fmt.Printf(header_message)

  sc := bufio.NewScanner(os.Stdin)

  for {
  REPLLOOP:

    fmt.Printf("> ")
    sc.Scan()
    text := sc.Text()

    switch text {
    case "":
      goto REPLLOOP

    case ":q", ":quit", ":exit":
      goto EXITLOOP

    case ":l", ":lambda":
      ev.SetOutputMode(eval.OUTPUT_LAMBDA)
      goto REPLLOOP

    case ":i", ":index":
      ev.SetOutputMode(eval.OUTPUT_INDEX)
      goto REPLLOOP

    case ":e", ":expand":
      ev.SetOutputMode(eval.OUTPUT_EXPAND)
      goto REPLLOOP

    case ":r", ":reduce":
      ev.SetOutputMode(eval.OUTPUT_REDUCE)
      goto REPLLOOP

    case ":a", ":ascii":
      ev.SetOutputMode(eval.OUTPUT_ASCII)
      goto REPLLOOP
    }

    if strings.HasPrefix(text, ":f") || strings.HasPrefix(text, ":file") {
      s := strings.Split(text, " ")
      if len(s) != 2 || strings.TrimSpace(s[1]) == "" {
        fmt.Printf("Error: command invalid - usage :file FILENAME\n\n")
        goto REPLLOOP
      } else {
        file := strings.TrimSpace(s[1])
        fmt.Printf("loading file = %v\n", file)
        LoadAndEvalAndPrint(ev, file)
        goto REPLLOOP
      }
    }

    perr, o := ev.ParseAndEval(text)

    if perr != nil {
      fmt.Printf("ParserError:\n")
      for _, s := range perr {
        fmt.Printf("%v\n", s)
      }
      fmt.Printf("\n")
      continue
    }

    if o.IsError {
      fmt.Printf("EvalError: %v\n", o.Text)
      fmt.Printf("\n")
      continue
    }

    if ok, info := ev.GetDebugInfo() ; ok {
      for _, s := range info {
        fmt.Printf("[debug] %v\n", s)
      }
      ev.ClearDebugInfo()
    }

    fmt.Printf("%v\n\n", o.Text)
  }

  EXITLOOP:
}


func LoadAndEvalAndPrint(ev *eval.Evaluator, filename string) {

  f, err := os.Open(filename)
  if err != nil {
    fmt.Printf("FileOpenError: %v", err)
  }
  defer f.Close()

  s := bufio.NewScanner(f)
  line := 0

  for s.Scan() {
    line += 1
    text := strings.TrimSpace(s.Text())
    if text == "" {
      continue
    }

    perr, o := ev.ParseAndEval(text)

    if perr != nil {
      fmt.Printf("ParserError Line %d:\n", line)
      for _, s := range perr {
        fmt.Printf("%v\n", s)
      }
    }

    if o.IsError {
      fmt.Printf("EvalError Line %d:\n", line)
      fmt.Printf("%v\n", o.Text)
    }

    if o.IsNothing {
      continue //print nothing
    }

    if ok, info := ev.GetDebugInfo() ; ok {
      for _, s := range info {
        fmt.Printf("[debug] %v\n", s)
      }
      ev.ClearDebugInfo()
    }

    fmt.Printf("%v\n", o.Text)
  }
}

