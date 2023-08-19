package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("error")
	var pointer *error
	pointer = &err
	var arr [3]error

	fmt.Println(arr, err, pointer)

	ints := [2]int{1, 3}
	fmt.Println(ints)

	floats := [...]float32{1, 2, 3, 4, 5, 6}
	fmt.Println(floats)

	slice := []int{1, 2, 3, 45, 5}
	fmt.Println(slice)

	fmt.Println("TypeOf(slice)", reflect.TypeOf(slice))
	fmt.Println("TypeOf(arr)", reflect.TypeOf(arr))

	slice2 := floats[1:4]
	fmt.Println(slice2)

	any_slice := []int{1, 4, 5, 67, 8}
	var slice_pointer *[]int
	slice_pointer = &any_slice

	fmt.Println("any_slice", any_slice)
	fmt.Println("slice_pointer", slice_pointer)

	any_slice = append(any_slice, 12)

	fmt.Println("any_slice", any_slice)
	fmt.Println("slice_pointer", slice_pointer)

	arr2 := [3]string{"test1", "test2", "test3"}
	slice_arr2 := arr2[0:2]
	fmt.Println("arr2", arr2)
	fmt.Println("slice_arr2", slice_arr2)

	slice_arr2[0] = "Cardeal"
	fmt.Println("arr2", arr2)
	fmt.Println("slice_arr2", slice_arr2)
}
