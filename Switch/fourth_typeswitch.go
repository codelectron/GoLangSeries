package main
import "fmt"


func main() {
    fmt.Println(typeName(1))
    fmt.Println(typeName("1"))
}
func typeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
