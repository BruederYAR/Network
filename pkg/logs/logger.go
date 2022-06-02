package logs

import (
	"log"
	"os"
)

type Logger struct {
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
}

func NewLogger(path string, username string) (*Logger, error) {
	file, err := os.OpenFile(path+"log_"+username+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		Info:    log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile),
		Warning: log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile),
		Error:   log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile),
	}

	return logger, nil
}

func (logger *Logger) LogError(err error) {
	logger.Error.Println(err)
}

func (logger *Logger) LogWarning(message string) {
	logger.Warning.Println(message)
}

func (logger *Logger) LogInfo(message string) {
	logger.Info.Println(message)
}
