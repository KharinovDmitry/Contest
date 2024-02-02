package linux

import (
	"contest/internal/compiler"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unsafe"
)

type LinuxCompiler struct {
}

func NewLinuxCompiler() LinuxCompiler {
	return LinuxCompiler{}
}

func createFileWithCode(code string, extension string) (*os.File, error) {
	fileName := time.Now().String() + extension
	file, err := os.Create(fileName)

	if err != nil {
		return nil, fmt.Errorf("")
	}
	if _, err = file.WriteString(code); err != nil {
		return nil, fmt.Errorf("In LinuxCompiler(createFileWithCode): %w", err)
	}
	return file, nil
}

func setFilePermission(fileName string) error {
	cmd := exec.Command("chmod", "+x", fileName)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("In LinuxCompiler(setFilePermission): %s", *(*string)(unsafe.Pointer(&output)))
	}
	return nil
}

func (c LinuxCompiler) CompileCPP(code string) (string, error) {
	file, err := createFileWithCode(code, ".cpp")
	defer os.Remove(file.Name())
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("In LinuxCompiler(CompileCPP): %w", err)
	}

	cmd := exec.Command("g++", "-o", file.Name()+".exe", "-x", "c++", file.Name())
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("%w: %s", compiler.ErrCompile, *(*string)(unsafe.Pointer(&output)))
	}

	if err := setFilePermission(file.Name() + ".exe"); err != nil {
		return "", fmt.Errorf("In LinuxCompiler(CompileCPP): %w", err)
	}
	return file.Name() + ".exe", nil
}
