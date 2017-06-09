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
func Init(hSize int, hFileSize int, hFileName string) []Hist {

	var entry string
	var split []string

	histSlice := make([]Hist, hSize, hFileSize)

	hFile, err := os.Open(hFileName)

	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		// log error
	}
	defer hFile.Close()

	scanner := bufio.NewScanner(hFile)
	for scanner.Scan() {
		entry = scanner.Text()
		split = strings.Fields(entry)

		// debug output
		fmt.Println(split[0])
		fmt.Println(split[1:])

		// construct struct
		temp := new(Hist)
		temp.LineNumber = split[0]
		temp.Data = split[1:]

		// place into slace
		histSlice = append(histSlice, Hist{split[0],split[1:]})

	}

	return histSlice

}
