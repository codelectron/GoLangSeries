package main

import "fmt"


func main() {
	callMe()
}
func callMe() {
	defer fmt.Println("world")
	fmt.Println("hello")
	for i := 0; i < 5; i++ {
 	   defer fmt.Printf("%d ", i)
	}
}
