package main

import (
	"github.com/michaelfranzl/bmp180"
	"log"
	"fmt"
	"golang.org/x/exp/io/i2c"
)

func main() {
	// Fetch barometric pressure, output it for now
	fs := i2c.Devfs{Dev: "/dev/i2c-1"}
	device, err := i2c.Open(&fs, 0x77)

	if err != nil {
		log.Fatal(fmt.Sprintf("Fatal error: %+v", err))
	}

	defer device.Close()

	bmp180 := bmp180.NewSensor(device)
	id, err := bmp180.ID()
	if err != nil {
		log.Fatal(fmt.Sprintf("Fatal error: %+v", err))
	}
	temp, _ := bmp180.Temperature()
	ruf, _ := bmp180.Pressure(0)
	pres, _ := bmp180.Pressure(3)
	sea, _ := bmp180.PressureSealevel(3, 500)

	fmt.Printf("ID=0x%x,  t=%.3f°C,  pRough=%.3fmbar,  pAccurate=%.3fmbar,  pSealevel=%.3fmbar\n", id, temp, ruf, pres, sea)
}
