package main

import "testing"

func TestMountNoDevise(t *testing.T) {
	// используй битый путь
	resultNoDevice := 1

	noDeviceTest := MountI("/dev/sdb1", "/media/passed3")
	if noDeviceTest != resultNoDevice {
		t.Error("expected:", resultNoDevice, "but received:", noDeviceTest)
	} else {
		t.Log("noDeviceTest: PASSED")
	}
	// нужна примонтированнная флешка

}

// before testing u must be sure, that flash <devPath> is mounted to <mountPath>
func TestBusyDevice(t *testing.T) {
	devPath := "/dev/sdb1"
	mountPath := "/media/passed3"
	deviceBusy := 3
	if busyDevice := MountI(devPath, mountPath); busyDevice != deviceBusy {
		t.Error("expected 3, but received:", busyDevice)
	}

}
