//if you want to use Pulldown with hardware not firmware, you can use this firmware

package main

import (
	"fmt"
    "github.com/stianeikeland/go-rpio"
    "os"
)

var pinButton = rpio.Pin(26)
var pinLED = rpio.Pin(7)

func main() {
	// Open and map memory to access gpio, check for errors
    if err := rpio.Open(); err != nil {
         fmt.Println(err)
         os.Exit(1)
    }

    // Unmap gpio memory when done
    defer rpio.Close()

    // Set pin to output mode
    pinButton.Input()
	pinLED.Output()

    for {
		fmt.Printf("PullDown: %d\n", pinButton.Read())
		
		if pinButton.Read()==1{
			pinLED.High()
		}else{
			pinLED.Low()
		}		
    }
}
