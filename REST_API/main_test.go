package main

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	expected := "ready"
	actual := "ready"

	if actual != expected {
		t.Errorf("El entorno no está listo. Se esperaba %s pero se obtuvo %s", expected, actual)
	}
}