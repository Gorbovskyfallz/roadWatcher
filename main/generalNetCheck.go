package main

import (
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

// shitty func with sudo for stopping rtsp-simple-service
func stopServerService(netErr error) {
	if netErr != nil {
		admin := "sudo"
		who := "systemctl"
		what := "stop"
		whom := "docker"
		cmd := exec.Command(admin, who, what, whom)
		log.Print(cmd.Output())
	}

}

func unmountFlash(netErr error) {

}
