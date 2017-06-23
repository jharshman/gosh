package history

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jharshman/gosh/xerrors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Hist struct similar to GNU History Library
type Hist struct {
	lineNumber int32
	timeStamp  string
	data       string
	context    string
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
			log.Fatalf("Failed reading file: %v", err)
		}
		defer hFile.Close()

		scanner := bufio.NewScanner(hFile)
		for scanner.Scan() {
			entry = scanner.Text()
			splitEntry = strings.Fields(entry)

			// construct struct
			num, err := strconv.ParseUint(splitEntry[0], 10, 32)
			if err != nil {
				fmt.Println(xerrors.ErrInternal)
				log.Fatalf("Failed parsing int: %v", err)
			}

			num32 := int32(num)
			data := strings.Join(splitEntry[3:], " ")

			hEntry := &Hist{}
			hEntry.setLineNumber(num32)
			hEntry.setTimeStamp(splitEntry[1])
			hEntry.setContext(splitEntry[2])
			hEntry.setData(data)

			_ = (*hList).PushBack(hEntry)

		}
	}
}

// WriteHistory ...
func WriteHistory(hList **list.List) {

	// write history file
	var entryList []*HistoryEntry
	entryList = make([]*HistoryEntry, (*hList).Len()-1)
	index := 0
	for e := (*hList).Front(); e.Next() != nil; e = e.Next() {

		// create an entry
		temp := &HistoryEntry{
			LineNumber: e.Value.(*Hist).GetLineNumber(),
			Data:       e.Value.(*Hist).GetData(),
			TimeStamp:  e.Value.(*Hist).GetTimeStamp(),
			Context:    e.Value.(*Hist).GetContext(),
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
		log.Fatalf("Failed marshalling data: %v", err)
	}

	if err := ioutil.WriteFile("testfile", out, 0644); err != nil {
		fmt.Println(xerrors.ErrInternal)
		log.Fatalf("Failed writting file: %v", err)
	}
}

/*
*	accessor functions
 */
func (h *Hist) setLineNumber(pLineNumber int32) {
	h.lineNumber = pLineNumber
}

func (h *Hist) setTimeStamp(pTimeStamp string) {
	h.timeStamp = pTimeStamp
}

func (h *Hist) setData(pData string) {
	h.data = pData
}

func (h *Hist) setContext(pContext string) {
	h.context = pContext
}

// GetLineNumber ...
func (h *Hist) GetLineNumber() int32 {
	return h.lineNumber
}

// GetTimeStamp ...
func (h *Hist) GetTimeStamp() string {
	return h.timeStamp
}

// GetData ...
func (h *Hist) GetData() string {
	return h.data
}

// GetContext ...
func (h *Hist) GetContext() string {
	return h.context
}
