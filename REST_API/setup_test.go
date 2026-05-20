package main

import (
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// conexiónBDTests nos ayuda a centralizar la configuración de la BD de la nube de Github Actions
func conexionBDTests(t *testing.T) (*sql.DB, http.Handler) {
	// Una única conexión para todo el suite de pruebas
	dsn := "root:password_local@tcp(127.0.0.1:3306)/employees?parseTime=true"  // <-- funciona para hacer la conexión a la BD MySQL en la nube de Github Actions
	//dsn := "root:paraPHPescolar@tcp(127.0.0.1:3306)/employees?parseTime=true"  //<-- funciona para hacer la conexión a mi BD Local
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("No se pudo conectar a la BD de pruebas: %v", err)
	}

	// Creamos el enrutador con la función de main.go
	apiRouter := SetupRoutes(db)

	// Devolvemos ambos objetos listos para usarse
	return db, apiRouter
}