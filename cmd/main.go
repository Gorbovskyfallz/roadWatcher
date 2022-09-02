package main

import (
	"kek/logger"
	"kek/parseConf"
	"sync"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига
	err := logger.CreateLogFile("log.txt")
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			conf := new(parseConf.Config)
			conf.ParseFromTwoDirs("regConfig.yaml", "/etc/regConfig.yaml")
		}
		wg.Done()

	}()
	wg.Wait()

}
