package store

import (
	"database/sql"
	"rest_api/internal/model"
)

type DepartmentStore interface {
	GetAll() ([]*model.Department, error)
	GetByID(id string) (*model.Department, error)
	Create(department *model.Department) (*model.Department, error)
	Update(id string, department *model.Department) (*model.Department, error)
	Delete(id string) error
}

type departmentStore struct {
	db *sql.DB
}

func NewDepartmentStore(db *sql.DB) DepartmentStore {
	return &departmentStore{db: db}
}

func (s *departmentStore) GetAll() ([]*model.Department, error) {
	q := `SELECT dept_no, dept_name FROM departments LIMIT 50`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*model.Department
	for rows.Next() {
		d := model.Department{}
		if err := rows.Scan(&d.DeptNo, &d.DeptName); err != nil {
			return nil, err
		}

		departments = append(departments, &d)
	}

	return departments, nil
}

func (s *departmentStore) GetByID(id string) (*model.Department, error) {
	q := `SELECT dept_no, dept_name FROM departments WHERE dept_no = ?`

	d := model.Department{}
	err := s.db.QueryRow(q, id).Scan(&d.DeptNo, &d.DeptName)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (s *departmentStore) Create(department *model.Department) (*model.Department, error) {
	q := `SELECT CONCAT('d', LPAD(MAX(SUBSTRING(dept_no, 2)) + 1, 3, '0')) FROM departments;`
	var id string
	err := s.db.QueryRow(q).Scan(&id)
	if err != nil {
		return nil, err
	}

	q = `INSERT INTO departments (dept_no, dept_name) VALUES (?, ?)`
	_, err = s.db.Exec(q, id, department.DeptName)
	if err != nil {
		return nil, err
	}

	department.DeptNo = id
	return department, nil
}

func (s *departmentStore) Update(id string, department *model.Department) (*model.Department, error) {
	q := `UPDATE departments SET dept_name = ? WHERE dept_no = ?`

	_, err := s.db.Exec(q, department.DeptName, id)
	if err != nil {
		return nil, err
	}

	department.DeptNo = id
	return department, nil
}

func (s *departmentStore) Delete(id string) error {
	q := `DELETE FROM departments WHERE dept_no = ?`
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
