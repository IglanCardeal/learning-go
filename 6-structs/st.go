package main

import "fmt"

type address struct {
	street  string
	state   string
	country string
}

type user struct {
	name    string
	age     int8
	address address
}

func main() {
	var user_data user
	user_data.age = 10
	user_data.name = "test"

	fmt.Println(user_data.name)
	fmt.Println(user_data.age)

	breakLine()

	user_address := address{street: "test", state: "PA", country: "Lizarb"}

	other_user := user{"test2", 20, user_address}
	print(other_user)
	other_user = user{name: "test3"}
	print(other_user)
	other_user = user{"test4", 17, user_address}
	print(other_user)
}

func print(param any) {
	fmt.Println(param)
}

func breakLine() {
	fmt.Println("======================")
}
