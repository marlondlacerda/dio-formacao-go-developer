package main

import "fmt"

func main() {
	var kelvin float64

	fmt.Println("=== Conversor de Kelvin para Celsius ===")
	fmt.Print("Digite a temperatura em Kelvin: ")

	fmt.Scan(&kelvin)

	celsius := kelvin - 273

	fmt.Printf("%.2f Kelvin equivalem a %.2f Celsius\n", kelvin, celsius)
}
