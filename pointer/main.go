package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	fmt.Printf("%p, %p\n", &a, &b)

	var p *int = &a
	fmt.Printf("%p, %v\n", &p, p)
	fmt.Printf("p type: %T\n", p)

	*p = 100
	fmt.Println(a)

	// array pointer
	arr := [3]int{1, 2, 3}
	ap := &arr
	fmt.Printf("ap type: %T, ap value: %p, ap address value: %v\n", ap, ap, *ap)

	// porinter array
	arr2 := [3]*int{&a, &b, p}
	pa := &arr2
	fmt.Printf("pa type: %T, pa value: %p, pa address value: %v\n", pa, pa, *pa)
	for p := range pa {
		fmt.Printf("pa[%d]: %v, original value: %v\n", p, pa[p], *pa[p])
	}

	// pointer default value
	var p2 *int
	fmt.Printf("p2 type: %T, p2 value: %v\n", p2, p2)
	if p2 == nil {
		fmt.Println("p2 is nil")
	}
}
