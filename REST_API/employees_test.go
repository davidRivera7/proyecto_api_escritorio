package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv" 
	"testing"
)

func TestEmployeeCRUD(t *testing.T) {
	// Reutilizamos la conexión centralizada que lee de MySQL en la nube o local
	db, apiRouter := conexionBDTests(t)
	defer db.Close()        
		
	// Probar POST /employees	
	t.Run("CreateEmployee", func(t *testing.T) {
		jsonPayload := []byte(`{"birth_date":"1888-08-08","first_name":"Alison","last_name":"Flores","gender":"F","hire_date":"2028-08-08"}`)
		
		req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated && rr.Code != http.StatusOK {
			t.Errorf("POST /employees falló. Código obtenido: %v", rr.Code)
		} else {
			t.Logf("POST /employees exitosamente. Respuesta: %s", rr.Body.String())       
		}
	})          

	// Capturamos el ID que la API acaba de insertar en la BD
	var idReciente int
	err := db.QueryRow("SELECT MAX(emp_no) FROM employees").Scan(&idReciente)
	if err != nil {
		t.Fatalf("Error al obtener el ID más reciente: %v", err)
	}
	
	// Convertimos el entero a string real ("10003")
	strID := strconv.Itoa(idReciente)
			
	// Probar GET /employees/idReciente	
	t.Run("ReadEmployeeByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/employees/" + strID, nil) 
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {           
			t.Errorf("GET /employees/%s falló. Código obtenido: %v", strID, rr.Code)
		} else {
			t.Logf("GET /employees/%s exitosamente. Respuesta: %s", strID, rr.Body.String())     
		}
	})  
		
	// Probar PUT /employees/idReciente	
	t.Run("UpdateEmployee", func(t *testing.T) {
		jsonPayload := []byte(`{"birth_date":"1999-09-09","first_name":"Alison-M","last_name":"Flores-M","gender":"M","hire_date":"2016-06-06"}`)
		
		req, _ := http.NewRequest("PUT", "/employees/" + strID, bytes.NewBuffer(jsonPayload)) 
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {           
			t.Errorf("PUT /employees/%s falló. Código obtenido: %v", strID, rr.Code)
		} else {
			t.Logf("PUT /employees/%s exitosamente. Respuesta: %s", strID, rr.Body.String())     
		}
	})
	
	// Probar DELETE /employees/idReciente	
	t.Run("DeleteEmployee", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/employees/" + strID, nil) 
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("DELETE /employees/%s falló. Código obtenido: %v", strID, rr.Code)
		} else {
			t.Logf("DELETE /employees/%s exitosamente.", strID)        
		}
	})
	
	// Probar GET /employees	
	t.Run("ReadEmployees", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/employees", nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("GET /employees falló. Código obtenido: %v", rr.Code)
		}

		t.Logf("GET /employees exitosamente. Respuesta: %s", rr.Body.String())        
	})      
}