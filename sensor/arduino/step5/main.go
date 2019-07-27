package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

func main() {
	machine.InitADC()
	machine.InitPWM()

	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	green := machine.PWM{machine.D10}
	green.Configure()

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	touch := machine.D9
	touch.Configure(machine.PinConfig{Mode: machine.PinInput})

	bzrPin := machine.D8
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	dial := machine.ADC{machine.A0}
	dial.Configure()

	for {
		green.Set(dial.Get())

		if !button.Get() {
			blue.Low()
		} else {
			blue.High()
		}

		if touch.Get() {
			bzr.On()
		} else {
			bzr.Off()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
