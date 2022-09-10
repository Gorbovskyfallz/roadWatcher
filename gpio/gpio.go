package gpio

import (
	gpio "github.com/stianeikeland/go-rpio/v4"
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
