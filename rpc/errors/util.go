package errors

import (
	"fmt"
	"regexp"
	"strconv"

	goerr "errors"
)

var (
	regErrCode = regexp.MustCompile(`\[errcode:(\d+)\]`)
)

func ExtractErrCode(err error) (int, bool) {
	matches := regErrCode.FindStringSubmatch(err.Error())
	if len(matches) <= 1 {
		// no matches
		return 0, false
	}

	code := matches[1]
	if c, err := strconv.Atoi(code); err == nil {
		return c, true
	}
	return 0, false
}

// WrapError 包装一个内部 error
func WrapError(err error, ierr error) error {
	if ierr == nil {
		return err
	}
	return fmt.Errorf("%v: %w", err, ierr)
}

// WrapErrorMsg 包装一个 msg 作为内部 error
func WrapErrorMsg(err error, msg string) error {
	if msg == "" {
		return WrapError(err, nil)
	}
	return WrapError(err, goerr.New(msg))
}
