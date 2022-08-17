package main

import "fmt"

func main() {
	x := 101

	if x > 100 {
		fmt.Println("x é maior que 100")
	} else if x == 100 {
		fmt.Println("x é igual a 100")
	} else {
		fmt.Println("x é menor que 100")
	}
}
