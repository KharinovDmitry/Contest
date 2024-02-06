package cpp

import (
	"contest/internal/executors"
	"contest/pkg/byteconv"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	code, err := os.ReadFile("testFiles/test.cpp")
	assert.Nil(t, err)
	compiler := NewCPPCompiler()

	file, err := compiler.Compile(byteconv.String(code))
	assert.Nil(t, err)
	defer os.Remove(file)

	cmd := exec.Command("./" + file)
	actual, err := cmd.CombinedOutput()
	assert.Nil(t, err)

	expected := "Hello, World!\n"
	assert.Equal(t, expected, byteconv.String(actual))
}

func TestRuntimeError(t *testing.T) {
	code, err := os.ReadFile("testFiles/compileErrorTest.cpp")
	assert.Nil(t, err)
	compiler := NewCPPCompiler()

	file, err := compiler.Compile(byteconv.String(code))
	assert.ErrorIs(t, err, executors.CompileError)
	defer os.Remove(file)
}
