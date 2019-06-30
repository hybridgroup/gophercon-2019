package main

import (
	"bufio"
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"os"
)

var drone = tello.NewDriver("8888")

func main() {
	var currentFlightData *tello.FlightData

	work := func() {

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		drone.On(tello.FlipEvent, func(data interface{}) {
			fmt.Println("Flip")
		})

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			printFlightData(currentFlightData)
			switch scanner.Text()[0] {
			// [ ] to Take off and land
			case '[':
				fmt.Println("takeoff...")
				drone.TakeOff()
			case ']':
				fmt.Println("land...")
				drone.Land()

			// WSAD to control forward, backward, left, and right
			case 'w':
				fmt.Println("forward...")
				drone.Forward(20)
			case 's':
				fmt.Println("backward...")
				drone.Backward(20)
			case 'a':
				fmt.Println("left...")
				drone.Left(20)
			case 'd':
				fmt.Println("right...")
				drone.Right(20)

			// IKJL to control up, down, spin counter clockwise, spin clockwise
			case 'i':
				fmt.Println("up...")
				drone.Up(20)
			case 'k':
				fmt.Println("down...")
				drone.Down(20)
			case 'j':
				fmt.Println("spin counter clockwise...")
				drone.CounterClockwise(tello.ValidatePitch(20, 10))
			case 'l':
				fmt.Println("spin clockwise...")
				drone.Clockwise(tello.ValidatePitch(20, 10))

			// r to stop motion of the tello
			case 'r':
				fmt.Println("stop movement...")
				drone.Clockwise(0)
				drone.CounterClockwise(0)
        drone.Forward(0)
        drone.Backward(0)
        drone.Left(0)
        drone.Right(0)
        drone.Up(0)
        drone.Down(0)

			// TGFH to flip front, flip back, flip left, flip right
			case 't':
				fmt.Println("front flip...")
				drone.FrontFlip()
			case 'g':
				fmt.Println("back flip...")
				drone.BackFlip()
			case 'f':
				fmt.Println("left flip...")
				drone.LeftFlip()
			case 'h':
				fmt.Println("right flip...")
				drone.RightFlip()
			}
		}
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
