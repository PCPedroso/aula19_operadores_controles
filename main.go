package main

import "fmt"

func main() {
	// var a, b, soma, subtracao, divisao, multiplicacao, resto int

	// a = 6
	// b = 3
	// soma = a + b
	// subtracao = a - b
	// divisao = a / b
	// multiplicacao = a * b
	// resto = a % b

	// fmt.Printf("A soma de (a) %v e (b) %v é igual a %v\n", a, b, soma)
	// fmt.Printf("A subtracao de (a) %v e (b) %v é igual a %v\n", a, b, subtracao)
	// fmt.Printf("A divisao de (a) %v e (b) %v é igual a %v\n", a, b, divisao)
	// fmt.Printf("A multiplicacao de (a) %v e (b) %v é igual a %v\n", a, b, multiplicacao)
	// fmt.Printf("A resto de (a) %v e (b) %v é igual a %v\n", a, b, resto)

	// var a, b bool

	// c := ((5 * 3) == (20 - 5))
	// d := ((5 * 3) == (20 - 2))
	// a = c
	// b = d

	// fmt.Printf("(a) %v && (b) %v = %v\n", c, d, a && b)
	// fmt.Printf("(a) %v || (b) %v = %v\n", c, d, a || b)
	// fmt.Printf("(a) %v && (b) %v = %v\n", c, !b, a && !b)
	// fmt.Printf("(!a) %v && (!b) %v = %v\n", !a, !b, !a && !b)

	a := 10

	if a < 5 {
		fmt.Println("5 é menor do que a (", a, ")")
	} else {
		fmt.Println("a (", a, ") NÃO é menor do que 5 ")
	}

	dia := "domingo"

	switch dia {
	case "segunda", "quarta", "sexta":
		fmt.Println("Hoje é segunda, quarta ou sexta feira")
	case "terça", "quinta":
		fmt.Println("Hoje é terça ou quinta feira")
	default:
		fmt.Println("Não sei que dia é hoje")
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println()

	numeros := []int{1, 2, 3, 4}

	for i, valor := range numeros {
		fmt.Printf("Indice %v com valor %v\n", i, valor)
	}

}
