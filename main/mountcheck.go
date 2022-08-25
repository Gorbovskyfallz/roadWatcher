package main

import (
	"golang.org/x/sys/unix"
	"log"
)

func MountI(mountpath string) {
	//infoMount, mountErr := minfo.Mounted(mountpath) // а есть ли маунт в эту директорию?
	//if mountErr != nil {
	//	log.Println("dir ", mountpath, "have no mountpoints: ", mountErr)
	//}
	//fmt.Println(infoMount)
	//
	//kek := unix.Unmount("/run/media/jupyter/E264-9720", 0)
	//if kek != nil {
	//	log.Panic(kek)
	//}

	if mountErr := unix.Mount("/dev/sdb1", "/media/passed3", "exfat", 0, ""); mountErr != nil {
		log.Println("error:", mountErr)
		if mountErr.Error() == "no such file or diectory" {
			log.Println("passed assert err")
		}

	} else {
		log.Println("mounted")
	}

}
