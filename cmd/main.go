package main

import (
	"fmt"
	"kek/logger"
	"kek/parseConf"
	"sync"
)

func main() {

	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	mainConf := new(parseConf.Config)
	homePath := "regConfig.yaml"
	etcPath := "/etc/regConfig.yaml"
	wg := sync.WaitGroup{}
	wg.Add(2)

	watcher := mainConf.ConfWatcher(homePath, etcPath)
	defer watcher.Close()

	go func() {
		defer wg.Done()
		fmt.Println("kekekek")
		mainConf.Parse(watcher, homePath, etcPath)
	}()

	wg.Wait()

}
