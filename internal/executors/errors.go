package executors

import (
	"errors"
)

var (
	ErrUnknownLanguage = errors.New("unknown language")

	CompileError     = errors.New("Compile error")
	MemoryLimitError = errors.New("Memory limit")
	TimeLimitError   = errors.New("Time limit")
	RuntimeError     = errors.New("Runtime error")
)
