package history

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/jharshman/gosh/xerrors"
	"os"
	"strings"
)

// Hist struct similar to GNU History Library
type Hist struct {
	LineNumber string
	Data       []string
	TimeStamp  string
	Context    string
}

const histFile string = ".gosh_history"

// Init initializes history slice and reads .gosh_history file
func Init(hList **list.List) {

	var entry string
	var splitEntry []string

	if _, err := os.Stat(histFile); err == nil {
		hFile, err = os.Open(histFile)
		defer hFile.Close()
		if err != nil {
			fmt.Println(xerrors.ErrInternal)
			// log error
		}
	}

	scanner := bufio.NewScanner(hFile)
	for scanner.Scan() {
		entry = scanner.Text()
		splitEntry = strings.Fields(entry)

		// construct struct
		hEntry := new(Hist)
		hEntry.LineNumber = splitEntry[0]
		hEntry.Data = splitEntry[1:]

		_ = (*hList).PushBack(hEntry)

	}

}
