package errlib

import (
	"log"
	"strings"
)

func WithCode(e error, status int) error {

	if x, ok := e.(XError); ok {
		x.status = status
		return x
	}
	x := XError{
		error:  e,
		status: status,
	}
	return x
}

func Code(err error) (code int, ok bool) {
	var x XError
	if x, ok = err.(XError); ok {
		code = x.status
	}
	return
}

func (errors Errors) Error() string {
	s := &strings.Builder{}
	for _, err := range errors {
		s.WriteString(err.Error())
		s.WriteRune('\n')
		s.WriteRune('\r')
	}
	return s.String()
}

func PanicOnErr(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
	panic(err)
}
