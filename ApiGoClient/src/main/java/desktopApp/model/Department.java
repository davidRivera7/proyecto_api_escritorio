
package desktopApp.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Department implements Model {     
    @JsonProperty(value = "dept_no", access = JsonProperty.Access.WRITE_ONLY)
    public String deptNo;
    @JsonProperty("dept_name")
    public String deptName;    
    
    @JsonProperty(access = JsonProperty.Access.WRITE_ONLY)
    private Object[] data = new Object[2];

    public Department() {
    }
    
    public Department(String deptNo) {
        this.deptNo = deptNo;
    }

    public Department(String deptNo, String deptName) {
        this.deptNo = deptNo;
        this.deptName = deptName;
    }

    @Override
    public String toString() {
        return "Department{" + "DeptNo=" + deptNo + ", DeptName=" + deptName + '}';
    }
    
    @Override
    public Object[] getData() {        
        data[0] = deptNo;        
        data[1] = deptName;        
        return data;
    }
            
}
