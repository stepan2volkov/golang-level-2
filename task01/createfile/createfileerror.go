package createfile

import (
	"fmt"
	"time"
)

type CreateFileError struct {
	created_at time.Time
	err        error
}

func (e CreateFileError) Error() string {
	return fmt.Sprintf("CreateFileError at %v: %v", e.created_at, e.err)
}

func (e CreateFileError) Unwrap() error {
	return e.err
}

func WrapCreateFileError(err error) *CreateFileError {
	return &CreateFileError{
		created_at: time.Now(),
		err:        err,
	}
}
