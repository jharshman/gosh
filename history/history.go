package history

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/golang/protobuf/proto"
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

// Init initializes history and reads .gosh_history file
// if .gosh_history doesn't exist, nothing happens and the list
// remains empty.
// .gosh_history will be written / appended to later
func Init(hList **list.List) {

	var entry string
	var splitEntry []string
	var hFile *os.File

	if _, err := os.Stat(histFile); err == nil {
		hFile, err = os.Open(histFile)
		if err != nil {
			fmt.Println(xerrors.ErrInternal)
			// log error
		}
		defer hFile.Close()

		scanner := bufio.NewScanner(hFile)
		for scanner.Scan() {
			entry = scanner.Text()
			splitEntry = strings.Fields(entry)

			// construct struct
			hEntry := new(Hist)
			hEntry.LineNumber = splitEntry[0]
			hEntry.TimeStamp = splitEntry[1]
			hEntry.Context = splitEntry[2]
			hEntry.Data = splitEntry[3:]

			_ = (*hList).PushBack(hEntry)

		}
	}
}

func WriteHistory(hList **list.List) {

	// write history file

	for e := (*hList).Front(); e.Next() != nil; e = e.Next() {

		// create an entry
		entry := &HistoryEntry{
			LineNumber: e.Value.(*Hist).LineNumber,
			Data:       e.Value.(*Hist).Data,
			TimeStamp:  e.Value.(*Hist).TimeStamp,
			Context:    e.Value.(*Hist).Context,
		}

	}
}
