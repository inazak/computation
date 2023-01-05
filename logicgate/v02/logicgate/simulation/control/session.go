package control

import (
  "fmt"
  "github.com/inazak/computation/logicgate/v02/logicgate/simulation/model"
)

var LO = model.LO
var HI = model.HI

const CLOCK_INTERVAL = 1000

const (
  TYPE_HALT        = 1
  TYPE_CLOCK       = 2
  TYPE_NAND        = 3
  TYPE_LO          = 4
  TYPE_HI          = 5
  TYPE_LED         = 6
  TYPE_CONNECTED   = 7
  TYPE_UNCONNECTED = 8
)

type Session struct {
  circuit     *model.Circuit
  interval    int //millisecond
  port        map[string]*stat //name:stat
  halt        *model.Port
}

type stat struct {
  Port            *model.Port
  Type            int
  Used            bool
  AvailableAsPrev bool
  AvailableAsNext bool
}

type SessionError struct {
  Message string
}

func MakeSession() *Session {
  se := &Session{
    circuit:     model.MakeCircuit(),
    interval:    CLOCK_INTERVAL,
    port:        map[string]*stat{},
  }
  se.halt = se.circuit.MakePort()
  return se
}

func (se *Session) PreUpdate() {
  se.circuit.PreUpdate()
}

func (se *Session) ClockUpAndUpdate() {
  if se.halt.State == HI {
    return
  }
  se.circuit.ClockUp()
  se.circuit.Update()
}

func (se *Session) ClockDownAndUpdate() {
  if se.halt.State == HI {
    return
  }
  se.circuit.ClockDown()
  se.circuit.Update()
}

func (se *Session) Cycle() {
  if se.halt.State == HI {
    return
  }
  se.circuit.Cycle()
}

func (se *Session) GetInterval() int {
  return se.interval
}

func (se *Session) SetInterval(i int) {
  se.interval = i
}

func (se *Session) GetState(name string) int {
  if stat, ok := se.port[name] ; ok {
    return stat.Port.State
  } else {
    return -1 //TODO
  }
}

func (se *Session) IsHalt() bool {
  return se.halt.State == HI
}

func (se *Session) GetLEDPortName() []string {
  t := []string{}
  for n, v := range se.port {
    if v.Type == TYPE_LED {
      t = append(t, n)
    }
  }
  return t
}

func (se *Session) GetUnconnectedPortName() []string {
  t := []string{}
  for n, v := range se.port {
    if v.Type == TYPE_UNCONNECTED {
      t = append(t, n)
    }
  }
  return t
}

func (se *Session) AddNand(inputA, inputB, output string) *SessionError {

  if inputA == inputB || inputA == output || inputB == output {
    return &SessionError {
      Message: fmt.Sprintf("duplicated port name"),
    }
  }

  if ! se.isAvailablePortName(inputA) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available port name", inputA),
    }
  }

  if ! se.isAvailablePortName(inputB) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available port name", inputB),
    }
  }

  if ! se.isAvailablePortName(output) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available port name", output),
    }
  }

  a := se.circuit.MakePort()
  b := se.circuit.MakePort()
  o := se.circuit.Nand(a, b)

  se.port[inputA] = &stat {
    Port: a,
    Type: TYPE_NAND,
    Used: true,
    AvailableAsPrev: false,
    AvailableAsNext: true,
  }
  se.port[inputB] = &stat {
    Port: b,
    Type: TYPE_NAND,
    Used: true,
    AvailableAsPrev: false,
    AvailableAsNext: true,
  }
  se.port[output] = &stat {
    Port: o,
    Type: TYPE_NAND,
    Used: true,
    AvailableAsPrev: true,
    AvailableAsNext: false,
  }

  return nil
}


