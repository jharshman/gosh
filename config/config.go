package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jharshman/gosh/xerrors"
	"os"
)

// Rc defines the settable variables in .goshrc
type Rc struct {
	History         int
	HistoryFileSize int
}

const rcFile string = ".goshrc"

// Init reads configuration and sets variables
func Init() Rc {
	_, err := os.Stat(rcFile)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		// TODO: log error
	}

	var config Rc
	_, err = toml.DecodeFile(rcFile, &config)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		// TODO: log error
	}
	return config
}
