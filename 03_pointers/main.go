package main

import "fmt"

func zero(z *int) {
  fmt.Println(z) // this memory address is the same like the one in func main
  *z = 0 // setting the value of the memory address to 0
}

func main() {
  x := 5
  fmt.Println(&x) // this memory address is the same like the one in func main
  zero(&x) // memory address is passed by value
  fmt.Println(x) // x is 0
}