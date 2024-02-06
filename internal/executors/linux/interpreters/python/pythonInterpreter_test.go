package python

import (
	"contest/internal/executors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test(t *testing.T) {
	file, err := os.Open("testFiles/test.py")
	assert.Nil(t, err)
	interpreter := NewPythonInterpreter(file)
	actual, err := interpreter.Execute("2", 1024, 1000)
	assert.Nil(t, err)

	expected := "4\n"
	assert.Equal(t, expected, actual)
}

func TestRuntimeError(t *testing.T) {
	file, err := os.Open("testFiles/runtimeErrorTest.py")
	assert.Nil(t, err)
	interpreter := NewPythonInterpreter(file)
	_, err = interpreter.Execute("2", 1024, 1000)
	assert.ErrorIs(t, err, executors.RuntimeError)
}
