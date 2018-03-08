package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	adaptor := firmata.NewAdaptor("/dev/tty.usbmodem14131")
	led := gpio.NewLedDriver(adaptor, "13")
	button := gpio.NewButtonDriver(adaptor, "12")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{led, button},
		work,
	)

	robot.Start()
}
