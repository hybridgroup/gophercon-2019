# Sphero SPRK+

The Sphero SPRK+ and Sphero Ollie, and Sphero BB-8 all use the same API. However,
they have separate Gobot drivers to accommodate their other differences.

## What you need

    - Sphero Ollie, SPRK+, or BB-8
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux or OS X

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your SPRK+, Ollie or BB-8 as needed.

### OS X

To run any of the Gobot BLE code on OS X, you must use the `GODEBUG=cgocheck=0` flag.

For example:

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step01/main.go BB-128E
```

### Linux

On Linux the BLE code will need to run as a root user account. The easiest way to accomplish this is probably to use `go build` to build your program, and then to run the requesting executable using `sudo`.

For example:

```
$ go build -o step01 rover/ollie/step01/main.go
$ sudo ./step01 2B-123E
```

## Code

### step01

This tests that the Sphero SPRK+ or Ollie is connected correctly to your computer, by blinking the built-in LED.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step01/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step01/main.go 2B-1234
```

#### Linux

```
$ go build -o step01 rover/sprkplus/step01/main.go
$ sudo ./step01 SK-1234
```

OR

```
$ go build -o step01 rover/ollie/step01/main.go
$ sudo ./step01 2B-1234
```

### step02

Rolls around at random.


#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step02/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step02/main.go 2B-1234
```

#### Linux

```
$ go build -o step02 rover/sprkplus/step02/main.go
$ sudo ./step02 SK-1234
```

OR

```
$ go build -o step02 rover/ollie/step02/main.go
$ sudo ./step02 2B-1234
```

### step03

Gets collision notifications from robot.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step03/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step03/main.go 2B-1234
```

#### Linux

```
$ go build -o step03 rover/sprkplus/step03/main.go
$ sudo ./step03 SK-1234
```

OR

```
$ go build -o step03 rover/ollie/step03/main.go
$ sudo ./step03 2B-1234
```

### step04/main.go

This step has us receiving a heartbeat signal from the "base station" using the MQTT machine to machine messaging protocol. No additional hardware needs to be connected. 

You will need the server location of the MQTT server to use for the base station.

When the heartbeat data is received from the base station, the built-in LED will change color.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step04/main.go SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step04/main.go 2B-1234 tcp://192.168.1.55:1883
```

#### Linux

```
$ go build -o step04 rover/sprkplus/step04/main.go
$ sudo ./step04 SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ go build -o step04 rover/ollie/step04/main.go
$ sudo ./step04 2B-1234 tcp://192.168.1.55:1883
```

### step05/main.go

Control robot using keyboard arrow keys.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step05/main.go SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step05/main.go 2B-1234 tcp://192.168.1.55:1883
```

#### Linux

```
$ go build -o step05 rover/sprkplus/step05/main.go
$ sudo ./step05 SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ go build -o step05 rover/ollie/step05/main.go
$ sudo ./step05 2B-1234 tcp://192.168.1.55:1883
```

### step06/main.go

Control robot using keyboard to collect data and send to base station.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step06/main.go SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step06/main.go 2B-1234 tcp://192.168.1.55:1883
```

#### Linux

```
$ go build -o step06 rover/sprkplus/step06/main.go
$ sudo ./step06 SK-1234 tcp://192.168.1.55:1883
```

OR

```
$ go build -o step06 rover/ollie/step06/main.go
$ sudo ./step06 2B-1234 tcp://192.168.1.55:1883
```

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
