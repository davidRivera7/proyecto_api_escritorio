package transport

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strconv"
	"strings"
)

type DeptManagerHandler struct {
	service *service.DeptManagerService
}

func NewDeptManagerHandler(s *service.DeptManagerService) *DeptManagerHandler {
	return &DeptManagerHandler{service: s}
}

func (h *DeptManagerHandler) HandleDeptManagers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		deptManagers, err := h.service.GetAllDeptManagers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deptManagers)

	case http.MethodPost:
		var deptManager model.DeptManager
		if err := json.NewDecoder(r.Body).Decode(&deptManager); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := h.service.CreateDeptManager(&deptManager)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *DeptManagerHandler) HandleDeptManager(w http.ResponseWriter, r *http.Request) {
	// La URL puede ser:
	//  /dept_managers/10001/d001
	// o
	//  /dept_managers/10001

	// El Path será algo como ["", "dept_managers", "10001", "d001"]
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		http.Error(w, "URL inválida. Se requiere ID de empleado y departamento", http.StatusBadRequest)
		return
	}

	var indexEmpID, indexDeptID int
	if len(parts) == 3 {
		indexEmpID = 2
	}

	if len(parts) == 4 {
		indexEmpID = 2
		indexDeptID = 3
	}

	// Obtener y convertir el emp_no
	empNo, err := strconv.Atoi(parts[indexEmpID])
	if err != nil {
		http.Error(w, "ID de empleado inválido", http.StatusBadRequest)
		return
	}
	// Obtener el dept_no
	deptNo := parts[indexDeptID]

	switch r.Method {
	case http.MethodGet:
		deptManagers, err := h.service.GetDeptManagerByEmployeeId(empNo)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deptManagers)

	case http.MethodDelete:
		if err := h.service.DeleteDeptManagerBy_EmployeeID_DeptID(empNo, deptNo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "metodo no disponible", http.StatusMethodNotAllowed)
	}
}
