package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

// func checking general (not vpn) network with net.DIalTimeout
func generalNetCheck() error {

	host := "com"
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
func stopServerService(netErr error) {
	if netErr != nil {
		admin := "sudo"
		who := "systemctl"
		what := "stop"
		whom := "rtsp-simple-server.service" // пробы идут на докере
		cmd := exec.Command(admin, who, what, whom)
		log.Print("systemctl: ", cmd.Run())
		//дописать проверку остановки потока во избежание косяка флешки сделать булевую перменную
		// на ретурн
		fmt.Println("stopped server service")

	}

}

// unmount flash func
func unmountFlash(netErr error) {
	if netErr != nil {
		admin := "sudo"
		who := "umount"
		what := "/media/flash"
		_ = exec.Command(admin, who, what)
		//expectedError := errors.New("")
		//umountError := fmt.Errorf( "%w", cmd.Run())

	}
	// добавить проверку, извлечена ли флешка

}
