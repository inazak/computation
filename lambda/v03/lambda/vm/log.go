package vm

import (
  "log"
  "os"
)

func (vm *VM) EnableLogging() {
  vm.logger = log.New(os.Stderr, "", log.LstdFlags)
}

func (vm *VM) DisableLogging() {
  vm.logger = nil
}

func (vm *VM) logf(format string, v ...interface{}) {
  if vm.logger == nil {
    return
  }
  vm.logger.Printf(format, v...)
}

func (vm *VM) debugPrint() {
  if vm.logger == nil {
    return
  }

  vm.logger.Printf("---- stack ----", )
  for _, v := range vm.stack {
    vm.logger.Printf("%s", v)
  }

  vm.logger.Printf("---- env ----", )
  for k, v := range vm.env {
    vm.logger.Printf("%s: %s", k, v)
  }

  vm.logger.Printf("---- code ----", )
  for _, v := range vm.code {
    vm.logger.Printf("%s", v)
  }

  vm.logger.Printf("**************")
  vm.logger.Printf("")
}


