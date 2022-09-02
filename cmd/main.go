package main

import (
	"github.com/fsnotify/fsnotify"
	"kek/logger"
	"log"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига
	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("cant create watcher")
	}
	defer watch.Close()
	err = watch.Add("./regConfig.yaml")
	if err != nil {
		log.Fatalln("cant add file to watcher")
	} else {
		log.Println("succsess")
	}

}
