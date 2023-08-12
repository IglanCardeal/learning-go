package main

import (
	"errors"
	"fmt"
)

func sum(n1 int8, n2 int8) (int8, error) {
	result := n1 + n2
	if result > 10 {
		return 0, errors.New("Bigger than 10")
	}
	return result, nil
}

func main() {
	result, err := sum(5, 10)
	fmt.Println("err", err)
	fmt.Println("result", result)

	printTxt := func(txt string) {
		fmt.Println(txt)
	}
	printTxt("show me")
}
