package transport

import (
	"encoding/json"
	"net/http"
	"rest_api/internal/model"
	"rest_api/internal/service"
	"strings"
)

type DepartmentHandler struct {
	service *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}

func (h *DepartmentHandler) HandleDepartments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		departments, err := h.service.GetAllDepartments()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(departments)

	case http.MethodPost:
		var department model.Department
		if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := h.service.CreateDepartment(&department)
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

func (h *DepartmentHandler) HandleDepartmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/departments/")

	switch r.Method {
	case http.MethodGet:
		department, err := h.service.GetDepartmentById(idStr)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(department)

	case http.MethodPut:
		var department model.Department
		if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}

		updated, err := h.service.UpdateDepartment(idStr, &department)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		if err := h.service.DeleteDepartment(idStr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "metodo no disponible", http.StatusMethodNotAllowed)
	}
}
