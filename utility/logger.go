package utility

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(message string) {
	if l.infoLogger != nil {
		l.infoLogger.Println(message)
	}
}

func (l *Logger) Error(message string) {
	if l.errorLogger != nil {
		l.errorLogger.Println(message)
	}
}

func (l *Logger) Debug(message string) {
	if l.debugLogger != nil {
		l.debugLogger.Println(message)
	}
}

func LogAndPrint(logger *Logger, message string) {
	log.Println(message)
	if logger != nil {
		logger.Info(message)
	}
}
