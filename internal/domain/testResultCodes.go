package domain

type TestResultCode string

const (
	TimeLimitCode       TestResultCode = "TL"
	MemoryLimitCode     TestResultCode = "ML"
	CompileErrorCode    TestResultCode = "CE"
	RuntimeErrorCode    TestResultCode = "RE"
	SuccesCode          TestResultCode = "SC"
	IncorrectAnswerCode TestResultCode = "IA"
)
