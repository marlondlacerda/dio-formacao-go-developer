package main

import "errors"

func Somar(a, b float64) float64 {
	return a + b
}

func Subtrair(a, b float64) float64 {
	return a - b
}

func Multiplicar(a, b float64) float64 {
	return a * b
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}

	return a / b, nil
}
