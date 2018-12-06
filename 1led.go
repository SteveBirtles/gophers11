package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
)

func main() {

	firmataAdaptor := firmata.NewAdaptor("COM3")

	led := gpio.NewDirectPinDriver(firmataAdaptor, "2")

	work := func() {
		led.On()
	})

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		work,
	)

	robot.Start()

}
