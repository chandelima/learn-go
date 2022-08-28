/*
### Na prática: exercício #5

- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
- Solução: https://play.golang.org/p/RkpqPpRWuo
*/

package main

import "fmt"

func main() {
	x := `variável
	
		raw string \n\n\n
		
	literal`

	fmt.Print(x)
}
