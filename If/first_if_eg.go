package main

import "fmt"

func main() {
   var check_Variable int = 100;
 
   if check_Variable < 50  {
      fmt.Printf("check_Variable is less than 20\n" );
   } else {
      fmt.Printf("check_Variable is not less than 20\n" );
   }
   fmt.Printf("value of check_Variable is : %d\n", check_Variable);
}
