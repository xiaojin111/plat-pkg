package errors

//go:generate ./gen.sh

import "fmt"

// WrapError 包装一个内部 error
func WrapError(err error, ierr error) error {
	if ierr == nil {
		return err
	}
	return fmt.Errorf("%v: %w", err, ierr)
}
