/*
### Na prática: exercício #4

- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG
*/

package main

import "fmt"

func main() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Chocolate",
		sabor:   "doce",
		ondetem: []string{"mercearias", "mercadinhos", "supermercados", "padarias"},
		vaibemcom: map[string]string{
			"no café da manhã": "com café puro",
			"no almoço":        "na sobremesa",
			"no jantar":        "não vai bem, pq comer doce a noite engorda",
		},
	}

	fmt.Println(x)
}
