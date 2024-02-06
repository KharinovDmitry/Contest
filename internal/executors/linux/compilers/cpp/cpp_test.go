package cpp

import (
	"contest/pkg/byteconv"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	code, err := os.ReadFile("cpp")
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
