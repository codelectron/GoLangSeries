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
    case 4,5,6,7:
        fmt.Println("Anytting between 4 to 7")
    default:
        fmt.Println("Unhandled at this moment")
    }
}
