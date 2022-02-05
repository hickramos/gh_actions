package main

import "testing"

func TestSoma(t *testing.T) {
	resultadoEsperado := 30
	total := Soma(15, 15)

	if total != resultadoEsperado {
		t.Errorf("Resultado da soma é inválido. Resultado: %d. Esperado: %d", total, resultadoEsperado)
	}
}
