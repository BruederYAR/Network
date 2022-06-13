package logs

type ILogger interface {
	LogInfo(message string)
	LogWarning(message string)
	LogError(err error)
	LogPanic(message string)
}
