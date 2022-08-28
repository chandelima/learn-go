package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct 2 JSON
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//JSON 2 struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "Violão Gianini", "preco": 600.00, "tags": ["instrumento", "musica"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
