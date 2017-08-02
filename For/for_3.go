package main

import "fmt"
import "time"

func main() {
    breakpoint := 0

    // This loop continues infinitely until broken.
    for {

        fmt.Println(breakpoint)
        // Break if id is past a certain number.
        if breakpoint > 10 {
	    fmt.Println("Breaking out of the loop")
            break
        }
	time.Sleep(1)
        breakpoint += 2
    }
}
