package main

import (
	"machine"

	"time"
)

func main() {

	time.Sleep(time.Millisecond * 10000)

	println("Beginning...")

	led := machine.GP0 // the pin connected to the led
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	
	switchPin := machine.GP12

	switchPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		switchState := switchPin.Get()
		
		if switchState {
			println("switchPin true (door open)") // true = open
			led.Low()
		} else {
			println("switchPin false (door closed)") // false = closed
			led.High()
		}

		time.Sleep(time.Millisecond * 1000)
	
	}
}

	
	

	
	