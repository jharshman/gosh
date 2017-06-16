package history

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jharshman/gosh/xerrors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Hist struct similar to GNU History Library
type Hist struct {
	LineNumber int32
	TimeStamp  string
	Data       string
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
			lineNumber, err := strconv.ParseUint(splitEntry[0], 10, 32)
			if err != nil {
				fmt.Println(xerrors.ErrInternal)
				// log error
			}

			lineNumber32 := int32(lineNumber)
			data := strings.Join(splitEntry[3:], " ")

			hEntry := new(Hist)
			hEntry.LineNumber = lineNumber32
			hEntry.TimeStamp = splitEntry[1]
			hEntry.Context = splitEntry[2]
			hEntry.Data = data

			_ = (*hList).PushBack(hEntry)

		}
	}
}

func WriteHistory(hList **list.List) {

	// write history file
	var entryList []*HistoryEntry
	entryList = make([]*HistoryEntry, (*hList).Len()-1)
	index := 0
	for e := (*hList).Front(); e.Next() != nil; e = e.Next() {

		// create an entry
		temp := &HistoryEntry{
			LineNumber: e.Value.(*Hist).LineNumber,
			Data:       e.Value.(*Hist).Data,
			TimeStamp:  e.Value.(*Hist).TimeStamp,
			Context:    e.Value.(*Hist).Context,
		}
		entryList[index] = temp
		index++
	}

	entries := &HistoryLog{
		History: entryList,
	}

	out, err := proto.Marshal(entries)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		fmt.Println(err)
		// log error
	}

	if err := ioutil.WriteFile("testfile", out, 0644); err != nil {
		fmt.Println(xerrors.ErrInternal)
		// log error
	}
}
