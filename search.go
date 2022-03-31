package main

import "fmt"

func main() {
    name := []string {"bola","dave","naomi","dave","oluwa"}
    for x := 0; x < len(name); x++ {
    if name[x] == "dave" {
    fmt.Println ("hello Dave")
    }
    }
}