func (se *Session) Connect(input, output string) *SessionError {

  if input == output {
    return &SessionError {
      Message: fmt.Sprintf("duplicated port name"),
    }
  }

  if se.isConnectableNextPortName(output) {

    // case1. defined-port   CONNECT   defined-port
    if se.isConnectablePrevPortName(input) {

      a := se.port[input].Port
      b := se.port[output].Port
      se.circuit.Connect(a, b)

      se.port[input].Used  = true
      se.port[output].Used = true
      se.port[output].AvailableAsNext = false

      if se.port[output].Type == TYPE_UNCONNECTED {
        se.port[output].Type = TYPE_CONNECTED
      }

    // case2. undefined-port CONNECT   defined-port
    } else if se.isAvailablePortName(input) {

      a := se.circuit.MakePort()
      b := se.port[output].Port
      se.circuit.Connect(a, b)

      se.port[input] = &stat {
        Port: a,
        Type: TYPE_UNCONNECTED,
        Used: true,
        AvailableAsPrev: true,
        AvailableAsNext: true,
      }

      se.port[output].Used = true
      se.port[output].AvailableAsNext = false

      if se.port[output].Type == TYPE_UNCONNECTED {
        se.port[output].Type = TYPE_CONNECTED
      }

    } else {
      return &SessionError {
        Message: fmt.Sprintf("`%s` is not available prev port name", input),
      }
    }

  } else { //output is NOT ConnectableNextPortName

    // case3. defined-port   CONNECT undefined-port
    if se.isConnectablePrevPortName(input) && se.isAvailablePortName(output) {

      a := se.port[input].Port
      b := se.circuit.MakePort()
      se.circuit.Connect(a, b)

      se.port[input].Used = true
      se.port[output] = &stat {
        Port: b,
        Type: TYPE_CONNECTED,
        Used: true,
        AvailableAsPrev: true,
        AvailableAsNext: false,
      }

    } else {
      return &SessionError {
        Message: fmt.Sprintf("`%s` is not available next port name", output),
      }
    }
  }

  return nil
}

func (se *Session) AddLO(name string) *SessionError {

  if ! se.isAvailablePortName(name) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available next port name", name),
    }
  }
  se.port[ name ] = &stat {
    Port: se.circuit.MakePort(),
    Type: TYPE_LO,
    Used: false,
    AvailableAsPrev: true,
    AvailableAsNext: false,
  }
  se.port[ name ].Port.State = LO
  return nil
}

func (se *Session) AddHI(name string) *SessionError {

  if ! se.isAvailablePortName(name) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available next port name", name),
    }
  }
  se.port[ name ] = &stat {
    Port: se.circuit.MakePort(),
    Type: TYPE_HI,
    Used: false,
    AvailableAsPrev: true,
    AvailableAsNext: false,
  }
  se.port[ name ].Port.State = HI
  return nil
}

func (se *Session) AddCLOCK(name string) *SessionError {

  if ! se.isAvailablePortName(name) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available next port name", name),
    }
  }
  se.port[ name ] = &stat {
    Port: se.circuit.GetClock(),
    Type: TYPE_CLOCK,
    Used: false,
    AvailableAsPrev: true,
    AvailableAsNext: false,
  }
  return nil
}

func (se *Session) AddLED(name string) *SessionError {

  if ! se.isAvailablePortName(name) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available next port name", name),
    }
  }
  se.port[ name ] = &stat {
    Port: se.circuit.MakePort(),
    Type: TYPE_LED,
    Used: false,
    AvailableAsPrev: false,
    AvailableAsNext: true,
  }
  return nil
}

func (se *Session) AddHALT(name string) *SessionError {

  if ! se.isAvailablePortName(name) {
    return &SessionError {
      Message: fmt.Sprintf("`%s` is not available next port name", name),
    }
  }
  se.port[ name ] = &stat {
    Port: se.halt,
    Type: TYPE_HALT,
    Used: false,
    AvailableAsPrev: false,
    AvailableAsNext: true,
  }
  return nil
}


func (se *Session) isAvailablePortName(name string) bool {
  if _, ok := se.port[name] ; ok {
    return false
  }
  return true
}

func (se *Session) isConnectablePrevPortName(name string) bool {
  if stat, ok := se.port[name] ; ok && stat.AvailableAsPrev {
    return true
  }
  return false
}

func (se *Session) isConnectableNextPortName(name string) bool {
  if stat, ok := se.port[name] ; ok && stat.AvailableAsNext {
    return true
  }
  return false
}

