package inheritanceornot

import "fmt"

type person struct {
	name string
	age  int8
}

type student struct {
	person
	code   string
	course string
	room   int8
}

func main() {
	student := student{person: person{name: "t", age: 12}, code: "123", course: "go", room: 56}

	fmt.Print(student)
	fmt.Print(student.name)
}
