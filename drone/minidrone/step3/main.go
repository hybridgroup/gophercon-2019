package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := minidrone.NewDriver(bleAdaptor)

	work := func() {
		drone.On(drone.Event("battery"), func(data interface{}) {
			fmt.Printf("battery: %d\n", data)
		})

		drone.On(minidrone.Hovering, func(data interface{}) {
			fmt.Println("hovering!")
		})

		drone.On(minidrone.Landing, func(data interface{}) {
			fmt.Println("landing!")
		})

		drone.On(minidrone.Landed, func(data interface{}) {
			fmt.Println("landed.")
		})

		drone.TakeOff()
		gobot.After(10*time.Second, func() {
			fmt.Println("forwards...")
			drone.Forward(5)
		})
		gobot.After(12*time.Second, func() {
			fmt.Println("backwards...")
			drone.Backward(10)
		})
		gobot.After(14*time.Second, func() {
			fmt.Println("hovering...")
			drone.Stop()
		})

		gobot.After(20*time.Second, func() {
			fmt.Println("right...")
			drone.Right(5)
		})
		gobot.After(22*time.Second, func() {
			fmt.Println("left...")
			drone.Left(10)
		})
		gobot.After(24*time.Second, func() {
			fmt.Println("hovering...")
			drone.Stop()
		})
		gobot.After(30*time.Second, func() {
			fmt.Println("landing...")
			drone.Land()
		})
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
