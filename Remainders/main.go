package main

import "fmt"

func main() {

	x := 13 % 3
	fmt.Println(x)

	for i := 0; i < 100; i++ {

		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}

// Slash means divide - % gives you the remainder
// x==1 checks if x is equal to one; x=1 asigns the value of one to x
