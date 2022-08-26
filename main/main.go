package main

import "log"

func main() {

	res := MountI("/dev/sdb1", "/media/pased3")
	log.Println(res)

}
