package main

import "log"

func main() {

	res := MountI("/dev/sdb1", "/media/passed3")
	log.Println(res)

}
