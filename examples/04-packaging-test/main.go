package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem14131")
	button1 := gpio.NewButtonDriver(firmataAdaptor, "7")
	button2 := gpio.NewButtonDriver(firmataAdaptor, "5")
	button3 := gpio.NewButtonDriver(firmataAdaptor, "3")
	led1 := gpio.NewLedDriver(firmataAdaptor, "6")
	led2 := gpio.NewLedDriver(firmataAdaptor, "4")
	led3 := gpio.NewLedDriver(firmataAdaptor, "2")

	work := func() {
		button1.On(gpio.ButtonPush, func(data interface{}) {
			led1.On()
		})

		button1.On(gpio.ButtonRelease, func(data interface{}) {
			led1.Off()
		})

		button2.On(gpio.ButtonPush, func(data interface{}) {
			led2.On()
		})

		button2.On(gpio.ButtonRelease, func(data interface{}) {
			led2.Off()
		})

		button3.On(gpio.ButtonPush, func(data interface{}) {
			led3.On()
		})

		button3.On(gpio.ButtonRelease, func(data interface{}) {
			led3.Off()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2, led3, button1, button2, button3},
		work,
	)

	robot.Start()
}
