package main

import "fmt"

func main() {
	fmt.Println("Hello world")

  a := 43

  fmt.Println(a)
  fmt.Println(&a)

  var b *int = &a
  fmt.Println(b)
  fmt.Println(&b)
  fmt.Println(*b)

  *b = 42
  fmt.Println(a)
  fmt.Println(b)
}
