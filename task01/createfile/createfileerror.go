package createfile

import (
	"fmt"
	"time"
)

// CreateFileError includes error and time, when error raised.
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

// WrapCreateFileError saves error as a field of CreateFileError and
// registers time, when error raised
func WrapCreateFileError(err error) *CreateFileError {
	return &CreateFileError{
		created_at: time.Now(),
		err:        err,
	}
}
