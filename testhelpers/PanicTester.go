package testhelpers

import (
	"fmt"
	"testing"
)

//AssertPanic - prevents panic from exiting and compares the messages
func AssertPanic(t *testing.T, expected string) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
			t.Errorf("Expected: %v but got: %v", expected, err)
		}
	}
}
