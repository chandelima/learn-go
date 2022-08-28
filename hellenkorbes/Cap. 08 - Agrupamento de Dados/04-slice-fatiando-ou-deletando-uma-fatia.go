// Slice: fatiando ou deletando de uma fatia
// (SLICING A SLICE)

package main

import "fmt"

func main() {
	pizza := []string{"calabresa", "quatro queijos", "abacaxi", "frango com catupiry", "margherita"}

	fatia := pizza[0:2] //Obs.: O segundo argumento é não-inclusivo.
	fmt.Println(fatia)

	//Acessa todos os ítens do slice:
	fmt.Println(pizza[:]) // Insere automaticamente o primeiro e o último ítem do slice

	fmt.Println(pizza[:2]) // Insere automaticamente o primeiro item do slice

	fmt.Println(pizza[:len(pizza)]) //Caminho mais verboso pra acessar todos os ítens...

	fmt.Println("==================================")
	// Usando for sem slice pra imprimir todos os elementos
	for i := 0; i < len(pizza); i++ {
		fmt.Println(pizza[i])
	}

	//Excluindo o sabor abacaxi da pizza
	pizza = append(pizza[:2], pizza[3:]...)
	fmt.Println(pizza)
}
