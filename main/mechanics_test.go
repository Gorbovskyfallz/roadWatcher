package main

import "testing"

func TestCheckPID(t *testing.T) {

	processName := "htop"

	res := CheckPID(processName)

	if res == true {
		t.Log("passed", res)
	} else {
		t.Error("not passed", res)
	}

}
