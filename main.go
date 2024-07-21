// -------------------------------------------------------------------------------------------------------------------------------------------------------
// GoLight
// Firmware for lighting a WS2812 LED Strip for various chipset using TinyGo
//
// Author: olive@code7.dev
// Date: 2023-10-23
// -------------------------------------------------------------------------------------------------------------------------------------------------------
//
// This program use the WS2812 tinygo driver, a Arduino Board and a push button to switch between different light modes.
// You can adjust pin usage and light modes in the constants section below.
//
// See README.md for more information.
//
// -------------------------------------------------------------------------------------------------------------------------------------------------------
package main

import (
	"image/color"
	"machine"

	"time"

	"tinygo.org/x/drivers/ws2812"
)

// Light Modes Enumeration
type Mode uint8

// Light Modes list
const (
	WhiteHigh Mode = iota
	WhiteMedium
	WhiteLow
	WarmWhiteHigh
	WarmWhiteMedium
	WarmWhiteLow
	Turquoise
	Purple
	Blue
	Red
	AnimeOne
	AnimeTwo
	AnimeThree
)

// chip and leds constants
const (
	// modesCount is the number of modes
	modesCount = 12

	// number of LEDs in the strip
	ledCount = 120

	// pin of the ws2812 data line is connected to
	ledsPin = machine.D2

	// on card led pin
	ledPin = machine.LED

	// button digital pin
	buttonPin = machine.D3
)

// button control constants
const (
	// delay between each loop iteration
	mainLoopDelay = 100 * time.Millisecond

	// button debounce delay
	buttonDebounceDelay = 800 * time.Millisecond

	// button scan delay
	buttonScanDelay = 10 * time.Millisecond
)

var currentMode Mode = WhiteMedium
var leds [ledCount]color.RGBA
var counter uint = 0

// main function initializes the LED strip, chip and starts the main program loop
func main() {

	println("----------------------------------------")
	println(" goLight")
	println(" the light is coming.")
	println("----------------------------------------")

	println("Initializing Chip and LEDs")

	// card pin initializations
	buttonPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledsPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// initialize the LED driver
	ledsDevice := ws2812.New(ledsPin)

	println("Initializing Button control routine")
	// start button scan go routine
	go scanButton()

	println("Current light mode is :", currentMode)

	// main program loop
	println("Entering main program loop")
	for {
		DisplayCurrentMode(ledsDevice, leds[:], currentMode)
		time.Sleep(mainLoopDelay)
	}

}

// DisplayCurrentMode displays the current mode on the LED strip
// leds is the LED strip
// mode is the current mode
// device is the LED strip device
// returns nothing
func DisplayCurrentMode(device ws2812.Device, leds []color.RGBA, mode Mode) {

	switch mode {
	case WhiteHigh:
		FillColor(leds[:], color.RGBA{R: 0x11, G: 0x11, B: 0x11})
	case WhiteMedium:
		FillColor(leds[:], color.RGBA{R: 0x7f, G: 0x7f, B: 0x7f})
	case WhiteLow:
		FillColor(leds[:], color.RGBA{R: 0x3f, G: 0x3f, B: 0x3f})
	case WarmWhiteHigh:
		FillColor(leds[:], color.RGBA{R: 0xff, G: 0xff, B: 0x7f})
	case WarmWhiteMedium:
		FillColor(leds[:], color.RGBA{R: 0x7f, G: 0x7f, B: 0x3f})
	case WarmWhiteLow:
		FillColor(leds[:], color.RGBA{R: 0x3f, G: 0x3f, B: 0x1f})
	case Turquoise:
		FillColor(leds[:], color.RGBA{R: 0x00, G: 0xff, B: 0xff})
	case Purple:
		FillColor(leds[:], color.RGBA{R: 0xff, G: 0x00, B: 0xff})
	case Blue:
		FillColor(leds[:], color.RGBA{R: 0x00, G: 0x00, B: 0xff})
	case Red:
		FillColor(leds[:], color.RGBA{R: 0xff, G: 0x00, B: 0x00})
	case AnimeOne:
		FillTwoColors(leds[:], color.RGBA{R: 0xff, G: 0xff, B: 0xff}, color.RGBA{R: 0x7f, G: 0x7f, B: 0x3f})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
		FillTwoColors(leds[:], color.RGBA{R: 0xff, G: 0xff, B: 0x7f}, color.RGBA{R: 0xff, G: 0xff, B: 0xff})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)

	case AnimeTwo:
		FillThreeColors(leds[:], color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
		FillThreeColors(leds[:], color.RGBA{R: 0x00, G: 0x00, B: 0xff}, color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
		FillThreeColors(leds[:], color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff}, color.RGBA{R: 0xff, G: 0x00, B: 0x00})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
		FillThreeColors(leds[:], color.RGBA{R: 0xff, G: 0x00, B: 0x00}, color.RGBA{R: 0x00, G: 0xff, B: 0x00}, color.RGBA{R: 0x00, G: 0x00, B: 0xff})
		device.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
	}

	// write the colors array to the LED strip
	device.WriteColors(leds[:])
}

