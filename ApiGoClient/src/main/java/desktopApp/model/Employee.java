
package desktopApp.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Employee implements Model {      
    @JsonProperty(value = "emp_no", access = JsonProperty.Access.WRITE_ONLY)
    public int empNo;    
    @JsonProperty("birth_date")
    public String birthDate;    
    @JsonProperty("first_name")
    public String firstName;
    @JsonProperty("last_name")
    public String lastName;
    @JsonProperty("gender")
    public String gender;
    @JsonProperty("hire_date")
    public String hireDate;
    
    @JsonProperty(access = JsonProperty.Access.WRITE_ONLY)
    private Object[] data = new Object[6];

    // Constructor vacío obligatorio para Jackson
    public Employee() {}
       
    public Employee(int empNo) {
        this.empNo = empNo;
    }
    
    public Employee(int empNo, String birthDate, String firstName, String lastName, String gender, String hireDate) {
        this.empNo = empNo;
        this.birthDate = birthDate;
        this.firstName = firstName;
        this.lastName = lastName;
        this.gender = gender;
        this.hireDate = hireDate;                
    }    

    @Override
    public String toString() {
        return "Employee{" + "empNo=" + empNo + ", birthDate=" + birthDate + ", firstName=" + firstName + ", lastName=" + lastName + ", gender=" + gender + ", hireDate=" + hireDate + '}';
    }         
    
    @Override
    public Object[] getData() {        
        data[0] = empNo;        
        data[1] = birthDate;
        data[2] = firstName;
        data[3] = lastName;
        data[4] = gender;
        data[5] = hireDate;
        return data;
    }
    
}