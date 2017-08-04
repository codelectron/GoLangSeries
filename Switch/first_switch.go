package main
import "fmt"


func main() {
    var input int 
    fmt.Scanln(&input)

    switch int(input)  {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("Unhandled at this moment")
    }
}
