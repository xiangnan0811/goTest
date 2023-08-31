package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Index("hello", "x"))

    fmt.Println(strings.Contains("hello", "x"))
    
    fmt.Println(strings.HasPrefix("hello", "hx"))
    fmt.Println(strings.HasSuffix("hello", "hx"))
    
    fmt.Println(strings.Count("hello", "l"))
    
    fmt.Println(strings.ToUpper("hello"))

    fmt.Println(strings.TrimLeft("hello", "he"))
    fmt.Println(strings.Trim("hello", "h"))
    fmt.Println(strings.TrimSpace(" hello      "))

    fmt.Println(strings.Split("hello,world", ","))

    fmt.Println(strings.Join([]string{"hello", "world"}, " , "))

    fmt.Println(strings.Replace("hello,world", "l", "x", -2))
}
