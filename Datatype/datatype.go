package main

import (
    "fmt"
    "reflect"
    "math/cmplx"
)

func main() {
    b := true
    s := ""
    n := 1
    f := 1.0
    a := []string{"foo", "bar", "baz"}
    c := 0xff
    c_1 := int64(0xfffffff)
    complex := cmplx.Sqrt(-5 + 12i)
    complex_1 := complex64(complex)


    fmt.Println(reflect.TypeOf(b))
    fmt.Println(reflect.TypeOf(s))
    fmt.Println(reflect.TypeOf(n))
    fmt.Println(reflect.TypeOf(f))
    fmt.Println(reflect.TypeOf(a))
    fmt.Println(reflect.TypeOf(c))
    fmt.Println(reflect.TypeOf(c_1))
    fmt.Println(reflect.TypeOf(complex))
    fmt.Println(reflect.TypeOf(complex_1))
}
