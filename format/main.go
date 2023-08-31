package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "hello"
	age := 18
    fmt.Println("name: "+name+", age: "+ strconv.Itoa(age))
	fmt.Printf("name: %#v, age: %#v\n", name, age)
	fmt.Printf("name: %T, age: %T\n", name, age)

	// return a string
	desc := fmt.Sprintf("name: %s, age: %d", name, age)
	fmt.Println(desc)

	data:= 65
	fmt.Printf("data: %c\n", data)
	fmt.Printf("data: %q\n", data)

    fmt.Printf("data: %e\n", 1.23456789)
	fmt.Printf("data: %f\n", 1.23456789)

	var n string
	fmt.Println("Please input your name:")
	fmt.Scanln(&n)
	fmt.Println("Hello, " + n)
}