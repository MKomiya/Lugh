package handler

type Handler interface {
	On() bool
	Call() error
}
