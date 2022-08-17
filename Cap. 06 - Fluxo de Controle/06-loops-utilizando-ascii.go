package main

import "fmt"

func main() {
	// LOOPS USANDO ASCII
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %c\n", x, x)
	}
}
