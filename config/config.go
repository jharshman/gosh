package config

import (
  "fmt"
  "os"
  "github.com/jharshman/gosh/xerrors"
  "github.com/BurntSushi/toml"
)

// Rc defines the settable variables in .goshrc
type Rc struct {
  History           int
  HistoryFileSize   int
}

// Init reads configuration and sets variables
func Init(cfgFile string) Rc {
  _, err := os.Stat(cfgFile)
  if err != nil {
    fmt.Println(xerrors.ErrInternal)
    // TODO: log error
  }

  var config Rc
  _, err = toml.DecodeFile(cfgFile, &config)
  if err != nil {
    fmt.Println(xerrors.ErrInternal)
    // TODO: log error
  }
  return config
}
