package model

type DeptManager struct {
	EmpNo     int    `json:"emp_no"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	DeptNo    string `json:"dept_no"`
	DeptName  string `json:"dept_name,omitempty"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
}
