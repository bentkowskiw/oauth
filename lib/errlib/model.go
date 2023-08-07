package errlib

const (
	ErrCodeNotUnique = 1
)

type XError struct {
	error
	status int
}

type Errors []error
