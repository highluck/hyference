package errors

import (
	"errors"
	"runtime"
	"strings"
)

type Error struct {
	StackTrace []string
	Err        error
	status     ResponseStatus
}

func New(log string) *Error {
	var stack []string
	stack = append(stack, log)
	return &Error{
		StackTrace: stack,
	}
}

func (e *Error) WithStatus(status ResponseStatus) *Error {
	e.status = status
	return e
}

func NewError(err error) *Error {
	return new(err)
}

func new(err error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return &Error{Err: err}
}

func (e *Error) Status(status ResponseStatus) *Error {
	e.status = status
	return e
}

func (e *Error) Error() string {
	if e.Err == nil {
		e.Err = errors.New("")
	}
	ops := []string{e.Err.Error()}
	ops = append(ops, e.StackTrace...)
	return strings.Join(ops, "/")
}

func Wrapper(err error, log string) *Error {
	e := new(err)
	e.StackTrace = append(e.StackTrace, log)
	return e
}

func WrapperWithStatus(err error, log string, status ResponseStatus) error {
	e := new(err)
	e.StackTrace = append(e.StackTrace, log)
	e.status = status
	return e
}

func ParseHttpStatus(err error) ResponseStatus {
	if e, ok := err.(*Error); ok && e.status == Success || e.status == 0 {
		return e.status
	}
	return InternalServerError
}

func GetMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc).Name()
	splits := strings.Split(caller, "/")
	return strings.Join(splits[:], ".")
}
