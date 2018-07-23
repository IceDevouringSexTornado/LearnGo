package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var *int b = &a
	fmt.Println(b)
	fmt.Println(*b)
}
