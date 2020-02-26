package errors

import (
	"fmt"

	"github.com/jinmukeji/plat-pkg/v2/rpc/errors/codes"
)

type RpcError struct {
	Code    codes.Code `json:"code"`    // 错误码
	Message string     `json:"message"` // 额外的错误提示消息
	Cause   error      `json:"cause"`   // 导致本 error 的内部 error
}

// Code returns the Code of the error if it is a RpcError error, codes.OK if err
// is nil, or codes.Unknown otherwise.
func Code(err error) codes.Code {
	// Don't use FromError to avoid allocation of OK status.
	if err == nil {
		return codes.OK
	}
	if re, ok := err.(*RpcError); ok {
		return re.Code
	}
	return codes.Unknown
}

// New returns a Status representing c and msg.
func New(c codes.Code, msg string) *RpcError {
	return &RpcError{
		Code:    c,
		Message: msg,
	}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c codes.Code, format string, a ...interface{}) *RpcError {
	return New(c, fmt.Sprintf(format, a...))
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c codes.Code, msg string) error {
	return New(c, msg)
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface{}) error {
	return Newf(c, format, a...)
}

// Error returns an error with cause and msg.
func ErrorWithCause(c codes.Code, cause error, msg string) error {
	re := New(c, msg)
	return re.WithCause(cause)
}

// Error returns an error with cause and formatted msg.
func ErrorfWithCause(c codes.Code, cause error, format string, a ...interface{}) error {
	re := Newf(c, format, a...)
	return re.WithCause(cause)
}

// WithCause encupsale an err as cause
func (e *RpcError) WithCause(err error) *RpcError {
	if e != nil {
		e.Cause = err
	}

	return e
}

func (e *RpcError) leading() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("[errcode:%d] %s", e.Code, e.Code.Message())
}

func (e *RpcError) Error() string {
	if e == nil {
		return ""
	}

	if len(e.Message) > 0 {
		return fmt.Sprintf("%s: %s", e.leading(), e.Message)
	}

	return e.leading()
}

func (e *RpcError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

func (e *RpcError) DetailedError() string {
	if e == nil {
		return ""
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s ╭∩╮ %v", e.Error(), e.Cause)
	}

	return e.Error()
}
