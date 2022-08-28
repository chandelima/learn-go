/*
### Na prática: exercício #10

- Anote (à mão) o resultado das expressões:
    - fmt.Println(true && true)
    - fmt.Println(true && false)
    - fmt.Println(true || true)
    - fmt.Println(true || false)
    - fmt.Println(!true)
- Ninja nível 3! Parabéns!
*/

package main

import "fmt"

func main() {
	fmt.Println(true && true)  // retorna true
	fmt.Println(true && false) // retorna false
	fmt.Println(true || true)  // retorna true
	fmt.Println(true || false) // retorna true
	fmt.Println(!true)         // retorna false
}
