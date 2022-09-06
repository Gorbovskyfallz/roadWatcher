package tests

import (
	"kek/flash"
	"kek/reghandler"
	"testing"
)

func TestFlashUse_CheckPid(t *testing.T) {
	testStruct := new(flash.FlashUse)

	testStruct.CheckPid("ffplay")
	// start the htop of ffmpeg!! must be true while process in running
	if testStruct.ProcessWork == true {
		t.Log("passed")
	} else {
		t.Errorf("ecpected \"true\", but received %t, or testing process os not runnig at this moment.", testStruct.ProcessWork)
	}

}

func TestFlashUse_CheckService(t *testing.T) {

	testStruct := new(flash.FlashUse)

	testStruct.CheckService("docker.service")

	//while docker service is active!!

	if testStruct.ServiceWork == true {
		t.Log("passed")
	} else {
		t.Errorf("expected value is \"true\", but received value is %t", testStruct.ServiceWork)
	}

}

func TestNewRegHand(t *testing.T) {

	testStruct := reghandler.NewRegHand()

	if testStruct.NetworkStatus.VpnStatus == true {
		t.Log("VPN: must passed if vpn is enabled")
	} else {
		t.Errorf("VPN: expected \"true\", but received %t", testStruct.NetworkStatus.VpnStatus)
	}
	if testStruct.NetworkStatus.ModemStatus == true {
		t.Log("NetAddress: must passed if network is enabled")
	} else {
		t.Errorf("NetAddress: expected \"true\", but received %t", testStruct.NetworkStatus.ModemStatus)
	}
	if testStruct.FlashUse.ProcessWork == true {
		t.Log("ProcessWork: must passed if ffmpeg is executing now")
	} else {
		t.Errorf("ProcessWork: expected \"true\", but received %t", testStruct.FlashUse.ProcessWork)
	}
	if testStruct.FlashUse.ServiceWork == true {
		t.Log("ServiceWork: must passed if rtsp-simple-server.service is executing now")
	} else {
		t.Errorf("ServiceWork: expected \"true\", but received %t", testStruct.FlashUse.ServiceWork)
	}

}
