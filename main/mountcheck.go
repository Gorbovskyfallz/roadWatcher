package main

import (
	"fmt"
	minfo "github.com/moby/sys/mountinfo"
	"golang.org/x/sys/unix"
	"log"
)

func MountI(mountpath string) {
	infoMount, _ := minfo.Mounted(mountpath) // а есть ли маунт в эту директорию?
	fmt.Println(infoMount)

	kek := unix.Unmount("/run/media/jupyter/E264-9720", 0)
	if kek != nil {
		log.Panic(kek)
	}

}
