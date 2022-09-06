package network

import (
	"log"
	"net"
	"os/exec"
	"time"
)

type Network struct {
	ModemStatus bool
	VpnStatus   bool
}

func (n *Network) ModemNetCheck(host string, port string) bool {
	funcName := "ModemNetCheck"
	timeout := 5 * time.Second
	_, globalNetErr := net.DialTimeout("tcp", host+":"+port, timeout)

	if globalNetErr != nil {
		log.Printf("%s: %s:%s: %v\n", funcName, host, port, globalNetErr)
		n.ModemStatus = false
	} else {
		n.ModemStatus = true
	}

	return n.ModemStatus

}

func (n *Network) VpnNetCheck(host string) bool {
	funcName := "VPNNetCheck"
	util := "ping"
	args := "-c 3"
	connCheck := exec.Command(util, args, host).Run()
	log.Printf("%s: %v\n", funcName, connCheck)

	if connCheck != nil {
		n.VpnStatus = false
	} else {
		n.VpnStatus = true
	}

	return n.VpnStatus

}
