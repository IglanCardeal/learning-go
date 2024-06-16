package main

import "fmt"

func main() {

  num := 10

  if num2 := num; num2 > 2 { // if nit
    fmt.Println("num2 is greater than 2")
  } else {
    fmt.Println("num2 is lower than 2")
  }

  fmt.Println()
}
