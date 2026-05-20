package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetEmployeesAPIIntegration(t *testing.T) {
	// 1. Conectar a la base de datos que levantará GitHub Actions de fondo
	dsn := "root:password_local@tcp(127.0.0.1:3306)/employees?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("No se pudo conectar a la BD de pruebas: %v", err)
	}
	defer db.Close()

	// 2. Encendemos los controladores y rutas reales de la API de GO usando la BD de pruebas
	apiRouter := SetupRoutes(db)

	// 3. Simular el 'curl' de Windows: Disparamos un GET real hacia la ruta de la API
	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Fatalf("No se pudo crear la petición HTTP: %v", err)
	}

	// 4. Creamos un grabador para capturar la respuesta JSON que genere mi Handler
	rr := httptest.NewRecorder()

	// 5. Enviamos la petición a través de mi enrutador real
	apiRouter.ServeHTTP(rr, req)

	// 6. Evaluamos que el código de Go responda exitosamente (Código 200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("La API devolvió un código de error: %v. Respuesta: %s", status, rr.Body.String())
	}

	// 7. Evaluamos que la respuesta contenga los datos de prueba insertados en la nube
	bodyString := rr.Body.String()
	if bodyString == "" || bodyString == "null" || bodyString == "[]" {
		t.Errorf("Error de integración: La API respondió 200, pero el JSON de empleados está vacío: %s", bodyString)
	} else {
		t.Logf("¡PRUEBA DE INTEGRACIÓN PASADA! La API ejecutó store, service y transport con éxito. Respuesta: %s", bodyString)
	}
}