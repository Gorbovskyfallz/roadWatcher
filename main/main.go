package main

import "log"

func main() {

	res := MountFlash("/dev/sda", "/media/passed1")
	log.Println(res)

}
