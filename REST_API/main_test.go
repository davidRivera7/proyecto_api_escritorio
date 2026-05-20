package main

import (
	"testing"
)

func TestSetupRoutesCompiles(t *testing.T) {
	db, apiRouter := conexionBDTests(t)
	defer db.Close()

	if apiRouter == nil {
		t.Fatal("El enrutador de la API se inicializó como nil")
	}
	t.Log("¡Confirmado! SetupRoutes inyecta dependencias y compila correctamente.")
}