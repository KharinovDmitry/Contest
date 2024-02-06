package python

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test(t *testing.T) {
	file, err := os.Open("python.py")
	assert.Nil(t, err)
	interpreter := NewPythonInterpreter(file)
	actual, err := interpreter.Execute("2", 1024, 100)
	assert.Nil(t, err)

	expected := "4\n"
	assert.Equal(t, expected, actual)
}
