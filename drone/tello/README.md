# DJI Tello Drone

The DJI Tello from Ryze Robotics uses a WiFi interface with a UDP-based API.

## What you need

    - DJI Tello
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go installed
    - Works on Linux (kernel v4.14+), macOS, or Windows

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Tello using the WiFi interface.

Therefore, you must connect to the Tello drone which acts as a WiFi access point before you will be able to run any of the code.

Further instructions here...

## Code

### step01/main.go

Let's start with a simple takeoff, and then land. Make sure the drone is turned on and you're connected to its wifi access point, then run the code.

```go run step1/main.go```

<hr>

### step02/main.go

The drone will hover and return some flight data info. Run this code:

```go run step2/main.go```

<hr>

### step03/main.go

**NOTE: Ctrl-C will now land the drone if you get in trouble!**

The drone can move forward, backward, to the right, and the left, all while maintaining a steady altitude. Run the code. 

```go run step3/main.go```

<hr>

### step04/main.go

The drone can perform flips while flying. Run the code.

```go run step4/main.go```

<hr>

### step05/main.go

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

* Triangle    - Takeoff
* X           -  Land
* Left stick  - altitude
* Right stick - direction

**macOS**
`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

`go run step5/main.go`

**Linux**
`sudo apt-get install libsdl2-dev`
...
<hr>

### step06/main.go

Now that you have mastered the flight controls, let's grab the drone video feed. You'll want to make sure that you have mplayer installed first. Upon running the code, you should see an mplayer window open with the camera feed.

**macOS**:
`brew install mplayer`

**Ubuntu Linux**:
`sudo apt-get install mplayer`

```go run step6/main.go```

<hr>
### keyboard/main.go

Control the tello with your keyboard!

- [, ] control take off and landing
- w, s, a, d control moving forward, backward, strafe left, and strafe right
- i, k, j, l control moving up, down, turning counter clockwise, and clockwise
- t, g, f, h control front flip, back flip, left flip, right flip
- r stop all movement on the tello to allow it to simply hover
