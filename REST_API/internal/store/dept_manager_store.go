package store

import (
	"database/sql"
	"rest_api/internal/model"
)

type DeptManagerStore interface {
	GetAll() ([]*model.DeptManager, error)
	GetByEmployeeID(id int) ([]*model.DeptManager, error)
	Create(deptManager *model.DeptManager) (*model.DeptManager, error)
	Delete(emp_no int, dept_no string) error
}

type deptManagerStore struct {
	db *sql.DB
}

func NewDeptManagerStore(db *sql.DB) DeptManagerStore {
	return &deptManagerStore{db: db}
}

func (s *deptManagerStore) GetAll() ([]*model.DeptManager, error) {
	q := `
		select e.emp_no, e.first_name, e.last_name, d.dept_no, d.dept_name, dm.from_date, dm.to_date
		from employees as e 
		join dept_manager as dm
		on e.emp_no = dm.emp_no
		join departments as d
		on dm.dept_no = d.dept_no		
		`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departManagers []*model.DeptManager
	for rows.Next() {
		dm := model.DeptManager{}
		if err := rows.Scan(&dm.EmpNo, &dm.FirstName, &dm.LastName, &dm.DeptNo, &dm.DeptName, &dm.FromDate, &dm.ToDate); err != nil {
			return nil, err
		}

		departManagers = append(departManagers, &dm)
	}

	return departManagers, nil
}

func (s *deptManagerStore) GetByEmployeeID(id int) ([]*model.DeptManager, error) {
	q := `
		select e.emp_no, e.first_name, e.last_name, d.dept_no, d.dept_name, dm.from_date, dm.to_date
		from employees as e 
		join dept_manager as dm
		on e.emp_no = dm.emp_no
		join departments as d
		on dm.dept_no = d.dept_no	
		where e.emp_no = ?	
		`

	rows, err := s.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departManagers []*model.DeptManager
	for rows.Next() {
		dm := model.DeptManager{}
		if err := rows.Scan(&dm.EmpNo, &dm.FirstName, &dm.LastName, &dm.DeptNo, &dm.DeptName, &dm.FromDate, &dm.ToDate); err != nil {
			return nil, err
		}

		departManagers = append(departManagers, &dm)
	}

	return departManagers, nil
}

func (s *deptManagerStore) Create(deptManager *model.DeptManager) (*model.DeptManager, error) {
	q := `INSERT INTO dept_manager (dept_no, emp_no, from_date, to_date) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(q, deptManager.DeptNo, deptManager.EmpNo, deptManager.FromDate, deptManager.ToDate)
	if err != nil {
		return nil, err
	}

	return deptManager, nil
}

func (s *deptManagerStore) Delete(emp_no int, dept_no string) error {
	q := `DELETE FROM dept_manager WHERE emp_no = ? and dept_no = ?`
	_, err := s.db.Exec(q, emp_no, dept_no)
	if err != nil {
		return err
	}

	return nil
}
