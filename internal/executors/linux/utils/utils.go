package utils

import (
	"contest/pkg/byteconv"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func CreateTempFileWithText(text string, extension string) (*os.File, error) {
	fileName := time.Now().String() + extension
	file, err := os.Create(fileName)

	if err != nil {
		return nil, fmt.Errorf("")
	}
	if _, err = file.WriteString(text); err != nil {
		return nil, fmt.Errorf("In utils(CreateFileWithText): %w", err)
	}
	return file, nil
}

func AddFileExecutablePermission(fileName string) error {
	cmd := exec.Command("chmod", "+x", fileName)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("In utils(AddFileExecutablePermission): %s", byteconv.String(output)+err.Error())
	}
	return nil
}
