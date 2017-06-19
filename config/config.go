package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jharshman/gosh/xerrors"
	"github.com/sirupsen/logrus"
	"os"
)

// Rc defines the settable variables in .goshrc
type Rc struct {
	HistorySize     int32
	HistoryFileSize int32
}

const rcFile string = ".goshrc"

// Init reads configuration and sets variables
func Init() Rc {
	_, err := os.Stat(rcFile)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		logrus.Fatalf("Could not stat file : %v", err)
	}

	var config Rc
	_, err = toml.DecodeFile(rcFile, &config)
	if err != nil {
		fmt.Println(xerrors.ErrInternal)
		logrus.Fatalf("Failed toml decode: %v", err)
	}

	return config
}
