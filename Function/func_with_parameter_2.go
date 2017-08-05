package main

import "fmt"

func main() {
fmt.Println(addMe(2,3)) 
fmt.Println(addMe(1,3)) 
fmt.Println(addMe(2,4)) 
fmt.Println(addMe(3,3)) 
}

func addMe(first int, second int   ) ( string, int) { 
  value :=  fmt.Sprintf("The sum of %d and %d is ", first, second)
  return value, first + second
}
