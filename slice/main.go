package main

import "fmt"

func main() {
	var courses = [8]string{1: "Go", 4: "Ruby", 5: "Javascript"}
	for i, course := range courses {
		fmt.Printf("courses[%d]: %v\n", i, course)
	}

	fmt.Println("slice from array...")
	var courseSlice = courses[1:5]
	fmt.Println(courseSlice)
	fmt.Printf("courseSlice length: %d, capacity: %d\n", len(courseSlice), cap(courseSlice))

	fmt.Println("slice from slice...")
	var courseSlice2 = courseSlice[:4]
	fmt.Println(courseSlice2)
	fmt.Printf("courseSlice2 length: %d, capacity: %d\n", len(courseSlice2), cap(courseSlice2))

	fmt.Println("append to slice2...")
	courseSlice2 = append(courseSlice2, "Python")
	// courseSlice2 = append(courseSlice2, "Python", "Java", "C++", "C#")
	fmt.Printf("courseSlice: %v\n", courseSlice)
	fmt.Printf("courseSlice length: %d, capacity: %d\n", len(courseSlice), cap(courseSlice))
	fmt.Printf("courseSlice2: %v\n", courseSlice2)
	fmt.Printf("courseSlice2 length: %d, capacity: %d\n", len(courseSlice2), cap(courseSlice2))

	fmt.Println("modify array 1...")
	courses[1] = "Rust"
	fmt.Printf("courseSlice: %v\n", courseSlice)
	fmt.Printf("courseSlice length: %d, capacity: %d\n", len(courseSlice), cap(courseSlice))
	fmt.Printf("courseSlice2: %v\n", courseSlice2)
	fmt.Printf("courseSlice2 length: %d, capacity: %d\n", len(courseSlice2), cap(courseSlice2))
}