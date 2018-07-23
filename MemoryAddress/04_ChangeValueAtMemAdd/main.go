package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) //(mem location of a)

	var b *int = &a // b is referencing a memory address where an int is stored
	fmt.Println(b)  //(mem location of a)
	fmt.Println(*b) //43

	*b = 42
	fmt.Println(a) // 42
}
