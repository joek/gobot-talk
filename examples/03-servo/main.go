package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	adaptor := firmata.NewAdaptor("/dev/cu.usbmodem146621")
	servo := gpio.NewServoDriver(firmataAdaptor, "3")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
