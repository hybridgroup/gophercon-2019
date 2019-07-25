package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !button.Get() {
			led.Low()
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
