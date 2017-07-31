package main
import "fmt" 	
import "strings"

func main(){
VariableStr:="This is a Test Variable"
fmt.Println(strings.Count(VariableStr,"i"))
fmt.Println(strings.ContainsAny(VariableStr,"z"))
fmt.Println(strings.HasSuffix(VariableStr,"ble"))
fmt.Println(strings.Repeat(VariableStr,2))
fmt.Println(strings.ToLower(VariableStr))
}
