package main

import "fmt"

func main() {
   age := 20

   if (age > 21) {
   fmt.Println("you can enter the bar because you are old enough")
   } else if age == 21 {
   fmt.Println("danizzy you may enter the bar")
   } else {
   fmt.Println("you shall not enter")
   }
}
