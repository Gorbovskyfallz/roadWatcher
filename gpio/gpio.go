// Package gpio global settings of gpio on raspberry pi 4 for using it as roadawatcher
package gpio

import (
	gpio "github.com/stianeikeland/go-rpio/v4"
	"log"
	"time"
)

// 3 - пин для выключения малины
// 17 - пин кнопки запуска и отключения сервисов монтирования и пот
// 4 - пин индикатора потока (светодиод)
// 27 - пин для кнопки монтирования флешки
// 22 - пин для индикации монтирования флешки

type IoPins struct {
	ShutDownButton gpio.Pin
	MountAndStart  gpio.Pin
	NetworkLed     gpio.Pin
	SystemLed      gpio.Pin
}

// EnableGpio opens memory for using pins of rpi 4
func EnableGpio() (enableGpioErr error) {
	enableGpioErr = gpio.Open()
	if enableGpioErr != nil {
		log.Fatalf("EnableGpio: %v", enableGpioErr)
	}
	return nil
}

// DisableGpio closes memory, uses with defer func
func DisableGpio() (disableGpioErr error) {
	disableGpioErr = gpio.Close()
	if disableGpioErr != nil {
		log.Printf("DisableGpio memory warning: %v", disableGpioErr)
	}
	return nil
}

// SetShutDownButton sets mode for shutdown pin
func (i *IoPins) SetShutdownButton() (setPinErr error) {
	// требуется проверка на срабатываение, возможно, нужно указывать притягивание к
	// земле или высокому сигналу для нормально работы
	i.ShutDownButton = gpio.Pin(3)
	i.ShutDownButton.Input()
	i.ShutDownButton.PullDown() // уточнить состояние кнопки, к чему подключена распайка
	return nil

}

// SetMountAndStartButton sets mode for starting roadWatcher
func (i *IoPins) SetMountAndStartButton() (setPinErr error) {
	// требует то же самое, что и для предидущей функции
	i.MountAndStart = gpio.Pin(17)
	i.MountAndStart.Input()
	i.MountAndStart.PullDown() //
	return nil
}

//SetNetworkLed set mode for network status led
func (i *IoPins) SetNetworkLed() (setPinErr error) {
	i.NetworkLed = gpio.Pin(4)
	i.NetworkLed.Output()
	i.NetworkLed.PullDown() // check!!
	return nil
}

/*tet*/
func Testgpio() {

	_ = gpio.Open()

	defer gpio.Close()

	pin := gpio.Pin(4)

	pin.Output() // Output mode
	for {
		gpio.TogglePin(pin)
		time.Sleep(1 * time.Second)

	}

	//pin.Toggle() // Toggle pin (Low -> High -> Low)

	//pin.Input()    // Input mode
	//_ = pin.Read() // Read state from pin (High / Low)
	//
	//pin.Mode(gpio.Output) // Alternative syntax
	//pin.Write(gpio.High)  // Alternative syntax
}
