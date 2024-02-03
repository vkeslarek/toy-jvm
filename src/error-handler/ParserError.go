package errorhandler

import (
	"fmt"
	"strings"
)

type ParserError struct {
	Field   string
	Message string
	Cause   *ParserError
}

func Wrap(field string, err error) ParserError {
	return ParserError{
		Field:   field,
		Message: err.Error(),
		Cause:   nil,
	}
}

func New(field string, message string, cause ParserError) ParserError {
	return ParserError{
		Field:   field,
		Message: message,
		Cause:   &cause,
	}
}

func (e *ParserError) Panic() {
	panic(e)
}

func (e *ParserError) Error() string {
	return e.errorTabulated(0)
}

func (e *ParserError) errorTabulated(numberOfTabs int) string {
	// Print in the format of Java stack trace, idented by tabs
	return fmt.Sprintf("%sField: %s, Message: %s\nCaused by: %s", strings.Repeat("\t", numberOfTabs), e.Field, e.Message, e.Cause.errorTabulated(numberOfTabs+1))
}
