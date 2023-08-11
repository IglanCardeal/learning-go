package main

import "fmt"

func main() {
	var text string
	text2 := "hello, world2!" // := is for declaration + assignment
	text = "hello, world!"    // = is for assignment only

	fmt.Println(text)
	fmt.Println(text2)

	var (
		var1 string = "test"
		var2 int    = 1
	)
	fmt.Println(var1, var2)

	hello, world := "hello", "world"
	const constant string = "constant"
	fmt.Println(hello, world, constant)
}
