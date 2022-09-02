package tests

import (
	"fmt"
	"kek/network"
	"testing"
)

func TestNetwork_GeneralNetCheck(t *testing.T) {
	testStruct := new(network.Network)

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

	testStruct := new(network.Network)

	testStruct.VpnNetCheck("10.0.0.1")

	if testStruct.VpnNetworkStatus == true {
		t.Log("passed")
	} else {
		t.Log("expected true, but received: ", testStruct.VpnNetworkStatus)
	}

}
