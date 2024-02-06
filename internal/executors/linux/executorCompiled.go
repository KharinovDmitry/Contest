package linux

import (
	"contest/internal/executors"
	"contest/pkg/byteconv"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type ExecutorCompiled struct {
	executableFile *os.File
}

func (e *ExecutorCompiled) Execute(input string, memoryLimitInKb int, timeLimitInMs int) (output string, err error) {
	timeout := time.Duration(timeLimitInMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "./"+e.executableFile.Name())
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("In CodeRunnerCompiled(Execute): %w", err)
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

func (e *ExecutorCompiled) Close() error {
	defer e.executableFile.Close()

	err := os.Remove(e.executableFile.Name())
	if err != nil {
		fmt.Errorf("In CodeRunnerCompiled(Close): %w", err)
	}
	return nil
}
