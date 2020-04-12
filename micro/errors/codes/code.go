package codes

import (
	"fmt"
	"strconv"
)

//go:generate ./gen.sh

// A Code is an unsigned 32-bit error code as defined in the spec.
type Code uint32

// Message returns a description message of the code
func (c *Code) Message() string {
	msg := codeToMsg[*c]

	return msg
}

// UnmarshalJSON unmarshals b into the Code.
func (c *Code) UnmarshalJSON(b []byte) error {
	// From json.Unmarshaler: By convention, to approximate the behavior of
	// Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as
	// a no-op.
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 32); err == nil {
		if _, ok := codeToMsg[Code(ci)]; !ok {
			return fmt.Errorf("invalid code: %q", ci)
		}

		*c = Code(ci)
		return nil
	}

	return fmt.Errorf("invalid code: %q", string(b))
}
