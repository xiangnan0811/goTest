package main

import "fmt"

func main() {
	//  int byte rune float bool string - have default value
	// pointer slice map function - default value is nil

	// 1. new
    // 1.1 incorrect usage
	// var p * int
	// *p = 10
	// fmt.Println(*p)

    // 1.2 correct usage
	var p *[3]int = new([3]int)
    *p = [3]int{1, 2, 3}
	fmt.Println(*p)

    // 2. make
    var info map[string]string = make(map[string]string)
    info["name"] = "zhangsan"
    info["age"] = "20"
    fmt.Println(info)

    // the difference between `new` and `make`
    // `new`
    // 1. `new(T)` creates a variable of type `T`, initializes it to its zero value, and returns a pointer to it.
    // 2. It's primarily used to allocate memory for value types like int, float64, struct types, etc.
    // 3. The memory allocated by new is initialized to zero value of the type.
    // 4. `new` does not initialize the memory, it only zeros it.

    // `make`
    // 1. `make(T, args)` returns an initialized (not zeroed) value of type `T`, not a pointer. `make` is primarily used for slice maps and channels.
    // 2. `make` does both memory allocation and initialization.
    // 3. For slices, `make` creates an array and returns a slice reference to it.
    // 4. For maps and channels, `make` initializes the internal data structure and prepares it for use.

    // my understanding
    // 1. `new` is used to allocate memory for all types, including basic types like int, string, bool, byte, float, rune, etc and complex types like struct and array and so on. `new` always returns a pointer that pointed the new allocated memory that has been initialized to the zero-value of type. Therefore, `new` is applied to types that have clearly defined zero-value.
    // 2. `make`, however, is dedicated to creating and initializing slices, maps, and channels. These data types are treated specifically in the Go language because their internal representation contains Pointers to other memory blocks. Therefore, `make` not only allocates memory, but also initializes the internal structure of these data types.
}