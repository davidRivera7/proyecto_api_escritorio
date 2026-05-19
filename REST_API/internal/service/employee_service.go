package service

import (
	"rest_api/internal/model"
	"rest_api/internal/store"
)

type EmployeeService struct {
	store store.EmployeeStore
}

func NewEmployeeService(s store.EmployeeStore) *EmployeeService {
	return &EmployeeService{store: s}
}

func (s *EmployeeService) GetAllEmployees() ([]*model.Employee, error) {
	return s.store.GetAll()
}

func (s *EmployeeService) GetEmployeeById(id int) (*model.Employee, error) {
	return s.store.GetByID(id)
}

func (s *EmployeeService) CreateEmployee(employee *model.Employee) (*model.Employee, error) {
	return s.store.Create(employee)
}

func (s *EmployeeService) UpdateEmployee(id int, employee *model.Employee) (*model.Employee, error) {
	return s.store.Update(id, employee)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.store.Delete(id)
}
