package domain

type RunTestRequest struct {
	TaskID          int
	Language        Language
	Code            string
	MemoryLimitInKb int
	TimeLimitInMs   int
}
