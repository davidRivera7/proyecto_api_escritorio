package service

import (
	"rest_api/internal/model"
	"rest_api/internal/store"
)

type DepartmentService struct {
	store store.DepartmentStore
}

func NewDepartmentService(s store.DepartmentStore) *DepartmentService {
	return &DepartmentService{store: s}
}

func (s *DepartmentService) GetAllDepartments() ([]*model.Department, error) {
	return s.store.GetAll()
}

func (s *DepartmentService) GetDepartmentById(id string) (*model.Department, error) {
	return s.store.GetByID(id)
}

func (s *DepartmentService) CreateDepartment(department *model.Department) (*model.Department, error) {
	return s.store.Create(department)
}

func (s *DepartmentService) UpdateDepartment(id string, department *model.Department) (*model.Department, error) {
	return s.store.Update(id, department)
}

func (s *DepartmentService) DeleteDepartment(id string) error {
	return s.store.Delete(id)
}
