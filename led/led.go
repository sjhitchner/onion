package main

import (
	"fmt"

	p2 "github.com/sjhitchner/onion-tutorial/pkg/onion/pins/v2plus"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	_, err := host.Init() // Init periph.io
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World")

	led := gpioreg.ByName(p2.ExpansionDock.RGBLed.Red)
	defer led.Halt() // Close when the program ends

	ledState := false
	for {
		// Keeps blinking each 500ms
		if ledState {
			led.Out(gpio.High)
		} else {
			led.Out(gpio.Low)
		}
		ledState = !ledState
		time.Sleep(500 * time.Millisecond)
	}
}
