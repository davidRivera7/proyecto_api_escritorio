package store

import (
	"database/sql"
	"rest_api/internal/model"
)

type EmployeeStore interface {
	GetAll() ([]*model.Employee, error)
	GetByID(id int) (*model.Employee, error)
	Create(libro *model.Employee) (*model.Employee, error)
	Update(id int, libro *model.Employee) (*model.Employee, error)
	Delete(id int) error
}

type employeeStore struct {
	db *sql.DB
}

func NewEmployeeStore(db *sql.DB) EmployeeStore {
	return &employeeStore{db: db}
}

func (s *employeeStore) GetAll() ([]*model.Employee, error) {
	q := `SELECT emp_no, birth_date, first_name, last_name, gender, hire_date FROM employees LIMIT 50`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*model.Employee
	for rows.Next() {
		e := model.Employee{}
		if err := rows.Scan(&e.EmpNo, &e.BirthDate, &e.FirstName, &e.LastName, &e.Gender, &e.HireDate); err != nil {
			return nil, err
		}

		employees = append(employees, &e)
	}

	return employees, nil
}

func (s *employeeStore) GetByID(id int) (*model.Employee, error) {
	q := `SELECT emp_no, birth_date, first_name, last_name, gender, hire_date FROM employees WHERE emp_no = ?`

	e := model.Employee{}
	err := s.db.QueryRow(q, id).Scan(&e.EmpNo, &e.BirthDate, &e.FirstName, &e.LastName, &e.Gender, &e.HireDate)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (s *employeeStore) Create(employee *model.Employee) (*model.Employee, error) {
	q := `SELECT COALESCE(MAX(emp_no), 0) + 1 FROM employees;`
	var id int
	err := s.db.QueryRow(q).Scan(&id)
	if err != nil {
		return nil, err
	}

	q = `INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = s.db.Exec(q, id, employee.BirthDate, employee.FirstName, employee.LastName, employee.Gender, employee.HireDate)
	if err != nil {
		return nil, err
	}

	employee.EmpNo = id
	return employee, nil
}

func (s *employeeStore) Update(id int, employee *model.Employee) (*model.Employee, error) {
	q := `UPDATE employees SET birth_date = ?, first_name = ?, last_name = ?, gender = ?, hire_date = ? WHERE emp_no = ?`

	_, err := s.db.Exec(q, employee.BirthDate, employee.FirstName, employee.LastName, employee.Gender, employee.HireDate, id)
	if err != nil {
		return nil, err
	}

	employee.EmpNo = id
	return employee, nil
}

func (s *employeeStore) Delete(id int) error {
	q := `DELETE FROM employees WHERE emp_no = ?`
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
