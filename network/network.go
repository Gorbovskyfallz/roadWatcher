package network

import (
	"log"
	"net"
	"os/exec"
	"time"
)

type Network struct {
	ModemNetWorkStatus bool
	VpnNetworkStatus   bool
}

func (n *Network) ModemNetCheck(host string, port string) bool {

	//host := "google.com"
	//port := "80"
	timeout := 5 * time.Second
	_, generalNeterr := net.DialTimeout("tcp", host+":"+port, timeout)

	if generalNeterr != nil {
		log.Println("ModemNetCheck (network package):"+host+":"+port, " not responding ", generalNeterr.Error())
		n.ModemNetWorkStatus = false
	} else {
		n.ModemNetWorkStatus = true
	}

	return n.ModemNetWorkStatus

}

func (n *Network) VpnNetCheck(host string) bool {

	who := "ping"
	with := "-c 3" // дописать в сигнутуру функции количество пинга, чтобы можно было гибко настраивать
	connCheck := exec.Command(who, with, host).Run()
	log.Println("VpnNetCheck (network package): ", connCheck)

	if connCheck != nil {
		n.VpnNetworkStatus = false
	} else {
		n.VpnNetworkStatus = true
	}

	return n.VpnNetworkStatus

}
