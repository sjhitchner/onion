package main

import (
	"fmt"
	"log"
	"time"

	// p2 "github.com/sjhitchner/onion-tutorial/pkg/onion/pins/v2plus"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	_, err := host.Init() // Init periph.io
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World")

	// led := gpioreg.ByName(p2.ExpansionDock.RGBLed.Red)
	blink("15", 500*time.Millisecond)
	blink("16", 300*time.Millisecond)
	blink("17", 100*time.Millisecond)

	<-time.After(time.Minute)
}

func blink(pin string, period time.Duration) {
	led := gpioreg.ByName(pin)
	go func() {
		defer led.Halt() // Close when the program ends
		defer led.Out(gpio.Low)

		ledState := false
		for {
			if ledState {
				led.Out(gpio.High)
			} else {
				led.Out(gpio.Low)
			}

			ledState = !ledState

			fmt.Println("LED", ledState)
			<-time.After(period)
		}
	}()
}
