package handler

import (
	"testing"
)

type testhandler struct {
}

func (*testhandler) On() bool {
	return true
}

func (*testhandler) Call() error {
	return nil
}

func NewTestHandler() Handler {
	return new(testhandler)
}

func TestListenCall(t *testing.T) {
	handlers := []Handler{NewTestHandler()}

	err := ListenCall(handlers)
	if err != nil {
		t.Error("Failed")
	}
}
