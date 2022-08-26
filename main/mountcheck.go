package main

import (
	"golang.org/x/sys/unix"
	"log"
)

func MountI(devPath, mountPath string) (exitStatus int) {
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

	if mountErr := unix.Mount(devPath, mountPath, "exfat", unix.MS_MGC_VAL, ""); mountErr != nil {
		switch {
		case mountErr.Error() == "no such device":
			log.Println(mountErr)
			exitStatus = 1
			return
		case mountErr.Error() == "no such file or directory":
			log.Println(mountErr)
			exitStatus = 2
		//проверить, чего нет, устройства или точки монтирования
		case mountErr.Error() == "device or resource busy":
			log.Println(mountErr)
			exitStatus = 3
		case mountErr.Error() == "invalid argument":
			log.Println(mountErr)
			exitStatus = 4
		default:
			log.Println(mountErr)
			exitStatus = 0

		}

	} else {
		log.Println("mounted device:", devPath, "on mountpoint: ", mountPath)
	}

	return

}
