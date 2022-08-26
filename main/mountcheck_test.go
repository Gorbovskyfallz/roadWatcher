package main

import (
	"testing"
)

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

// неправильный путь к точке монтирования
func TestNoSuchFileDir(t *testing.T) {
	devPath := "/dev/sdb1"
	badMountPath := "/media/passed"
	noFileExit := 2

	if noFileDir := MountI(devPath, badMountPath); noFileDir != noFileExit {
		t.Error("expected noFileExit=2, but received:", noFileExit)
	}

}

// invalid argument while trying to mount with bad dev path INVALID ARGUMENT
// flash card must be unmounted to pass this test
func TestBadDevNotMounted(t *testing.T) {
	//unix.Unmount("/media/passed3", 0)
	badDevPath := "/dev/sdb"
	MountPath := "/media/passed3"
	invalidArgExit := 4
	if badDevPath := MountI(badDevPath, MountPath); badDevPath != invalidArgExit {
		t.Errorf("expected %v, but received: %v", invalidArgExit, badDevPath)
	}
}
