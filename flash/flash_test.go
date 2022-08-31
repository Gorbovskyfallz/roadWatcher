package flash

import (
	"testing"
)

// before testing u must be sure, that flash <devPath> is mounted to <mountPath>
func TestBusyDevice(t *testing.T) {
	devPath := "/dev/sdb1"
	mountPath := "/media/passed3"
	deviceBusy := 3
	if busyDevice := MountFlash(devPath, mountPath); busyDevice != deviceBusy {
		t.Error("expected 3, but received:", busyDevice)
	}

}

// before testing u nust be sure, that flash is exist, and is not mounted
func TestNoErrMount(t *testing.T) {
	devPath := "/dev/sdb1"
	mountPath := "/media/passed3"
	noErrExit := 0
	if noErr := MountFlash(devPath, mountPath); noErr != noErrExit {
		t.Error("expected exit=0, but received:", noErr)
	}
}

// неправильный путь к точке монтирования
// так же учитывается ошибка неверного пути к блочному устройству
func TestNoSuchFileDir(t *testing.T) {
	devPath := "/dev/sdb1"
	badMountPath := "/media/passed"
	noFileExit := 2
	if noFileDir := MountFlash(devPath, badMountPath); noFileDir != noFileExit {
		t.Error("expected noFileExit=2, but received:", noFileExit)
	}

}

// invalid argument while trying to mount with bad dev path INVALID ARGUMENT
// flash card must be unmounted to pass this test
// тестирование ошибки, связанной с неверным аргументом раздела блочного устройства при отмонтированном
// блочном устройстве
func TestBadDevNotMounted(t *testing.T) {
	//unix.Unmount("/media/passed3", 0)
	badDevPath := "/dev/sdb"
	MountPath := "/media/passed3"
	invalidArgExit := 4
	if badDevPath := MountFlash(badDevPath, MountPath); badDevPath != invalidArgExit {
		t.Errorf("expected %v, but received: %v", invalidArgExit, badDevPath)
	}
}

//func TestFlashMount_UmountPoint(t *testing.T) {
//	testStruct := new(FlashMount)
//	mountPointPaths := "/media/"
//	testStruct.UmountPoint(mountPointPaths)
//	if testStruct.UnmountStatus != 0 {
//		t.Log("error while unmounting:")
//	}
//}

func TestFlashMount_MountInfo(t *testing.T) {
	//for umounted flash
	path := "/media/passed3"
	testStruct := new(FlashMount)
	testStruct.MountInfo(path)
	if testStruct.Mounted != true {
		t.Error("|while flash is mounted| expected true, but received:", testStruct.Mounted)
	}
}
