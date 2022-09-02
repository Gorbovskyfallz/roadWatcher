package main

import (
	"fmt"
	"kek/parseConf"
)

func main() {
	// 1. считываем конфиг, возможно даже мы должны делать это в отдельной горутине для того, чтобы конфиг считывался
	// конкурентно и подтятигивал новые параметры, которые могли бы быть измененены. Для начала просто считаем конфиг
	// и запихнем его в объект конфига

	testConf := new(parseConf.Config)
	// input by cli
	testConf.Security.EnableTokenConfigParse = false
	testConf.Security.TokenBotApi = "kjh"
	testConf.SwitchTokenInput()
	fmt.Println(testConf.Security.TokenBotApi)
}
