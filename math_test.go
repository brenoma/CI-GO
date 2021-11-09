package main

import "testing"

func testSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválida, Resultado: %d - Esperado: %d", total, 30)
	}
}
