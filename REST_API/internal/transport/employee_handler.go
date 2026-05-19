package transport

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strconv"
	"strings"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) HandleEmployees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		employees, err := h.service.GetAllEmployees()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)

	case http.MethodPost:
		var employee model.Employee
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := h.service.CreateEmployee(&employee)
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

func (h *EmployeeHandler) HandleEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "no lo encontre", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		employee, err := h.service.GetEmployeeById(id)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

	case http.MethodPut:
		var employee model.Employee
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}

		updated, err := h.service.UpdateEmployee(id, &employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.DeleteEmployee(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "metodo no disponible", http.StatusMethodNotAllowed)
	}
}
