package gpio

import gpio "github.com/stianeikeland/go-rpio/v4"

func testgpio() {
	err := gpio.Open()
	pin := gpio.Pin(10)

	pin.Output() // Output mode
	pin.High()   // Set pin High
	pin.Low()    // Set pin Low
	pin.Toggle() // Toggle pin (Low -> High -> Low)

	pin.Input()       // Input mode
	res := pin.Read() // Read state from pin (High / Low)

	pin.Mode(gpio.Output) // Alternative syntax
	pin.Write(gpio.High)  // Alternative syntax
}
