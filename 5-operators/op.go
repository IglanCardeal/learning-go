package main

import "fmt"

// nothing special here!
func main() {
	var t int
	t = 1
	fmt.Println((t))

	verd := true
	falso := false
	fmt.Println(falso && verd)

	var num float64 = 10
	num += 9
	num /= 2
	fmt.Println(num)
}
