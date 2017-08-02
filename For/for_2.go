package main

import "fmt"
import "time"

func main() {
    condition_to_run := true
    i := 0

    // This loop continues while "valid" is true.
    for condition_to_run {
        fmt.Println(i)
        if i == 5 {
            condition_to_run = false
            fmt.Println("Exiting out of the for loop")
        }
        time.Sleep(1)
        i++
    }
}

