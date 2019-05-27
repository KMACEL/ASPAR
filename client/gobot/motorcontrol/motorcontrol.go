package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	//led := gpio.NewLedDriver(r, "3")
	gpioPin := gpio.NewDirectPinDriver(r, "3")

	for {
		gpioPin.DigitalWrite(1)
		time.Sleep(3 * time.Second)
		gpioPin.DigitalWrite(0)
		time.Sleep(3 * time.Second)
	}

	/*work := func() {
		gobot.Every(1*time.Second, func() {
			led
		})
	}*/

	/*robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		//[]gobot.Device{led},
		led,
	)*/

	//robot.Start()
}
