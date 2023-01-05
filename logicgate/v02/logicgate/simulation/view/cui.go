package view

import (
  "fmt"
  "time"
  "sort"
  "strconv"

  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/control"
)

type CommandlineUI struct {
  session *control.Session
}

func MakeCommandlineUI(se *control.Session) *CommandlineUI {
  cui := &CommandlineUI {
    session: se,
  }
  return cui
}

func (ui *CommandlineUI) Start() {

  ui.session.PreUpdate()

  ports := ui.session.GetLEDPortName()
  sort.Strings(ports)

  count := 0

  ui.printLED(ports, count)

  // sleep, update, halt or print
  for {

    if ui.session.IsHalt() {
      break
    }

    time.Sleep( time.Duration(ui.session.GetInterval()) * time.Millisecond )

    ui.session.Cycle()
    count += 1

    ui.printLED(ports, count)

  } //for
}

func (ui *CommandlineUI) printLED(list []string, count int) {

  if len(list) == 0 {
    return
  }

  first := list[0]
  fmt.Printf("%06d: %s=[%d]  ", count, first, ui.session.GetState(first))
  prev, _ := strconv.Atoi(first)

  for i := 1 ; i < len(list) ; i++ {
    item := list[i]
    curr, _ := strconv.Atoi(item)
    if curr - prev != 1 {
      fmt.Printf("\n%06d: ", count)
    }
    fmt.Printf("%s=[%d]  ", item, ui.session.GetState(item))
    prev = curr
  }

  fmt.Printf("\n----\n")
}

