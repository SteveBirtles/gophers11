package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"time"
	"math/rand"
)

var (
	mode = 0
	t = 0
)

func main() {

	firmataAdaptor := firmata.NewAdaptor("COM3")

	led := []*gpio.DirectPinDriver{
		gpio.NewDirectPinDriver(firmataAdaptor, "2"),  //0
		gpio.NewDirectPinDriver(firmataAdaptor, "3"),  //PWM 	//1
		gpio.NewDirectPinDriver(firmataAdaptor, "6"),  //PWM 	//2
		gpio.NewDirectPinDriver(firmataAdaptor, "5"),  //PWM 	//3
		gpio.NewDirectPinDriver(firmataAdaptor, "4"),  //4
		gpio.NewDirectPinDriver(firmataAdaptor, "7"),  //5
		gpio.NewDirectPinDriver(firmataAdaptor, "13"), //6
		gpio.NewDirectPinDriver(firmataAdaptor, "12"), //7
		gpio.NewDirectPinDriver(firmataAdaptor, "11"), //PWM 	//8
		gpio.NewDirectPinDriver(firmataAdaptor, "10"), //PWM 	//9
		gpio.NewDirectPinDriver(firmataAdaptor, "9"),  //PWM 	//10
		gpio.NewDirectPinDriver(firmataAdaptor, "8")} 		    //11

	work := func() {
		gobot.Every(time.Millisecond*100, func() {
			t++

			if t%30 == 0 {
				mode = (mode + 1) % 8
			}

			switch mode {
			case 0:

				for i := 0; i < 12; i++ {
					led[i].On()
				}

			case 1:

				on := rand.Intn(12)
				led[on].On()
				off := rand.Intn(12)
				led[off].Off()

			case 2:

				led[t%12].On()
				led[(t+11)%12].Off()

			case 3:

				led[t%12].On()
				led[(t+11)%12].Off()
				led[(t+4)%12].On()
				led[(t+15)%12].Off()
				led[(t+8)%12].On()
				led[(t+19)%12].Off()

			case 4:

				groups := [4][12]int{
					{1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1}}

				for i := 0; i < 12; i++ {
					if groups[t%4][i] == 1 {
						led[i].On()
					} else {
						led[i].Off()
					}
				}

			case 5:

				led[0].Off()
				led[4].Off()
				led[5].Off()
				led[6].Off()
				led[7].Off()
				led[11].Off()

				led[1].PwmWrite(byte((t * 10) % 240))
				led[2].PwmWrite(byte((t*10 + 40) % 240))
				led[3].PwmWrite(byte((t*10 + 80) % 240))
				led[8].PwmWrite(byte((t*10 + 120) % 240))
				led[9].PwmWrite(byte((t*10 + 160) % 240))
				led[10].PwmWrite(byte((t*10 + 200) % 240))

			case 6:

				led[0].On()
				led[4].On()
				led[5].On()
				led[6].On()
				led[7].On()
				led[11].On()

				led[1].PwmWrite(byte((t * 10) % 240))
				led[2].PwmWrite(byte((t * 10) % 240))
				led[3].PwmWrite(byte((t * 10) % 240))
				led[8].PwmWrite(byte((t * 10) % 240))
				led[9].PwmWrite(byte((t * 10) % 240))
				led[10].PwmWrite(byte((t * 10) % 240))

			case 7:

				for i := 0; i < 12; i++ {
					if t%2 == 0 {
						led[i].On()
					} else {
						led[i].Off()
					}
				}

			}

		})

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		work,
	)

	robot.Start()

}