package main

import (
	"kek/Config"
	"kek/logger"
	"sync"
)

func main() {
	// сделать флаг для пути логов
	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	configPath := Config.ParsePathfromFlag()
	mainConf := new(Config.Config)
	mainConf.ParseFromYaml(configPath)
	wg := sync.WaitGroup{}
	wg.Add(2)

	watcher := mainConf.AddNotifyWatcher(configPath)
	defer watcher.Close()

	go func() {
		defer wg.Done()
		mainConf.CheckUpdate(watcher, configPath)
	}()

	wg.Wait()

}
