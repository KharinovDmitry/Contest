package cpp

import (
	"contest/internal/executors"
	"contest/internal/executors/linux/utils"
	"contest/pkg/byteconv"
	"fmt"
	"os"
	"os/exec"
)

type CPPCompiler struct {
}

func NewCPPCompiler() CPPCompiler {
	return CPPCompiler{}
}

func (c CPPCompiler) Compile(code string) (fileName string, err error) {
	file, err := utils.CreateTempFileWithText(code, ".cpp")
	defer os.Remove(file.Name())
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("In CPPCompiler(Compile): %w", err)
	}

	executableName := file.Name() + ".exe"
	cmd := exec.Command("g++", "-o", executableName, "-x", "c++", file.Name())
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("%w: %s", executors.CompileError, byteconv.String(output))
	}

	if err := utils.AddFileExecutablePermission(executableName); err != nil {
		return "", fmt.Errorf("In CPPCompiler(Compile): %w", err)
	}
	return executableName, nil
}
