package main

import "fmt"

func main() {
fmt.Println( sumMe(1, 2))
fmt.Println( sumMe(1, 2, 3))
}

func sumMe(nums ...int   ) (res int) { 
    fmt.Print(nums, " ")
for _, number := range nums {
        res += number
    }
    return res 
}
