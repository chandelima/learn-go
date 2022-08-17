package main

import "fmt"

func main() {

	x := 2

	// || = OU
	if x == 2 || x <= 10 {
		fmt.Println("x = 2 ou < 10")
	}

	// && = E
	if x == 2 && x <= 10 {
		fmt.Println("x = 2 e < 10")
	}
}
