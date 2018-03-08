package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// adaptor := firmata.NewAdaptor("/dev/tty.usbmodem14131")
	adaptor := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
