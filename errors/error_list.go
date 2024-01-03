package errors

import (
	"fmt"
	"sync"
)

type ErrorList struct {
	m      sync.Mutex
	errors []error
}

func (el *ErrorList) Len() int {
	return len(el.errors)
}

func (el *ErrorList) Add(e error) {
	el.m.Lock()
	el.errors = append(el.errors, e)
	el.m.Unlock()
}

func (el *ErrorList) Error() string {
	var message string
	for _, err := range el.errors {
		message += fmt.Sprintf("%s\n", err)
	}

	return message
}

func New() *ErrorList {
	return &ErrorList{}
}
