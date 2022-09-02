package flash

import (
	"github.com/moby/sys/mountinfo"
	"golang.org/x/sys/unix"
	"log"
	"os/exec"
	"strings"
)

type FlashMount struct {
	Mounted    bool   // from checker
	MountPoint string // from config
	DeviceName string // from config
	FlashUse   FlashUse
}

type FlashUse struct {
	ServiceWork bool // check service selected from config
	ProcessWork bool // check process from config
}

func (f *FlashMount) MountInfo(configMountPoint string) (mounted bool, mountInfoErr error) {

	if mounted, mountInfoErr = mountinfo.Mounted(configMountPoint); mountInfoErr != nil {
		log.Printf("MountInfo (flash package): an error \"%v\" occured while checking the mountpoints \n", mountInfoErr)
		return mounted, mountInfoErr
	}
	f.Mounted = mounted
	return f.Mounted, mountInfoErr
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
			exitStatus = 0

		}

	}

	return

}

// unmount all flash from mediamountdir
func (f *FlashMount) UmountPoint(mountPoint string) int {
	if unmountErr := unix.Unmount(mountPoint, 0); unmountErr != nil {
		log.Printf("UmountPoint (flash package): an error \"%s\" occured, while unmounting\n", unmountErr)
	} else {
		log.Print("UmountPoint (flash package): unmounted\n")
	}

	return 0 /////

}

// check potentional disaster procces using the flash
func (f *FlashUse) CheckPid(processName string) bool {
	who := "pidof"
	with := "-s"

	out, _ := exec.Command(who, with, processName).Output()
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
