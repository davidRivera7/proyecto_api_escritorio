package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDepartmentCRUD(t *testing.T) {
	// Reutilizamos la conexión centralizada que lee de MySQL en la nube o local
	db, apiRouter := conexionBDTests(t)
	defer db.Close()        
		
	// Probar POST /departments	
	t.Run("CreateDepartment", func(t *testing.T) {		
		jsonPayload := []byte(`{"dept_name":"Spanish Language Service"}`)
		
		req, _ := http.NewRequest("POST", "/departments", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated && rr.Code != http.StatusOK {
			t.Errorf("POST /departments falló. Código obtenido: %v", rr.Code)
		} else {
			t.Logf("POST /departments exitosamente. Respuesta: %s", rr.Body.String())       
		}
	})          

	// Obtenemos el dept_no del departamento que se acaba de crear dinámicamente
	var idReciente string
	err := db.QueryRow("SELECT dept_no FROM departments ORDER BY dept_no DESC LIMIT 1").Scan(&idReciente)
	if err != nil {
		t.Fatalf("Error al obtener el dept_no más reciente de la BD: %v", err)
	}
			
	// Probar GET /departments/{id}	
	t.Run("ReadDepartmentByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/departments/" + idReciente, nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {           
			t.Errorf("GET /departments/%s falló. Código obtenido: %v", idReciente, rr.Code)
		} else {
			t.Logf("GET /departments/%s exitosamente. Respuesta: %s", idReciente, rr.Body.String())     
		}
	})  
		
	// Probar PUT /departments/{id}	
	t.Run("UpdateDepartment", func(t *testing.T) {
		// JSON basado en tu curl de actualización
		jsonPayload := []byte(`{"dept_name":"Android Service"}`)
		
		req, _ := http.NewRequest("PUT", "/departments/" + idReciente, bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {           
			t.Errorf("PUT /departments/%s falló. Código obtenido: %v", idReciente, rr.Code)
		} else {
			t.Logf("PUT /departments/%s exitosamente. Respuesta: %s", idReciente, rr.Body.String())     
		}
	})
	
	// Probar DELETE /departments/{id}	
	t.Run("DeleteDepartment", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/departments/" + idReciente, nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
			t.Errorf("DELETE /departments/%s falló. Código obtenido: %v", idReciente, rr.Code)
		} else {
			t.Logf("DELETE /departments/%s exitosamente.", idReciente)        
		}
	})
	
	// Probar GET /departments	
	t.Run("ReadDepartments", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/departments", nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("GET /departments falló. Código obtenido: %v", rr.Code)
		} else {
			t.Logf("GET /departments exitosamente. Respuesta: %s", rr.Body.String())        
		}
	})      
}