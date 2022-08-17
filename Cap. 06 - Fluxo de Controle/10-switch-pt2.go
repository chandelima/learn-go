package main

import (
	"fmt"
)

var x interface{} //Pode receber qualquer tipo

func main() {
	x = -176

	switch x.(type) { // Neste caso o switch infere o tipo
	case int:
		fmt.Println("é um int")
	case string:
		fmt.Println("é um string")
	case bool:
		fmt.Println("é um bool")
	case float64:
		fmt.Println("é um float64")
	}

	switch x := 3; {
	case x == 1:
		fmt.Println("x é 1")
	case x == 2:
		fmt.Println("x é 2")
	case x == 3:
		fmt.Println("x é 3")
	case x == 4:
		fmt.Println("x é 4")
	case x == 5:
		fmt.Println("x é 5")
	}
}
