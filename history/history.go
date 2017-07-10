package history

import (
	"bufio"
	"bytes"
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
	lineNumber string
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
			data := strings.Join(splitEntry[3:], " ")

			hEntry := &Hist{}
			hEntry.setLineNumber(splitEntry[0])
			hEntry.setTimeStamp(splitEntry[1])
			hEntry.setContext(splitEntry[2])
			hEntry.setData(data)

			_ = (*hList).PushBack(hEntry)

		}
	}
}

// WriteHistory appends full command history to file
// starting at the point defiend by the passed in parameter
func WriteHistory(start *list.Element, hList **list.List) {

	// format for history entry:
	// <linenuber> <timestamp> <data/command> <context>

	f, err := os.OpenFile(histFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		log.Fatalf("Failed opening file for write: %v", err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	for e := start.Next(); e.Next() != nil; e = e.Next() {

		line := simpleHistoryEntry(e.Value)

		writer.WriteString(line)

	}

	writer.Flush()

}

// AddEntry creates a new Hist structure and pushes it onto the list
func AddEntry(command string, timestamp string, context string, hList **list.List) {

	listLength := strconv.Itoa((*hList).Len() + 1)

	hEntry := &Hist{}
	hEntry.setLineNumber(listLength)
	hEntry.setTimeStamp(timestamp)
	hEntry.setContext(context)
	hEntry.setData(command)
	_ = (*hList).PushBack(hEntry)
}

func simpleHistoryEntry(t interface{}) string {
	var buf bytes.Buffer
	buf.WriteString(t.(*Hist).GetLineNumber())
	buf.WriteString(" ")
	buf.WriteString(t.(*Hist).GetTimeStamp())
	buf.WriteString(" ")
	buf.WriteString(t.(*Hist).GetData())
	buf.WriteString(" ")
	buf.WriteString(t.(*Hist).GetContext())
	buf.WriteString("\n")
	return buf.String()
}

// WriteHistoryProtobuf ...
// Testing practicality of writing history log in protobuffer
func WriteHistoryProtobuf(hList **list.List) {

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
func (h *Hist) setLineNumber(pLineNumber string) {
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
func (h *Hist) GetLineNumber() string {
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
