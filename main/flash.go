package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"os/exec"
	"strings"
)

type FlashMount struct {
	Mounted       string
	MountPoint    string
	DeviceName    string
	MountStatus   int
	UnmountStatus int
}

type FlashUse struct {
	ServiceWork bool
	ProcessWork bool
}

// check potentional disaster procces using the flash
func (f *FlashUse) CheckPid(processName string) bool {
	who := "pidof"
	with := "-s"

	out, _ := exec.Command(who, with, processName).Output()

	fmt.Println(len(out))
	log.Print(string(out))
	if len(out) != 0 {
		f.ProcessWork = true
	} else {
		f.ProcessWork = false
	}

	return f.ProcessWork
}

// checkin service that's processes ffmpeg+gpio things
func (f *FlashUse) CheckService(serviceName string) bool {

	// sudo systemctl status docker.service | grep Active
	grep := exec.Command("grep", "Active")
	command := exec.Command("systemctl", "status", serviceName)
	pipe, _ := command.StdoutPipe()
	defer pipe.Close()
	grep.Stdin = pipe
	command.Start()
	res, _ := grep.Output()
	if strings.Contains(string(res), "active (running)") {
		f.ServiceWork = true
	} else {
		f.ServiceWork = false
	}

	return f.ServiceWork
}

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

var St *unix.Stat_t

func MountPoints() (pointPath string) {

	MountPoint := unix.Lstat("/media/passed3", St)

	return
}

// unmount all flash from mediamountdir
func (f *FlashMount) UmountPoint(mountPoint string) int {
	if unmountErr := unix.Unmount(mountPoint, 0); unmountErr != nil {
		log.Printf("an error \"%s\" occured, while unmounting", unmountErr)
	} else {
		log.Print("unmounted")
	}

	f.UnmountStatus = 0 // good
	return f.UnmountStatus

}
