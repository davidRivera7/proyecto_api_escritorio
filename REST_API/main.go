package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"rest_api/internal/service"
	"rest_api/internal/store"
	"rest_api/internal/transport"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configuración de conexión MySQL
	usuario := "root"
	password := "mi_password_secreto"  
	//password := "paraPHPescolar" //<-- funciona para mi BD MySQL local
	host := "db"
	//host := "localhost" //<-- funciona para mi BD MySQL local
	puerto := "3306"	
	nombreBD := "employees" 

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		usuario, password, host, puerto, nombreBD)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error abriendo base de datos:", err)
	}
	defer db.Close()

	// Intentar conectar varias veces esperando a que MySQL despierte
    var pingErr error
    for i := 0; i < 10; i++ {
        pingErr = db.Ping()
        if pingErr == nil {
            break
        }
        log.Printf("Aún no se pudo conectar a MySQL (intento %d/10). Esperando...", i+1)
        time.Sleep(3 * time.Second) 
    }

    if pingErr != nil {
        log.Fatal("No se pudo conectar a MySQL definitivamente:", pingErr)	
	}

	// Inyectar nuestras dependencias
	//Employee
	employeeStore := store.NewEmployeeStore(db)
	employeeService := service.NewEmployeeService(employeeStore)
	employeeHandler := transport.NewEmployeeHandler(employeeService)

	//Department
	departmentStore := store.NewDepartmentStore(db)
	departmentService := service.NewDepartmentService(departmentStore)
	departmentHandler := transport.NewDepartmentHandler(departmentService)

	//DeptManager
	deptManagerStore := store.NewDeptManagerStore(db)
	deptManagerService := service.NewDeptManagerService(deptManagerStore)
	deptManagerHandler := transport.NewDeptManagerHandler(deptManagerService)

	// Configurar las rutas
	//Employee
	http.HandleFunc("/employees", employeeHandler.HandleEmployees)
	http.HandleFunc("/employees/", employeeHandler.HandleEmployeeByID)

	//Department
	http.HandleFunc("/departments", departmentHandler.HandleDepartments)
	http.HandleFunc("/departments/", departmentHandler.HandleDepartmentByID)

	//DeptManager
	http.HandleFunc("/dept_managers", deptManagerHandler.HandleDeptManagers)
	http.HandleFunc("/dept_managers/", deptManagerHandler.HandleDeptManager)

	fmt.Println("🚀 Servidor ejecutándose en http://localhost:8080")
	fmt.Println("📚 API Endpoints:")
	fmt.Println("")
	fmt.Println("   EMPLOYEES")
	fmt.Println("   GET    /employees      - Obtener todos los empleados")
	fmt.Println("   POST   /employees      - Crear un nuevo empleado")
	fmt.Println("   GET    /employees/{id} - Obtener un empleado específico")
	fmt.Println("   PUT    /employees/{id} - Actualizar un empleado")
	fmt.Println("   DELETE /employees/{id} - Eliminar un empleado")
	fmt.Println("")
	fmt.Println("   DEPARTMENTS")
	fmt.Println("   GET    /departments      - Obtener todos los departamentos")
	fmt.Println("   POST   /departments      - Crear un nuevo departamento")
	fmt.Println("   GET    /departments/{id} - Obtener un departamento específico")
	fmt.Println("   PUT    /departments/{id} - Actualizar un departamento")
	fmt.Println("   DELETE /departments/{id} - Eliminar un departamento")
	fmt.Println("")
	fmt.Println("   DEPARTMENTS")
	fmt.Println("   GET    /dept_managers			- Obtener todos los gerentes de departamento")
	fmt.Println("   POST   /dept_managers			- Crear un nuevo gerente de departamento")
	fmt.Println("   GET    /dept_managers/{emp_id}	- Obtener los departamentos que maneja un empleado específico")
	fmt.Println("   DELETE /dept_managers/{emp_id}/{dept_id} - Eliminar un gerente de departamento específico")

	// Empezar y escuchar al servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}
