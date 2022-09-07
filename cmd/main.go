package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"kek/logger"
	"kek/parseConf"
	"log"
	"sync"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига

	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	mainConf := new(parseConf.Config)
	homePath := "regConfig.yaml"
	etcPath := "/etc/regConfig.yaml"
	wg := sync.WaitGroup{}
	wg.Add(2)

	_, parseErr := mainConf.ParseTwoDirs(homePath, etcPath)
	if parseErr != nil {
		log.Fatalf("%s: %v\n", parseErr)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	// Start listening for events.
	err = watcher.Add(homePath)
	if err != nil {
		log.Fatal(err)
	}

	//go func() {
	//	defer wg.Done()
	//	fmt.Println("2 loop go ")
	//	for {
	//
	//	}
	//}()

	go func() {
		defer wg.Done()
		fmt.Println("kekekek")
		mainConf.Parse(*watcher, homePath, etcPath)
	}()

	wg.Wait()

}
