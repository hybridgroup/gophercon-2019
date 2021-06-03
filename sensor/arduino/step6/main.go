package main

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"

	// comes from "tinygo.org/x/tinyfont/freemono"
	"github.com/hybridgroup/gophercon2019/freemono"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	dialValue  uint16
	buttonPush bool
	touchPush  bool
	pwm        = machine.TCC0
	green      = machine.D3
	channelA   uint8
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	go handleDisplay()

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
		dialValue = dial.Get()
		pwm.Set(channelA, pwm.Top()*uint32(dialValue)/0xffff)

		buttonPush = button.Get()
		if !buttonPush {
			blue.Low()
		} else {
			blue.High()
		}

		touchPush = touch.Get()
		if touchPush {
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

func handleDisplay() {
	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  32,
	})

	display.ClearDisplay()

	black := color.RGBA{1, 1, 1, 255}

	for {
		display.ClearBuffer()

		val := strconv.Itoa(int(dialValue))
		msg := []byte("dial: " + val) // + x)
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 10, 20, string(msg), black)

		var radius int16 = 4
		if buttonPush {
			tinydraw.FilledCircle(&display, 16+32*0, 32-radius-1, radius, black)
		} else {
			tinydraw.Circle(&display, 16+32*0, 32-radius-1, radius, black)
		}
		if touchPush {
			tinydraw.FilledCircle(&display, 16+32*1, 32-radius-1, radius, black)
		} else {
			tinydraw.Circle(&display, 16+32*1, 32-radius-1, radius, black)
		}

		display.Display()

		time.Sleep(100 * time.Millisecond)
	}
}
