# Arduino Sensor Station

## What you need

    - Arduino or compatible microcontroller
    - Grove IoT Starter Kit parts
    - Personal computer with Go 1.12+ and TinyGo installed, and a serial port.

## Installation

### Go 1.12

If somehow you have not installed Go 1.12 on your computer already, you can download it here:

https://golang.org/dl/

### TinyGo

Follow the instructions here:

https://tinygo.org/getting-started/

### Bossa

In order to "flash", meaning to move the binary code from your computer to the Arduino, you must install the "bossac" command line utility whihc is part of BOSSA.

#### Linux

On Linux, install from source:

```
git clone https://github.com/shumatech/BOSSA.git
cd BOSSA
make
```

#### macOS

On macOS, download the installer from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-1.9.1.dmg

One you have downloaded it, double click on the .dmg file to perform the installation.

## Connecting the Arduino to your computer

Plug the Arduino into your computer using a USB cable. There may be one provided in your starter kit.

## Running the code

The TinyGo programs will run directly on the Arduino's microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your Arduino.
- The program executes from the Arduino. You can disconnect the Arduino from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## Code

### step0.go - Built-in LED

![Arduino](./images/step0.png)

This tests that you can compile and flash your Arduino with TinyGo code, by blinking the built-in LED.

- Click on the "RST" button two times to put the Arduino into bootloader mode so you can load your own code onto it. The LED labeled "L" should remain lit to indicate that the Arduino is ready to receive your code.

Run the following command to compile your code, and flash it onto the Arduino:

```
$ tinygo flash -target arduino-nano33 ./sensor/arduino/step0/main.go
```

Once the Arduino is flashed correctly, the built-in LED labeled "L" should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### step1.go - Blue LED

![Arduino](./images/step1.png)

- Connect the "Ground" pin on the Arduino to breadboard's blue ground rail (-) using a black jumper cable.

- Connect the "3.3V" pin on the Arduino to breadboard's red power rail (+) using a red jumper cable.

- Plug the Grove blue LED into the provided cable with the Grove connector on one end, and the male jumpers on the other. Make sure the LED itself is plugged into the Grove board.

- Connect the black male end of the Grove cable to the breadboard's blue ground rail (-) on the same side of the breadboard as the black "Ground" cable.

- Connect the red male end of the Grove cable to the breadboard's red power rail (+) on the same side of the breadboard as the red "3.3V" cable.

- Connect the yellow male end of the Grove cable to pin D12 on the Arduino.

Run the code.

```
$ tinygo flash -t arduino ./sensor/arduino/step1/main.go
```

You should see the blue LED blink.

### step2.go - Blue LED, Button

![Arduino](./images/step2.png)

- Plug the Grove Button into a provided cable with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's blue ground rail (-) on the same side of the breadboard as the black "Ground" cable.

- Connect the red male end of the Grove cable to the breadboard's red power rail (+) on the same side of the breadboard as the red "3.3V" cable.

- Connect the yellow male end of the Grove cable to pin D11 on the Arduino.

Run the code.

```
$ tinygo flash -t arduino ./sensor/arduino/step2/main.go
```

When you press the button, the blue LED should turn on.

### step3.go - Blue LED, Button, Green LED

![Arduino](./images/step3.png)

- Plug the Grove greed LED into one of the provided cable with the Grove connector on one end, and the male jumpers on the other.

- Connect the black male end of the Grove cable to the breadboard's blue ground rail (-) on the same side of the breadboard as the black "Ground" cable.

- Connect the red male end of the Grove cable to the breadboard's red power rail (+) on the same side of the breadboard as the red "3.3V" cable.

- Connect the yellow male end of the Grove cable to pin D10 on the Arduino.

Run the code.

```
$ tinygo flash -t arduino ./sensor/arduino/step3/main.go
```

The green LED should light up. When you press the button, the blue LED should turn on, and the green LED should turn off. When you release the button, the blue LED should turn off, and the green LED should turn on again.

### step4.go - Blue LED, Button, Green LED, Buzzer, Touch

![Arduino](./images/step4.png)

...

### step5.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial

![Arduino](./images/step5.png)


### step6.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED

![Arduino](./images/step6.png)


### step7.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED, Sound Sensor

![Arduino](./images/step7.png)


### step8.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED, Sound Sensor, Light Sensor

![Arduino](./images/step8.png)


### step9.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED, Sound Sensor, Light Sensor

![Arduino](./images/step9.png)

### step10.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED, Sound Sensor, Light Sensor, OLED

![Arduino](./images/step10.png)

### step11.go - Blue LED, Button, Green LED, Buzzer, Touch, Dial, Temperature, Red LED, Sound Sensor, Light Sensor, OLED, MQTT Server

![Arduino](./images/step11.png)
