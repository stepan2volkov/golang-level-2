// Package createfile implements function which works with
// package "os". It also demonstrates wrapping error with
// inculding information about error-time
package createfile

import (
	"fmt"
	"os"
)

// CreateFile creates file or returns an error which type
// is CreateFileError
func CreateFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return WrapCreateFileError(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return nil
}
