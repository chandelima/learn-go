package main

import "fmt"

func main() {
	x := 5

	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	}

	turno := "Joana"

	switch turno {
	case "Marquinhos":
		fmt.Println("Hoje quem está no escritório é o Marquinhos.")
	case "Zezinho":
		fmt.Println("Hoje quem está no escritório é o Zezinho.")
		fallthrough
	case "Zica":
		fmt.Println("E Zica sempre vem quando Zezinho está.")
	case "Maria":
		fmt.Println("Hoje quem está no escritório é o Maria.")
	case "Joãozinho":
		fmt.Println("Hoje quem está no escritório é o Joãozinho.")
	case "Chandinho":
		fmt.Println("Hoje quem está no escritório é o Chandinho.")
	case "Joana", "Roberto": //case composto
		fmt.Println("Hoje quem está no escritório é o casal Joana e Roberto")
	case "":
		fmt.Println("Não tem ninguém no escritório")
	default:
		fmt.Println("Não há como inferir quem está no escritório.")
	}
}
