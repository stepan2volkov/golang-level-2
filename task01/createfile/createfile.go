package createfile

import (
	"fmt"
	"os"
)

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
