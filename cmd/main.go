package main

import (
	"github.com/fsnotify/fsnotify"
	"kek/logger"
	pc "kek/parseConf"
	"log"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига

	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	mainconfig := new(pc.Config)
	mainconfig.ParseFromTwoDirs("regConfig.yaml", "")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
				mainconfig.ParseFromTwoDirs("regConfig.yaml", "")
				log.Println(mainconfig.GlobalNetSettings.GlobalNetwork)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Add("./regConfig.yaml")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
