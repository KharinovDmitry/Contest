package mocks

type MockCompiler struct{}

func (с MockCompiler) CompileCPP(code string) (fileName string, err error) {
	panic("not implemented")
}
