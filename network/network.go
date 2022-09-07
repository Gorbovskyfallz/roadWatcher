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

func (n *Network) NetCheck(host string, port string) bool {
	funcName := "NetCheck"
	timeout := 5 * time.Second
	_, netErr := net.DialTimeout("tcp", host+":"+port, timeout)

	if netErr != nil {
		log.Printf("%s: %s:%s: %v\n", funcName, host, port, netErr)
		n.ModemStatus = false
	} else {
		n.ModemStatus = true
	}

	return n.ModemStatus

}

func (n *Network) VpnCheck(host string) bool {
	funcName := "VpnCheck"
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
