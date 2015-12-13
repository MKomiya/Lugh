package main

import (
	"./handler"
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

func NewTestHandler() handler.Handler {
	return new(testhandler)
}

func TestListenCall(t *testing.T) {
	handlers := []handler.Handler{NewTestHandler()}

	err := listenCall(handlers)
	if err != nil {
		t.Error("Failed")
	}
}
