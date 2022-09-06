package logger

import (
	"log"
	"os"
)

func CreateLogFile(name string) (createErr error) {
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(name, flags, 0666)
	if err != nil {
		return createErr
	}
	log.SetOutput(file)
	log.Println("logging started")
	return nil

}
