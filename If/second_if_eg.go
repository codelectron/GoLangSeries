package main
import "fmt"
import "strings"

func main() {
    fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)
    if strings.Contains(input, "This") { 

    fmt.Println("The string exist")
   }else {
    fmt.Println("This string does not exist")
   } 
}
