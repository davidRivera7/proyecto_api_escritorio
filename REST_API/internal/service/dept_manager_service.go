package service

import (
	"rest_api/internal/model"
	"rest_api/internal/store"
)

type DeptManagerService struct {
	store store.DeptManagerStore
}

func NewDeptManagerService(s store.DeptManagerStore) *DeptManagerService {
	return &DeptManagerService{store: s}
}

func (s *DeptManagerService) GetAllDeptManagers() ([]*model.DeptManager, error) {
	return s.store.GetAll()
}

func (s *DeptManagerService) GetDeptManagerByEmployeeId(id int) ([]*model.DeptManager, error) {
	return s.store.GetByEmployeeID(id)
}

func (s *DeptManagerService) CreateDeptManager(deptManager *model.DeptManager) (*model.DeptManager, error) {
	return s.store.Create(deptManager)
}

func (s *DeptManagerService) DeleteDeptManagerBy_EmployeeID_DeptID(emp_no int, dept_no string) error {
	return s.store.Delete(emp_no, dept_no)
}
