package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

// func checking general (not vpn) network with net.DIalTimeout
func generalNetCheck() error {

	host := "google.com"
	port := "80"
	timeout := time.Duration(5 * time.Second)
	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)

	if generalNeterr != nil {
		log.Print(host+":"+port, " not responding ", generalNeterr.Error())
		return generalNeterr
	}

	return nil

}

// func with sudo for stopping rtsp-server
func stopServerService(netErr error) (exitStatus *os.ProcessState) {
	if netErr != nil {
		admin := "sudo"
		who := "systemctl"
		what := "stop"
		whom := "docker.service"
		cmd := exec.Command(admin, who, what, whom)
		log.Print("systemctl: ", cmd.Run())
		//дописать проверку остановки потока во избежание косяка флешки сделать булевую перменную
		// на ретурн
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
