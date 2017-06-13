package history

import (
	"bufio"
	"fmt"
	"github.com/jharshman/gosh/xerrors"
	"os"
	"strings"
)

// Hist struct similar to GNU History Library
type Hist struct {
	LineNumber string
	Data       []string
	// TimeStamp  string
}

// Init initializes history slice and reads .gosh_history file
func Init(hSize int, hFileSize int, hFileName string) []*Hist {

	var entry string
	var split []string

	histSlice := make([]*Hist, hSize, hFileSize)

	hFile, err := os.Open(hFileName)

	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		// log error
	}
	defer hFile.Close()

	scanner := bufio.NewScanner(hFile)
	index := 0
	for scanner.Scan() {
		entry = scanner.Text()
		split = strings.Fields(entry)

		// construct struct
		hEntry := new(Hist)
		hEntry.LineNumber = split[0]
		hEntry.Data = split[1:]

		// place into slice
		// append is the wrong thing to use here
		// the function append grows the underlying array beyond
		// it's capacity
		histSlice[index] = hEntry
		index++

	}

	return histSlice

}
