package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(2, 4)

	if total != 6 {
		t.Error("Erro no teste soma(), resultado: " + string(rune(total)))
	}
}
