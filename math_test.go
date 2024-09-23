package main

import "testing"

func TestSo(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Error("Resultado da doma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
