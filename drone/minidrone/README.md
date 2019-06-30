# Parrot Minidrone

The various Parrot Mindrones such as the Rolling Spider all use the same API.

## What you need

    - Parrot Minidrone
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux (kernel v4.14+) or macOS. Sorry, no Windows yet.

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your Minidrone as needed.

### macOS

To run any of the Gobot BLE code on OS X, you must use the `GODEBUG=cgocheck=0` flag.

For example:

```
$ GODEBUG=cgocheck=0 go run drone/minidrone/step1/main.go Mars_1234
```

### Linux

On Linux the BLE code will need to run as a root user account. The easiest way to accomplish this is probably to use `go build` to build your program, and then to run the requesting executable using `sudo`.

For example:

```
$ go build -o step01 drone/minidrone/step1/main.go
$ sudo ./step01 Mars_1234
```

## Code

### step01/main.go

Let's start with a simple takeoff, and then land. Make sure the drone is turned on, then run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step1/main.go [dronename]

### Linux

    go build -o step1 drone/minidrone/step1/main.go
    sudo ./step1 [dronename]

### step02/main.go

The drone can return some flight data. Run this code:

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step2/main.go [dronename]

### Linux

    go build -o step2 drone/minidrone/step2/main.go
    sudo ./step2 [dronename]

### step03/main.go

The drone can move forward, backward, to the right, and the left, all while maintaining a steady altitude. Run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step3/main.go [dronename]

### Linux

    go build -o step3 drone/minidrone/step3/main.go
    sudo ./step3 [dronename]

### step04/main.go

The drone can perform flips while flying. Run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step4/main.go [dronename]

### Linux

    go build -o step4 drone/minidrone/step4/main.go
    sudo ./step4 [dronename]

### step05/main.go

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

Run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step5/main.go [dronename]

### Linux

    go build -o step5 drone/minidrone/step5/main.go
    sudo ./step5 [dronename]

### step06/main.go

Now that you have mastered the flight controls, let's report your drone's flight status to the Internet. We will connect to a MQTT machine to machine messaging server that is maintained by the Eclipse Foundation for public testing.

Run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step6/main.go [dronename] ssl://iot.eclipse.org:8883

### Linux

    go build -o step6 drone/minidrone/step6/main.go
    sudo ./step6 [dronename] ssl://iot.eclipse.org:8883

