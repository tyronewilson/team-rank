package util

import (
	"github.com/pkg/errors"
	"os"
)

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

func MaybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

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
