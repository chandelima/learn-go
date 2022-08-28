/*
### Na prática: exercício #6

- Crie e utilize uma função anônima.
- Solução: https://play.golang.org/p/Kgo6hVr5G5
*/

package main

import "fmt"

func main() {
	x := func(n int) int {
		return n * 10
	}(10)

	fmt.Println(x)
}
