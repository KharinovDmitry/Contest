package linux

import (
	"contest/internal/domain"
	"contest/internal/executors"
	"contest/internal/executors/linux/compilers/cpp"
	"contest/internal/executors/linux/interpreters/python"
	"contest/internal/executors/linux/utils"
	"contest/internal/services/runTestService"
	"errors"
	"fmt"
	"os"
)

type Compiler interface {
	Compile(code string) (fileName string, err error)
}

type ExecutorFactory struct {
	languageCompilerMap    map[domain.Language]Compiler
	languageInterpreterMap map[domain.Language]func(codeFile *os.File) runTestService.IExecutor
}

func NewExecutorFactory() *ExecutorFactory {
	return &ExecutorFactory{
		languageCompilerMap: map[domain.Language]Compiler{
			domain.CPP: cpp.NewCPPCompiler(),
		},
		languageInterpreterMap: map[domain.Language]func(codeFile *os.File) runTestService.IExecutor{
			domain.Python: python.NewPythonInterpreter,
		},
	}
}

func (c *ExecutorFactory) NewExecutor(code string, language domain.Language) (runTestService.IExecutor, error) {
	if compiler, exist := c.languageCompilerMap[language]; exist {
		fileName, err := compiler.Compile(code)
		if err != nil {
			if errors.Is(err, executors.CompileError) {
				return nil, err
			}
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		file, err := os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		return &ExecutorCompiled{executableFile: file}, nil
	}
	if constructor, exist := c.languageInterpreterMap[language]; exist {
		codeFile, err := utils.CreateTempFileWithText(code, ".code")
		if err != nil {
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		return constructor(codeFile), nil
	}
	return nil, executors.ErrUnknownLanguage
}
