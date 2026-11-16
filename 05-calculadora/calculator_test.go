package main

import "testing"

func TestSomar(t *testing.T) {
	resultado := Somar(10, 5)

	if resultado != 15 {
		t.Errorf("Esperado 15, recebido %v", resultado)
	}
}

func TestSubtrair(t *testing.T) {
	resultado := Subtrair(10, 5)

	if resultado != 5 {
		t.Errorf("Esperado 5, recebido %v", resultado)
	}
}

func TestMultiplicar(t *testing.T) {
	resultado := Multiplicar(10, 5)

	if resultado != 50 {
		t.Errorf("Esperado 50, recebido %v", resultado)
	}
}

func TestDividir(t *testing.T) {
	resultado, err := Dividir(10, 5)

	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if resultado != 2 {
		t.Errorf("Esperado 2, recebido %v", resultado)
	}
}

func TestDividirPorZero(t *testing.T) {
	_, err := Dividir(10, 0)

	if err == nil {
		t.Error("Esperado erro ao dividir por zero")
	}
}
