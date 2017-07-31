package main
import "fmt"
func main(){

VariableStr:="This is a Test Variable"
VariableSubstr:=VariableStr[2:5]

fmt.Println(VariableSubstr)

fmt.Println(VariableStr[:4])

fmt.Println(VariableStr[8:])

}
