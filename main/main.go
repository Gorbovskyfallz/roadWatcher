package main

import "log"

func main() {

	res := MountI("/dev/sdb1", "/media/passd3")
	log.Println(res)

}
