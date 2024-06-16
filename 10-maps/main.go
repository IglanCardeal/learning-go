package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Person struct {
	name   string
	age    int
	addres Address
}

func main() {
	person := Person{name: "cardeal", age: 32, addres: Address{city: "sao paulo", state: "SP"}}
	fmt.Println(person)
	fmt.Println(person.name)

	fmt.Println("MAPS")

	my_map := map[string]map[string]string{
		"person": {
			"name": "cardeal",
			"age":  "32",
		},
		"address": {
			"city":  "sao paulo",
			"state": "SP",
		},
	}
	fmt.Println(my_map)
	fmt.Println(my_map["person"]["name"])

	delete(my_map, "person")
	fmt.Println("after delete: ", my_map)
}
