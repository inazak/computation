package reader

import (
  "fmt"
  "os"
  "io/ioutil"
  "regexp"
)

type ReaderError struct {
  Message string
}

type SourceCode struct {
  Line []string
}

func ReadFromString(s string) (sc *SourceCode, err *ReaderError) {

  sc = &SourceCode {
    Line: []string{},
  }
  sc.Line = regexp.MustCompile(`\r\n|\n`).Split(s, -1)

  return sc, nil
}

func ReadFromFile(filepath string) (sc *SourceCode, err *ReaderError) {

  f, e := os.OpenFile(filepath, os.O_RDONLY, 0666)

  if e != nil {
    err = &ReaderError{ Message: fmt.Sprintf("%s", e.Error()), }
    return nil, err
  }
  defer f.Close()

  b, e := ioutil.ReadAll(f)
  if e != nil {
    err = &ReaderError{ Message: fmt.Sprintf("%s", e.Error()), }
    return nil, err
  }

  return ReadFromString(string(b))
}

