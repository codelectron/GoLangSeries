package main

import "fmt"

func main() {
fmt.Println(addMe(2,3)) 
}

func addMe(first int, second int   ) ( string, int) { 
  value :=  fmt.Sprintf("The sum of %d and %d is ", first, second)
  return value, first + second
}
