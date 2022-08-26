package main

import (
	"golang.org/x/sys/unix"
)

// not consider some errors of mount e.x. accsess perms (ronly...)
func MountFlash(devPath, mountPath string) (exitStatus int) {

	if mountErr := unix.Mount(devPath, mountPath, "exfat", unix.MS_MGC_VAL, ""); mountErr != nil {
		switch {
		case mountErr.Error() == "no such device":
			//log.Println(mountErr)
			exitStatus = 1
		case mountErr.Error() == "no such file or directory":
			//log.Println(mountErr)
			exitStatus = 2
		case mountErr.Error() == "device or resource busy":
			//log.Println(mountErr)
			exitStatus = 3
		case mountErr.Error() == "invalid argument":
			//log.Println(mountErr)
			exitStatus = 4
		default:
			//log.Println(mountErr)
			exitStatus = 0

		}

	}

	return

}
