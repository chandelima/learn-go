// Maps: range & deletando
package main

import "fmt"

func main() {

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)

	qualquercoisa2 := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa2)

	total := 0

	for key, _ := range qualquercoisa2 {
		total += key
	}

}
