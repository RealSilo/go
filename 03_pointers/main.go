package main

import "fmt"

func zero(z *int) {
  *z = 0 // setting the value of the memory address to 0
}

func main() {
  x := 5
  zero(&x) // memory address is passed by value
  fmt.Println(x) // x is 0
}