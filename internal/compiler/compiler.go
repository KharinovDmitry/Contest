package compiler

import (
	"errors"
)

var (
	ErrCompile = errors.New("Compile Error")
)

type Compiler interface {
	CompileCPP(code string) (fileName string, err error)
	CompilePython(code string) (fileName string, err error)
}