// FillColor fills the LED strip with a single color
func FillColor(leds []color.RGBA, color color.RGBA) {
	for i := range leds {
		leds[i] = color
	}
}

// FillTwoColors fills the LED strip with two alternating colors
func FillTwoColors(leds []color.RGBA, color1 color.RGBA, color2 color.RGBA) {
	for i := range leds {
		if i%2 == 0 {
			leds[i] = color1
		} else {
			leds[i] = color2
		}
	}
}

// FillThreeColors fills the LED strip with three alternating colors
func FillThreeColors(leds []color.RGBA, color1 color.RGBA, color2 color.RGBA, color3 color.RGBA) {
	for i := range leds {
		if i%3 == 0 {
			leds[i] = color1
		} else if i%3 == 1 {
			leds[i] = color2
		} else {
			leds[i] = color3
		}
	}
}

// FillGradientColors fills the LED strip with a gradient of colors between two colors
// buggy
func FillGradientColors(leds []color.RGBA, color1 color.RGBA, color2 color.RGBA) {

	// for each led fill with a color between color1 and color2
	// color is linerarly interpolated
	// color1 is the first led
	// color2 is the last led
	// color1 and color2 are included

	lenght := uint8(len(leds))

	for i := range leds {

		currentColor := color.RGBA{
			R: (color1.R*(lenght-uint8(i)) + color2.R*uint8(i)) / lenght,
			G: (color1.G*(lenght-uint8(i)) + color2.G*uint8(i)) / lenght,
			B: (color1.B*(lenght-uint8(i)) + color2.B*uint8(i)) / lenght}
		leds[i] = currentColor
	}

}

// DisplayIntegerValue displays an integer value on the LED strip by lighting up a number of LEDs
// leds is the LED strip
// color1 is the color of the lit LEDs
// value is the integer value to display
// returns nothing
func DisplayIntegerValue(leds []color.RGBA, color1 color.RGBA, value uint) {
	for i := range leds {
		if value > uint(i) {
			leds[i] = color1
		} else {
			leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}
	}
}

// scanButton scans the button state every 10 ms and changes the current mode of the LED strip accordingly
// uses a debounce delay of 800ms
// returns nothing
func scanButton() {

	var prestate uint = 1

	button := buttonPin

	for {

		buttonState := button.Get()

		if true == buttonState && prestate == 1 { // Button is up (not pressed)
			//println("Button is up (not pressed)")
			prestate = 0

		} else if false == buttonState && prestate == 0 { // On Button pressed

			//println("On Button pressed")
			// can increment counter and change mode here to not authorize multiple button presses
			prestate = 1

		} else if false == buttonState && prestate == 1 { // Button is pressed

			// increment counter and change mode here to authorize multiple button presses with the bounce delay
			counter++
			currentMode = Mode(counter % modesCount)

			println("Button Pressed ! new mode :", currentMode)
			time.Sleep(buttonDebounceDelay) // Add a debounce delay

			prestate = 1

		}

		time.Sleep(buttonScanDelay) // Adjust this delay for responsiveness

	}

}
