package main

import (
	"kek/logger"
	"kek/parseConf"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига

	logger.CreateLogFile("TESTFSNOTIFY.TXT")
	mainConf := new(parseConf.Config)
	mainConf.ConfigNotifier("regConfig.yaml", "/etc/regConfig.yaml")
}
