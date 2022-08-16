package main

import "fmt"

func main() {
	/*
		// EXEMPLO HORAS
		for horas := 1; horas <= 12; horas++ {
			fmt.Println("Hora:", horas)
			for x := 0; x < 60; x++ {
				fmt.Print(x, ", ")
			}
			fmt.Println()
		}
	*/

	// EXEMPLO MESES
	for meses := 1; meses <= 12; meses++ {
		fmt.Printf("Mês %v:\n", meses)
		for x := 1; x <= 30; x++ {
			fmt.Print(x, ", ")
		}
		fmt.Println()
	}
}
