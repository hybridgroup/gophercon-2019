package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

var (
	pwm      = machine.TCC0
	green    = machine.D3
	channelA uint8
)

func main() {
	machine.InitADC()
	initPWM()

	blue := machine.D12
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D11
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	touch := machine.D9
	touch.Configure(machine.PinConfig{Mode: machine.PinInput})

	bzrPin := machine.D8
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	bzr := buzzer.New(bzrPin)

	dial := machine.ADC{machine.A0}
	dial.Configure(machine.ADCConfig{})

	for {
		dialValue := dial.Get()
		pwm.Set(channelA, pwm.Top()*uint32(dialValue)/0xffff)

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

func initPWM() {
	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println("failed to configure PWM")
		return
	}

	// Configure the channel we'll use as output.
	channelA, err = pwm.Channel(green)
	if err != nil {
		println("failed to configure green channel")
		return
	}
}
