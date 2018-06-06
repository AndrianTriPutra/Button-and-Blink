package main

import (
	"fmt"
    "github.com/stianeikeland/go-rpio"
    "os"
)

var pinButton = rpio.Pin(26)
var pinLED = rpio.Pin(7)

func main() {
    if err := rpio.Open(); err != nil {
                fmt.Println(err)
                os.Exit(1)
    }
    
    defer rpio.Close()
	pinLED.Output()
	
	for {			
		pinButton.PullDown()
		fmt.Printf("PullDown: %d\n", pinButton.Read())
			
		if pinButton.Read()==1{
			pinLED.High()
		}else{
			pinLED.Low()
		}
	}
}
