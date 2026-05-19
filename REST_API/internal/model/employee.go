package model

type Employee struct {
	EmpNo     int    `json:"emp_no"`
	BirthDate string `json:"birth_date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	HireDate  string `json:"hire_date"`
}
