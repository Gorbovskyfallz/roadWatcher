package main

import (
	"testing"
)

func TestCheckPID(t *testing.T) {

	processName := "htop"

	res := CheckPID(processName)

	if res == true {
		t.Log("passed", res)
	} else {
		t.Error("not passed", res)
	}

}

func TestGeneralNetCheck(t *testing.T) {
	address := "google.com"
	port := "80"
	boolRes := GeneralNetCheck(address, port)
	if boolRes == true {
		t.Log("passed, value is:", boolRes)
	} else {
		t.Error("not passed, value is:", boolRes)
	}

	failAddress := "com"
	failPort := "81"
	boolRes = GeneralNetCheck(failAddress, failPort)
	if boolRes == false {
		t.Log("passed, value is:", boolRes)
	} else {
		t.Error("not passed, value is:", boolRes)
	}

}

func TestStopServerService(t *testing.T) {

}
