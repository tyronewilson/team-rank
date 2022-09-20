package util

import (
	"github.com/pkg/errors"
	"os"
)

// SplitOnLastSpace splits a string on the last space, returning the string before the space and the string after the space
func SplitOnLastSpace(str string) (string, string) {
	if len(str) == 0 {
		return "", ""
	}
	ptr := len(str) - 1
	for ptr >= 0 && str[ptr] != ' ' {
		ptr--
	}
	if ptr == -1 {
		return "", ""
	}
	return str[:ptr], str[ptr+1:]
}

// MaybePanic panics if err is not nil
func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckAllFilesExist checks that all files exist, and returns an error if any do not
func CheckAllFilesExist(filenames []string) error {
	for _, filename := range filenames {
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				return errors.Errorf("file %s does not exist", filename)
			}
			return err
		}
	}
	return nil
}
