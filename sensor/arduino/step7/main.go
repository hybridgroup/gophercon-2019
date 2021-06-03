package main

import (
	"image/color"
	"machine"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"

	// comes from "tinygo.org/x/tinyfont/freemono"
	"github.com/hybridgroup/gophercon2019/freemono"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/espat"
	"tinygo.org/x/drivers/net/mqtt"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	dialValue  uint16
	buttonPush bool
	touchPush  bool
	pwm        = machine.TCC0
	green      = machine.D3
	channelA   uint8

	uart = machine.UART1
	tx   = machine.PA22
	rx   = machine.PA23

	adaptor *espat.Device
	topic   = "tinygo"

	display ssd1306.Device
)

// IP address of the MQTT broker to use. Replace with your own info, if so desired.
const server = "ssl://test.mosquitto.org:8883"

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	rand.Seed(time.Now().UnixNano())

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

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

	// Init esp8266/esp32
	adaptor = espat.New(uart)
	adaptor.Configure()

	// first check if connected
	if connectToESP() {
		blue.High()
		println("Connected to wifi adaptor.")
		adaptor.Echo(false)

		blue.Low()
		connectToAP()
		blue.High()
	} else {
		println("")
		failMessage("Unable to connect to wifi adaptor.")
		return
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID("tinygo-client-" + randomString(10))

	blue.Low()
	println("Connectng to MQTT...")
	cl := mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	initDisplay()

	go handleDisplay()

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

		println("Publishing MQTT message...")
		data := []byte("{\"e\":[{ \"n\":\"hello\", \"sv\":\"world\" }]}")
		token := cl.Publish(topic, 0, false, data)
		token.Wait()
		if token.Error() != nil {
			println(token.Error().Error())
			break
		}

		time.Sleep(time.Millisecond * 100)
	}

	// Right now this code is only reached when there is an error. Need a way to trigger clean exit.
	println("Disconnecting MQTT...")
	cl.Disconnect(100)

	println("Done.")
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

func initDisplay() {
	display = ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  32,
	})

	display.ClearDisplay()
}

func handleDisplay() {
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

// connect to ESP8266/ESP32
func connectToESP() bool {
	for i := 0; i < 5; i++ {
		println("Connecting to wifi adaptor...")
		if adaptor.Connected() {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// connect to access point
func connectToAP() {
	println("Connecting to wifi network...")

	adaptor.SetWifiMode(espat.WifiModeClient)
	adaptor.ConnectToAP(ssid, pass, 10)

	println("Connected.")
	println(adaptor.GetClientIP())
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
