package gpio

import (
	gpio "github.com/stianeikeland/go-rpio/v4"
	"kek/flash"
	"kek/network"
	"time"
)

//индикация отсутствия интернета
// индикация отсутствия впн
// индикация ошибки потока
// индикация ошибки монтирования
// индикация "все хорошо"

// сначала мы накидаем просто названивая функций, надо попробовать развивать мы
//шление относительно уровней абстракции

type LedIndicator interface {
	NetworkIndicate(Network network.Network)
	SystemIndicate(StatFlash flash.StatFlash)
}

func (i *IoPins) VpnErrorToggle() {
	gpio.TogglePin(i.NetworkLed)
	time.Sleep(1 * time.Second)
}

func (i *IoPins) ModemErrorToggle() {
	gpio.TogglePin(i.NetworkLed)
	time.Sleep(300 * time.Millisecond)
}

func (i *IoPins) NetworkVpnErrorToggle() {
	gpio.TogglePin(i.NetworkLed)
	time.Sleep(5 * time.Second)
}
func (i *IoPins) AllFineConnectToggle() {
	i.NetworkLed.High()

}

func (i *IoPins) NetworkIndicate(Network network.Network) {
	if !Network.ModemStatus {
		i.ModemErrorToggle()
	} else if !Network.VpnStatus {
		i.VpnErrorToggle()
	} else {
		i.AllFineConnectToggle()
	}

}

func (i *IoPins) FlashMountErrorIndicate() {
	gpio.TogglePin(i.SystemLed)
	time.Sleep(200 * time.Millisecond)
}

func (i *IoPins) UnmountErrorIndicate() {
	gpio.TogglePin(i.SystemLed)
	time.Sleep(1 * time.Second)
}

func (i *IoPins) FlashMountedIndicate() {
	i.SystemLed.High()
}

func (i *IoPins) SystemIndicate(flashErrors flash.FlashErrors) {
	if flashErrors.UnmountExitStatus != 0 {
		i.UnmountErrorIndicate()
	} else if flashErrors.MountExitStatus != 0 {
		i.FlashMountErrorIndicate()
	} else {
		i.FlashMountedIndicate()
	}
}
