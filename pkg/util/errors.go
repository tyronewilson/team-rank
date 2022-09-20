package util

import (
	"strings"
	"sync"
)

type ErrorHandler interface {
	HandleError(err error)
}

type ErrorStreamer interface {
	ErrorHandler
	ReadErrs() error
}

type errorCollector struct {
	errors ErrorList
	mutex  sync.Mutex
}

func (e *errorCollector) HandleError(err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.errors = append(e.errors, err)
}

type ErrorList []error

func (e ErrorList) Error() string {
	msgs := make([]string, len(e))
	for i, err := range e {
		msgs[i] = err.Error()
	}
	return strings.Join(msgs, "\n")
}

func (e *errorCollector) ReadErrs() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.errors
}

func NewErrorCollector() ErrorStreamer {
	collector := &errorCollector{errors: make(ErrorList, 0)}
	return collector
}
