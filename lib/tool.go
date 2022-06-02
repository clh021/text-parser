package lib

import (
	"os"
	"path/filepath"
)

func GetProgramPath() (string, error) {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex), err
	}

	exReal, err := filepath.EvalSymlinks(ex)
	return filepath.Dir(exReal), err
}
