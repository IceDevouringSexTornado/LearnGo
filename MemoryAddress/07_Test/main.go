package main

import "fmt"

func main() {
	var mem float64
	fmt.Print("Enter number to store in memory: ")
	fmt.Scan(&mem)

	fmt.Println("The memory location of this number is: ", &mem)
}
