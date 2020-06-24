package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("Code.education Rocks!")

	if result != "<b>Code.education Rocks!</b>" {
		t.Error("Resultado inválido. Experado: Code.education Rocks!")
	}
}
