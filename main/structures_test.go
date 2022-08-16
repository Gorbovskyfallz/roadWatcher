package main

import (
	"fmt"
	"testing"
)

func TestNetwork_GeneralNetCheck(t *testing.T) {
	testStruct := new(Network)

	testStruct.ModemNetCheck("google.com", "80")
	fmt.Println(testStruct.ModemNetWorkStatus)
	if testStruct.ModemNetWorkStatus == true {
		t.Log("passed")
	} else {
		t.Error("not passed")
	}

	testStruct.ModemNetCheck("com", "80")
	fmt.Println(testStruct.ModemNetWorkStatus)
	if testStruct.ModemNetWorkStatus == false {
		t.Log("passed")

	} else {
		t.Error("not passed")
	}
}

func TestNetwork_VpnNetCheck(t *testing.T) {

	testStruct := new(Network)

	testStruct.VpnNetCheck("10.0.0.1")

	if testStruct.VpnNetworkStatus == true {
		t.Log("passed")
	} else {
		t.Log("expected true, but received: ", testStruct.VpnNetworkStatus)
	}

}

func TestFlashUse_CheckPid(t *testing.T) {
	testStruct := new(FlashUse)

	testStruct.CheckPid("htop")
	// start the htop of ffmpeg!!
	if testStruct.ProcessWork == true {
		t.Log("passed")
	} else {
		t.Errorf("ecpected \"true\", but received %t, or testing process os not runnig at this moment.", testStruct.ProcessWork)
	}

}
