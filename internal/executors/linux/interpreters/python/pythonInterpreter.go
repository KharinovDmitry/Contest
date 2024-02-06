package python

import (
	"contest/internal/executors"
	"contest/internal/services/runTestService"
	"contest/pkg/byteconv"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type PythonInterpreter struct {
	codeFile *os.File
}

func NewPythonInterpreter(codeFile *os.File) runTestService.IExecutor {
	return &PythonInterpreter{codeFile: codeFile}
}

func (c *PythonInterpreter) Execute(input string, memoryLimitInKb int, timeLimitInMs int) (output string, err error) {
	timeout := time.Duration(timeLimitInMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "python3", c.codeFile.Name())
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("In codeRunnerPython(Execute): %w", err)
	}
	defer stdin.Close()

	fmt.Fprintln(stdin, input)

	outputBytes, err := cmd.CombinedOutput()
	outputString := byteconv.String(outputBytes)
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok && status.Signaled() && status.Signal() == syscall.SIGKILL {
			return "", executors.TimeLimitError
		}

		return outputString, fmt.Errorf("%w: %s", executors.RuntimeError, err)
	}

	return outputString, nil
}

func (c *PythonInterpreter) Close() error {
	defer c.codeFile.Close()

	err := os.Remove(c.codeFile.Name())
	if err != nil {
		fmt.Errorf("In codeRunnerPython(Close): %w", err)
	}
	return nil
}
