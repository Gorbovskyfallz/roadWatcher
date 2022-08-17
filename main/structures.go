package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"
)

type RegHand struct {
	NetworkStatus Network
	FlashUse      FlashUse
	FlashMount    FlashMount
}

type FlashUse struct {
	ServiceWork bool
	ProcessWork bool
}

type Network struct {
	ModemNetWorkStatus bool
	VpnNetworkStatus   bool
}

type FlashMount struct {
	Mounted    string
	MountPoint string
	DeviceName string
}

func (n *Network) ModemNetCheck(host string, port string) bool {

	//host := "google.com"
	//port := "80"
	timeout := 5 * time.Second
	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)

	if generalNeterr != nil {
		log.Print(host+":"+port, " not responding ", generalNeterr.Error())
		n.ModemNetWorkStatus = false
	} else {
		n.ModemNetWorkStatus = true
	}

	return n.ModemNetWorkStatus

}

func (n *Network) VpnNetCheck(host string) bool {

	who := "ping"
	with := "-c 3"
	connCheck := exec.Command(who, with, host).Run()
	fmt.Println("this err: ", connCheck)

	if connCheck != nil {
		n.VpnNetworkStatus = false
	} else {
		n.VpnNetworkStatus = true
	}

	return n.VpnNetworkStatus

}

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

func NewRegHand() (completeRegHand RegHand) {
	completeRegHand.NetworkStatus.ModemNetCheck("google.com", "80")
	completeRegHand.NetworkStatus.VpnNetCheck("10.0.0.1")
	completeRegHand.FlashUse.CheckPid("ffmpeg")
	completeRegHand.FlashUse.CheckService("rtsp-simple-service.service")

	return completeRegHand
}
