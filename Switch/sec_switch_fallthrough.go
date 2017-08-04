package main
import "fmt"


func main() {
    var input int 
    fmt.Scanln(&input)

    switch int(input)  {
    case 1:
        fmt.Println("one")
	fallthrough
    case 2:
        fmt.Println("two")
	fallthrough
    case 3:
        fmt.Println("three")
	fallthrough
    default:
        fmt.Println("Unhandled at this moment")
    }
}
