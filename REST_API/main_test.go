package main

import (
	"database/sql"
	"testing"
		
	_ "github.com/go-sql-driver/mysql"
)

func TestGetEmployeesIntegration(t *testing.T) {
	// 1. Intentar conectar al MySQL que va a levantar GitHub Actions
	// El host será '127.0.0.1' y la contraseña la definiremos en el YAML
	dsn := "root:password_local@tcp(127.0.0.1:3306)/employees"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("No se pudo abrir la conexión a la base de datos: %v", err)
	}
	defer db.Close()

	// 2. Hacer una consulta real para verificar que existan datos
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM employees").Scan(&count)
	if err != nil {
		t.Fatalf("Error al ejecutar la consulta SELECT: %v", err)
	}

	// 3. Validar: Si el conteo es 0, significa que la integración falló
	if count == 0 {
		t.Errorf("Prueba de integración fallida: Se esperaba encontrar registros en la tabla 'employees', pero se obtuvo 0")
	} else {
		t.Logf("¡Éxito! La prueba de integración encontró %d empleados en la base de datos de la nube.", count)
	}
}