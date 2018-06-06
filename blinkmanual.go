//ref https://www.admfactory.com/raspberrypi/

package main

import (
	"fmt"
    "github.com/stianeikeland/go-rpio"
    "os"
    "time"
)

var pin = rpio.Pin(7)

func main() {
	// Open and map memory to access gpio, check for errors
    if err := rpio.Open(); err != nil {
         fmt.Println(err)
         os.Exit(1)
    }

    // Unmap gpio memory when done
    defer rpio.Close()

    // Set pin to output mode
    pin.Output()

    for {
		/*pin.High()
         time.Sleep(time.Second)
         pin.Low()
         time.Sleep(time.Second)*/
         
         //or you can use this
         pin.Write(1)
         time.Sleep(time.Second)
         pin.Write(0)
         time.Sleep(time.Second)
    }
}
