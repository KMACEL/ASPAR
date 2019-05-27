package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	//bmp180 := i2c.NewBMP180Driver(r)
	//adxl345 := i2c.NewADXL345Driver(r)

	
	gpioPin := gpio.NewDirectPinDriver(r, "7")

	for {
		get, err := gpioPin.DigitalRead()
		if err != nil {
			fmt.Println("Error", err.Error())
		}
		fmt.Println(get)
		time.Sleep(3 * time.Second)
	}

	/*work := func() {
		gobot.Every(1*time.Second, func() {
			t, _ := bmp180.Temperature()
			fmt.Println("Temperature", t)

			p, _ := bmp180.Pressure()
			fmt.Println("Pressure", p)

			x, y, z, err := adxl345.RawXYZ()
			if err != nil {
				fmt.Println("Error", err.Error())
			}
			fmt.Println("Deger : ", x, y, z)
		})
	}

	robot := gobot.NewRobot("Sensor Board",
		[]gobot.Connection{r},
		[]gobot.Device{bmp180},
		[]gobot.Device{adxl345},
		work,
	)

	robot.Start()*/
}
