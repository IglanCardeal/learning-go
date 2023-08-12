package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		i int8
		j int16
		k int32
		l int64
	)
	i, j, k, l = 1, 2, 3, 4
	fmt.Println(i, j, k, l)

	var num int = 2000000000000 // it will set int size according CPU architecture
	fmt.Println((num))

	var unsigned_integer uint8 = 255 // reject number signal minus
	fmt.Println((unsigned_integer))

	var x uint8 = 255

	x++                   // overflow
	fmt.Println("x: ", x) // show zero

	var y byte = 255 // byte = uint8
	fmt.Println("y: ", y)

	var r rune = 127 // rune = int32
	fmt.Println("r", r)

	var fl_num float32 = 123.45
	var fl_num2 float64 = 123.45
	fmt.Println("fl_num", fl_num)
	fmt.Println("fl_num2", fl_num2)

	char := 'A'
	fmt.Println(char) // print value 65 from ASCII table

	// zero (initial) value
	var (
		test    string // empty string
		test2   int    // 0
		err     error  // nil
		isRight bool   // false
	)
	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(err)
	fmt.Println(isRight)

	var customError error = errors.New("Custom Error")
	fmt.Println(customError)
}
