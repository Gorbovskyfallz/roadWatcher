package main

import "testing"

// before testing u must be sure, that flash <devPath> is mounted to <mountPath>
func TestBusyDevice(t *testing.T) {
	devPath := "/dev/sdb1"
	mountPath := "/media/passed3"
	deviceBusy := 3
	if busyDevice := MountI(devPath, mountPath); busyDevice != deviceBusy {
		t.Error("expected 3, but received:", busyDevice)
	}

}

// before testing u nust be sure, that flash is exist, and is not mounted
func TestNoErrMount(t *testing.T) {
	devPath := "/dev/sdb1"
	mountPath := "/media/passed3"
	noErrExit := 0
	if noErr := MountI(devPath, mountPath); noErr != noErrExit {
		t.Error("expected exit=0, but received:", noErr)
	}
}

func TestNoSuchFileDir(t *testing.T) {
	devPath := "/dev/sdb1"
	badMountPath := "/media/passed"
	noFileExit := 2
	if noFileDir := MountI(devPath, badMountPath); noFileDir != noFileExit {
		t.Error("expected noFileExit=2, but received:", noFileExit)
	}

}
