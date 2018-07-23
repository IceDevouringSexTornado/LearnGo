package main

import "fmt"

const meterstoyards float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)

	Yards := meters * meterstoyards

	fmt.Println(meters, " meters is", Yards, "yards")
}
