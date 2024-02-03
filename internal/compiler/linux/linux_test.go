package linux

import (
	"contest/pkg/byteconv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
	"unsafe"
)

func runFile(fileName string) (string, error) {
	cmd := exec.Command("./" + fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, *(*string)(unsafe.Pointer(&output)))
	}
	return *(*string)(unsafe.Pointer(&output)), nil
}

func TestCPP(t *testing.T) {
	compiler := NewLinuxCompiler()

	code, err := os.ReadFile("testFiles/cpp.cpp")
	assert.Nil(t, err)

	fileName, err := compiler.CompileCPP(byteconv.String(code))
	assert.Nil(t, err)
	defer os.Remove(fileName)

	actual, err := runFile(fileName)
	assert.Nil(t, err)

	expected := "Hello, World!\n"
	assert.Equal(t, expected, actual)
}

func TestPython(t *testing.T) {
	compiler := NewLinuxCompiler()

	code, err := os.ReadFile("testFiles/python.py")
	assert.Nil(t, err)

	fileName, err := compiler.CompilePython(byteconv.String(code))
	assert.Nil(t, err)
	defer os.Remove(fileName)

	actual, err := runFile(fileName)
	assert.Nil(t, err)

	expected := "Hello, World!\n"
	assert.Equal(t, expected, actual)
}
