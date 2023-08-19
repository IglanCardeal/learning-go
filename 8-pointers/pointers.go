package main

import "fmt"

func main() {
	var value float64 = 10
	var pointer *float64 // default is null
	fmt.Println(value, pointer)

	pointer = &value
	fmt.Println(value, pointer)

	// *pointer -> pointer dereference, show store address value
	fmt.Println(value, *pointer)
}
