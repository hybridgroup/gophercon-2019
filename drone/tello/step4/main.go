package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

var drone = tello.NewDriver("8888")

func main() {
	var currentFlightData *tello.FlightData

	work := func() {
		fmt.Println("takeoff...")

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		drone.On(tello.FlipEvent, func(data interface{}) {
			fmt.Println("Flip")
		})

		drone.TakeOff()

		gobot.Every(1*time.Second, func() {
			printFlightData(currentFlightData)
		})

		gobot.After(5*time.Second, func() {
			performFlips()
		})

		gobot.After(20*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func printFlightData(d *tello.FlightData) {
	if d.BatteryLow {
		fmt.Printf(" -- Battery low: %d%% --\n", d.BatteryPercentage)
	}

	displayData := `
Height:         %d
Ground Speed:   %d
Light Strength: %d

`
	fmt.Printf(displayData, d.Height, d.GroundSpeed, d.LightStrength)
}

func performFlips() {
	drone.FrontFlip()
	time.Sleep(time.Second * 3)
	drone.BackFlip()
}
