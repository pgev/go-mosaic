package moos

import (
	"fmt"
	"os"
)

// thanks Tendermint

// onlyUserDirPerm is the permissions used when creating directories.
const onlyUserDirPerm = 0700

// EnsureDir ensures the directory exists, and if not creates it recursively
// with permissions 0700, or errors if it cannot create the directory
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, onlyUserDirPerm)
		if err != nil {
			return fmt.Errorf("could not create directory %v: %w", dir, err)
		}
	}
	return nil
}
