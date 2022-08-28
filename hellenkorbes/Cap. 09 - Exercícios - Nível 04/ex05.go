/*
### Na prática: exercício #5

- Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]
- Solução: https://play.golang.org/p/26bT-UKmJH
*/

package main

import "fmt"

func main() {
	//          0   1   2   3   4   5   6   7   8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[0:3], x[len(x)-4:]...)
	//x = append(x[0:3], x[6:]...)     -> Podia ser assim tbm

	fmt.Println(x)
}
