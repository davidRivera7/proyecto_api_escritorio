package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv" 
	"testing"
)

func TestDeptManagerCRUD(t *testing.T) {
	// Reutilizamos la conexión centralizada
	db, apiRouter := conexionBDTests(t)
	defer db.Close()        

	// Buscamos un emp_no y un dept_no válidos que ya existan en las tablas base
	var idEmpleado int
	var idDepartamento string

	errEmp := db.QueryRow("SELECT emp_no FROM employees LIMIT 1").Scan(&idEmpleado)
	errDept := db.QueryRow("SELECT dept_no FROM departments LIMIT 1").Scan(&idDepartamento)

	if errEmp != nil || errDept != nil {
		t.Fatalf("Error al preparar datos semilla: EmpErr: %v, DeptErr: %v", errEmp, errDept)
	}

	// Convertimos el emp_no a string para poder concatenar las URLs del API
	strEmpID := strconv.Itoa(idEmpleado)
	
	// Probar POST /dept_managers	
	t.Run("CreateDeptManager", func(t *testing.T) {
		// Construimos el JSON dinámicamente usando los IDs reales de la BD
		jsonPayload := []byte(`{
			"dept_no": "` + idDepartamento + `",
			"emp_no": ` + strEmpID + `,
			"from_date": "2020-01-01",
			"to_date": "2026-01-01"
		}`)
		
		req, _ := http.NewRequest("POST", "/dept_managers", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		// Nota: Si la API devuelve StatusConflict (409) porque esa relación ya existía, 
		// también lo consideramos válido en entornos de prueba, pero usualmente con BD limpia da OK o Created.
		if rr.Code != http.StatusCreated && rr.Code != http.StatusOK {
			t.Errorf("POST /dept_managers falló. Código obtenido: %v. Body: %s", rr.Code, rr.Body.String())
		} else {
			t.Logf("POST /dept_managers exitosamente. Respuesta: %s", rr.Body.String())       
		}
	})          
			
	// Probar GET /dept_managers/{emp_id}	
	t.Run("ReadDeptManagerByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/dept_managers/" + strEmpID, nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {           
			t.Errorf("GET /dept_managers/%s falló. Código obtenido: %v", strEmpID, rr.Code)
		} else {
			t.Logf("GET /dept_managers/%s exitosamente. Respuesta: %s", strEmpID, rr.Body.String())     
		}
	})  
	
	// Probar DELETE /dept_managers/{emp_id}/{dept_id}	
	t.Run("DeleteDeptManager", func(t *testing.T) {		
		rutaDelete := "/dept_managers/" + strEmpID + "/" + idDepartamento
		
		req, _ := http.NewRequest("DELETE", rutaDelete, nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
			t.Errorf("DELETE %s falló. Código obtenido: %v", rutaDelete, rr.Code)
		} else {
			t.Logf("DELETE %s exitosamente.", rutaDelete)        
		}
	})
	
	// Probar GET /dept_managers	
	t.Run("ReadDeptManagers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/dept_managers", nil)
		rr := httptest.NewRecorder()

		apiRouter.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("GET /dept_managers falló. Código obtenido: %v", rr.Code)
		} else {
			t.Logf("GET /dept_managers exitosamente. Respuesta: %s", rr.Body.String())        
		}
	})      
}