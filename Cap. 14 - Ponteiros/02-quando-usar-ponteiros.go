package main

import "fmt"

func main() {
	x := 11

	pathbyvalue(x)
	fmt.Println(x)

	pathbyreference(&x)
	fmt.Println(x)
}

func pathbyvalue(x int) {
	x++
}

func pathbyreference(x *int) {
	*x++
}
