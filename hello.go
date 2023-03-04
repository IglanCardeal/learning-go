// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"time"
// )

// var c, java, python bool

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Current time: ", time.Now())
// 	fmt.Println("Rand: ", rand.Float64())
// 	fmt.Println("My favorite number is: ", rand.Int63n(20))
// 	fmt.Println("Pi: ", math.Pi)
// 	fmt.Println("Sum(2+10): ", sum(2, 10))
// 	a, b := swap("2", "10")
// 	fmt.Println("Swap(2,10): ", a, b)
// 	x, y := split(37)
// 	fmt.Println("Split(37): ", x, y)

// 	var i int
// 	fmt.Println(i, c, java, python)
// }

// func sum(x, y int) int {
// 	return x + y
// }

// // Multiple results
// func swap(a, b string) (string, string) {
// 	return b, a
// }

// // "Naked" return
// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

package main

import "fmt"

var i, j int = 1, 2
var test string = "true"

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java, test)
}
