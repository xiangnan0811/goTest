package main

import "fmt"

func main() {
	var courses = [5]string{"Go", "Ruby", "Javascript"}
	fmt.Println(courses)
	fmt.Printf("course[3]:%v,\n", courses[3])
	courses[3] = "Python"
	fmt.Printf("course[3]:%v,\n", courses[3])
}