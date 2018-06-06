//ref https://www.admfactory.com/raspberrypi/

package main

import (
	"fmt"
    "github.com/stianeikeland/go-rpio"
    "os"
    "time"
)


// Use GPIO pin 26
var pin = rpio.Pin(26)


func main() {
	// Open and map memory to access gpio, check for errors
    if err := rpio.Open(); err != nil {
         fmt.Println(err)
         os.Exit(1)
    }

    // Unmap gpio memory when done
    defer rpio.Close()

	for {
		// Pull up and read value
		//if you use this case, 
		//you must connected which one pin button to pin digital and another pin button to ground
		pin.PullUp()
		fmt.Printf("PullUp: %d\n", pin.Read())

		// Pull down and read value
		//if you use this case, 
		//you must connected which one pin button to pin digital and another pin button to 5V or 3.3 V		
		pin.PullDown()
		fmt.Printf("PullDown: %d\n", pin.Read())
			
		time.Sleep(time.Second)
	}
}
