
package desktopApp.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class DeptManager implements Model {
    @JsonProperty("emp_no")
    public int empNo;
    @JsonProperty(value = "first_name", access = JsonProperty.Access.WRITE_ONLY)    
    public String firstName;
    @JsonProperty(value = "last_name", access = JsonProperty.Access.WRITE_ONLY)    
    public String lastName;
    @JsonProperty("dept_no")
    public String deptNo;
    @JsonProperty(value = "dept_name", access = JsonProperty.Access.WRITE_ONLY)        
    public String deptName;
    @JsonProperty("from_date")
    public String fromDate;
    @JsonProperty("to_date")
    public String toDate;
    
    private Object[] data = new Object[7];    

    public DeptManager() {
    }    
    
    public DeptManager(int empNo, String deptNo, String fromDate, String toDate) {
        this.empNo = empNo;        
        this.deptNo = deptNo;        
        this.fromDate = fromDate;
        this.toDate = toDate;
    }

    @Override
    public String toString() {
        return "DeptManager{" + "empNo=" + empNo + ", firstName=" + firstName + ", lastName=" + lastName + ", deptNo=" + deptNo + ", deptName=" + deptName + ", fromDate=" + fromDate + ", toDate=" + toDate + '}';
    }
    
    @Override
    public Object[] getData() {        
        data[0] = empNo;        
        data[1] = firstName;
        data[2] = lastName;
        data[3] = deptNo;
        data[4] = deptName;
        data[5] = fromDate;
        data[6] = toDate;
        return data;
    }
        
}
