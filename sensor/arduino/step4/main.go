package main

import (
	"machine"
	"time"
)

func main() {
	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	red := machine.D10
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !button.Get() {
			blue.Low()
			red.High()
		} else {
			blue.High()
			red.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
