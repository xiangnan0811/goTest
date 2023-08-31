package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
    Name string
    Slug string
    Url  string
    Price float64
    Author []string
}

func (c Course) printCourseInfo() string {
    return fmt.Sprintf("Course Name: %s\nCourse Slug: %s\nCourse Url: %s\nCourse Price: %.2f\n", c.Name, c.Slug, c.Url, c.Price)
}

func (c Course) modifyPrice(price float64) {
    c.Price = price
}

func (c *Course) setPrice(price float64) {
    c.Price = price
}

func main() {
    c := Course {
        Name: "Golang: Exploring the Fundamentals",
        Slug: "golang-fundamentals",
        Url: "https://www.pluralsight.com/courses/golang-fundamentals",
        Price: 99.99,
    }
    fmt.Println(c)

    var c2 Course
    fmt.Println(c2)

    var c3 *Course = new(Course)
    fmt.Println(c3)

    var c4 *Course = &Course{}
    fmt.Println(c4)

    // struct is value type
    c5 := c
    c5.Name = "Python: Exploring the Fundamentals"
    c.Price = 199.99
    fmt.Println(c)
    fmt.Println(c5)

    fmt.Println("-------------- before modify price of c --------------")
    fmt.Println(c.Price)
    fmt.Println("-------------- modify price of c --------------")
    c.modifyPrice(299.99)
    fmt.Println(c.Price)
    fmt.Println("-------------- modify price of c by pointer --------------")
    c.setPrice(399.99)
    fmt.Println(c.Price)

    // the size of struct
    fmt.Println("the size of struct is", unsafe.Sizeof(c))

    // the method of struct
    fmt.Println("-------------- the method of struct --------------")
    res := c.printCourseInfo()
    fmt.Printf("%T\n", res)
    fmt.Println(res)
}