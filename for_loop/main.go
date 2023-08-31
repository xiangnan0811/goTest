package main

import "fmt"

func main() {
	name := "hello, 向南"
	for index, value := range name {
		fmt.Printf("index: %2d, value: %c\n", index, value)
	}
}