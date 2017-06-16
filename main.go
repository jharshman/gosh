package main

import (
	"container/list"
	"fmt"
	"github.com/jharshman/gosh/config"
	"github.com/jharshman/gosh/history"
)

/*
  TODO:
  - load config files
  - read history file
  - initialize history
  - loop and process input
  - run cleanup if any
*/

func main() {

	// read .goshrc file
	var conf config.Rc
	conf = config.Init()

	// debug print
	fmt.Println(conf.History)
	fmt.Println(conf.HistoryFileSize)

	// history use container/list instead of array / slice
	hList := list.New()
	history.Init(&hList)

	// debug print history
	for e := hList.Front(); e.Next() != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	history.WriteHistory(&hList)

}
