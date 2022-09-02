package logger

import (
	"log"
	"os"
)

func CreateLogFile(logFileName string) (logCreateErr error) {

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return logCreateErr
	}

	log.SetOutput(file)
	log.Println("logging started")

	return nil

}
