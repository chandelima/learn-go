package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
}

type dentista struct {
	pessoa
	dentesArrancados float64
	salario          float64
}

type arquiteto struct {
	pessoa
	tipoDeConstrucao string
	tamanhoDaLoucura string
}

func (x dentista) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesArrancados, "e ouve só: Bom dia!")
}

func (x arquiteto) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oiBomDia()
}

// Neste caso, implicitamente o compilador inferiu que os dois
// métodos oiBomDia() são compatíveis com a interface gente,
// e neste caso, definindo gente como um parâmetro da função
// ela pode receber os dois tipos (dentista/arquiteto)
func serHumano(g gente) {

	g.oiBomDia()
	switch g.(type) {
	case arquiteto:
		fmt.Println("O tamanho da minha loucura é:", g.(arquiteto).tamanhoDaLoucura)
	case dentista:
		fmt.Println("Eu já extraí", g.(dentista).dentesArrancados, "dentes.")
	}

}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mr. Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesArrancados: 8973,
		salario:          5999,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mr. Prédio",
			sobrenome: "Niemeyer",
			idade:     51,
		},
		tipoDeConstrucao: "Hospícios",
		tamanhoDaLoucura: "demais",
	}

	serHumano(mrdente)
	fmt.Println()
	serHumano(mrpredio)
}
