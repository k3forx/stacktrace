package stacktrace

import (
	"fmt"
	"io"
	"runtime"
)

func WithStack(err error) error {
	return WithStackDepth(err, 1)
}

func WithStackDepth(err error, depth int) error {
	if err == nil {
		return nil
	}
	return &errWitStack{cause: err, stack: callers(depth + 1)}
}

var _ error = (*errWitStack)(nil)
var _ fmt.Formatter = (*errWitStack)(nil)

type errWitStack struct {
	cause error
	*stack
}

func (e *errWitStack) Error() string {
	return e.cause.Error()
}

func (e *errWitStack) Unwrap() error {
	return e.cause
}

func (e *errWitStack) Format(s fmt.State, verb rune) {
	io.WriteString(s, "aaa")
}

// stack represents a stack of program counters.
type stack []uintptr

func callers(depth int) *stack {
	const numFrames = 32
	var pcs [numFrames]uintptr
	n := runtime.Callers(2+depth, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
