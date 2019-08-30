package errors

import (
	"regexp"
	"strconv"
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
