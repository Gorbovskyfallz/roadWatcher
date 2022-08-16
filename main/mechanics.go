package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// func checking general (not vpn) network with net.DIalTimeout
//func GeneralNetCheck(host string, port string) (uptime bool) {
//
//	//host := "google.com"
//	//port := "80"
//	timeout := 5 * time.Second
//	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)
//
//	if generalNeterr != nil {
//		log.Print(host+":"+port, " not responding ", generalNeterr.Error())
//		uptime = false
//	} else {
//		uptime = true
//	}
//
//	return
//
//}

// func with sudo for stopping rtsp-server
func StopServerService(uptime bool, serviceName string) (exitStatus *os.ProcessState) {
	if uptime == false {
		admin := "sudo"
		who := "systemctl"
		what := "stop"
		cmd := exec.Command(admin, who, what, serviceName)
		log.Print("systemctl: ", cmd.Run())
		fmt.Println("stopped server service")
		exitStatus = cmd.ProcessState

	}
	return exitStatus

}

func CheckPID(processName string) (pidResult bool) {
	who := "pidof"
	with := "-s"
	whom := "htop"

	out, _ := exec.Command(who, with, whom).Output()
	if len(out) != 0 {
		pidResult = true
	} else {
		pidResult = false
	}

	return pidResult
}

// unmount flash func
func unmountFlash(netErr error, serviceExitStatus *os.ProcessState) {
	if netErr != nil || serviceExitStatus.ExitCode() != 0 {
		admin := "sudo"
		who := "umount"
		what := "/media/flash"
		_ = exec.Command(admin, who, what)
		//expectedError := errors.New("")
		//umountError := fmt.Errorf( "%w", cmd.Run())
		log.Print("flash unmounted")

	}
	// добавить проверку, извлечена ли флешка

}
