package util

import (
	"strings"
	"sync"
)

// ErrorHandler is an interface which can handle errors.
type ErrorHandler interface {
	HandleError(err error)
}

// ErrorStreamer is an interface which can handle errors and allow you to read them as a single error as well.
type ErrorStreamer interface {
	ErrorHandler
	ReadErrs() error
}

type errorCollector struct {
	errors ErrorList
	mutex  sync.Mutex
}

// HandleError adds the error to the list of errors.
func (e *errorCollector) HandleError(err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.errors = append(e.errors, err)
}

// ErrorList is a list of errors.
type ErrorList []error

// Error returns a string representation of the error list.
func (e ErrorList) Error() string {
	msgs := make([]string, len(e))
	for i, err := range e {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "\n")
}

// ReadErrs returns the collected errors.
func (e *errorCollector) ReadErrs() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.errors
}

// NewErrorCollector returns a new ErrorStreamer which will simply collect errors.
func NewErrorCollector() ErrorStreamer {
	collector := &errorCollector{errors: make(ErrorList, 0)}
	return collector
}
