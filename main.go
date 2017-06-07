package main

import (
  "fmt"
  "github.com/jharshman/gosh/config"
)

/*
  TODO:
  - load config files
  - loop and process input
  - run cleanup if any
*/

func main() {
  var conf config.Rc
  conf = config.Init(".goshrc")
  fmt.Println(conf.History)
  fmt.Println(conf.HistoryFileSize)
}
