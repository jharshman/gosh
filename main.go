package main

import (
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

const rcFile string = ".goshrc"
const histFile string = ".gosh_history"

func main() {

	// read .goshrc file
	var conf config.Rc
	conf = config.Init(rcFile)

	// debug print
	fmt.Println(conf.History)
	fmt.Println(conf.HistoryFileSize)

	// initialize list to store command history
	// NOTE: history builtin functionality -
	// history builtin has access to slice length while
	// the total capacity is larger.  When the capacity is reached,
	// chunks of history slices are purged into the .gosh_history file
	histSlice := history.Init(/*conf.History, conf.HistoryFileSize, */histFile)

	// debug print history
	for i := range(histSlice) {
		anEntry := histSlice[i]
		fmt.Println(anEntry)
	}

}
