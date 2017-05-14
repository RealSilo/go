package main

import "fmt"

func zero(x int) {
  fmt.Println(&x) // this memory address is different from the one in main
  x = 0 // 5 wass passed in
}

func main() {
  x := 5
  fmt.Println(&x) // this memory address is different from the one in zero
  zero(x) // passed by value (5 is passed)
  fmt.Println(x) // x is still 5
}
