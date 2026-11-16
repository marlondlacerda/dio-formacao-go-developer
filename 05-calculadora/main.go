package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operacao string

	fmt.Println("=== Calculadora Go ===")

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scan(&operacao)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	var resultado float64

	switch operacao {
	case "+":
		resultado = Somar(a, b)
	case "-":
		resultado = Subtrair(a, b)
	case "*":
		resultado = Multiplicar(a, b)
	case "/":
		var err error
		resultado, err = Dividir(a, b)

		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
	default:
		fmt.Println("Operação inválida!")
		return
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
