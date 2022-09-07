package logger

import (
	"log"
	"os"
)

func CreateLogFile(fileName string) (createErr error) {
	kek := "CreateLogFile"
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, createErr := os.OpenFile(fileName, flags, 0666)
	if createErr != nil {
		return createErr
	}
	log.SetOutput(file)
	log.Printf("%s: logging started\n", kek)
	return nil

}
