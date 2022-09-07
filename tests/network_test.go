package tests

import (
	"fmt"
	"kek/network"
	"testing"
)

func TestNetwork_GeneralNetCheck(t *testing.T) {
	testStruct := new(network.Network)

	testStruct.NetCheck("google.com", "80")
	fmt.Println(testStruct.ModemStatus)
	if testStruct.ModemStatus == true {
		t.Log("passed")
	} else {
		t.Error("not passed")
	}

	testStruct.NetCheck("com", "80")
	fmt.Println(testStruct.ModemStatus)
	if testStruct.ModemStatus == false {
		t.Log("passed")

	} else {
		t.Error("not passed")
	}
}

func TestNetwork_VpnNetCheck(t *testing.T) {

	testStruct := new(network.Network)

	testStruct.VpnCheck("10.0.0.1")

	if testStruct.VpnStatus == true {
		t.Log("passed")
	} else {
		t.Log("expected true, but received: ", testStruct.VpnStatus)
	}

}
