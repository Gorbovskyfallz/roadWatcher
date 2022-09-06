package flash

import (
	flashInfo "github.com/moby/sys/mountinfo"
	"golang.org/x/sys/unix"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Flash struct {
	Mounted    bool   // from checker
	MountPoint string // from config
	DeviceName string // from config
	FlashUse   FlashUse
}

type FlashUse struct {
	ServiceWork bool // check service selected from config
	ProcessWork bool // check process from config
}

// Checking - mounted or not to mountpath
//configMountPoint (/media/passed3/flash for e.x.)
func (f *Flash) MountInfo(mountPoint string) (status bool, mountErr error) {
	funcName := "MountInfo"
	if status, mountErr = flashInfo.Mounted(mountPoint); mountErr != nil {
		log.Printf("%s: error \"%v\" occured\n", mountErr, funcName)
		return status, mountErr
	}
	f.Mounted = status
	return f.Mounted, mountErr
}

// not consider some errors of mount e.x. accsess perms (ronly...)
func MountFlash(dev, path string) (exitStatus int) {
	nameOfFunc := "MountFlash"
	mountErr := unix.Mount(dev, path, "exfat", unix.MS_MGC_VAL, "")
	if mountErr != nil {
		switch {
		case mountErr.Error() == "no such device":
			log.Printf("%s: device on path %s\n", nameOfFunc, mountErr.Error())
			exitStatus = 1
		case mountErr.Error() == "no such file or directory":
			log.Printf("%s: mountpath: %s\n", nameOfFunc, mountErr.Error())
			exitStatus = 2
		case mountErr.Error() == "device or resource busy":
			log.Printf("%s: work with device: %s\n", nameOfFunc, mountErr.Error())
			exitStatus = 3
		case mountErr.Error() == "invalid argument":
			log.Printf("%s: arguments: %s\n", nameOfFunc, mountErr.Error())
			exitStatus = 4
		default:
			log.Printf("%s: device %s mounted in path %s succsessfuly\n", nameOfFunc, dev, path)
			exitStatus = 0

		}

	}
	return
}

// unmount all flash from mediamountdir
func (f *Flash) UmountPoint(mountPoint string) int {
	nameOfFunc := "UmountPoint"
	if unmountErr := unix.Unmount(mountPoint, 0); unmountErr != nil {
		log.Printf("%s:error \"%s\" occured\n", nameOfFunc, unmountErr)
	} else {
		log.Printf("%s: %s unmounted\n", nameOfFunc, mountPoint)
	}

	return 0 /////

}

// check potentional disaster procces using the flash
// переделать на сисколы
func (f *FlashUse) CheckPid(processName string) bool {
	util := "pidof"
	withArgs := "-s"
	out, _ := exec.Command(util, withArgs, processName).Output()
	if len(out) != 0 {
		f.ProcessWork = true
	} else {
		f.ProcessWork = false
	}
	return f.ProcessWork
}

// checkin service that's processes ffmpeg+gpio things
func (f *FlashUse) CheckService(serviceName string) bool {
	// возможно, это стоит переписать на системных вызовах без использованися exec
	// sudo systemctl status docker.service | grep Active
	nameOfFunc := "CheckSrvice"
	grep := exec.Command("grep", "Active")
	command := exec.Command("systemctl", "status", serviceName)
	pipe, _ := command.StdoutPipe()
	defer func(pipe io.ReadCloser) {
		closePipeErr := pipe.Close()
		if closePipeErr != nil {
			log.Printf("%s: %w", nameOfFunc, closePipeErr)
		}
	}(pipe)
	grep.Stdin = pipe
	startErr := command.Start()
	if startErr != nil {
		log.Printf("%s: starting %s command: %w", nameOfFunc, grep.String(), startErr)
	}
	res, _ := grep.Output()
	if strings.Contains(string(res), "active (running)") {
		f.ServiceWork = true
	} else {
		f.ServiceWork = false
	}

	return f.ServiceWork
}
