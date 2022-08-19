// Slices multi-dimensionais

package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		// Índice: 0  1  2             // Índice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4]) // Pegando o elemento de índice 2 do slice
	// principal e o de índice 4 do subslice dimesional
}
