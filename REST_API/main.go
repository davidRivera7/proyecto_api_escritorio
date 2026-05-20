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

// SetupRoutes configura y devuelve el manejador con todas las rutas de la API
func SetupRoutes(db *sql.DB) http.Handler {
	// 1. Inyectar nuestras dependencias usando la BD que nos pasen
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

	// 2. Usamos un ServeMux local en lugar del global para poder probarlo aisladamente
	mux := http.NewServeMux()

	// Configurar las rutas
	//Employee
	mux.HandleFunc("/employees", employeeHandler.HandleEmployees)
	mux.HandleFunc("/employees/", employeeHandler.HandleEmployeeByID)

	//Department
	mux.HandleFunc("/departments", departmentHandler.HandleDepartments)
	mux.HandleFunc("/departments/", departmentHandler.HandleDepartmentByID)

	//DeptManager
	mux.HandleFunc("/dept_managers", deptManagerHandler.HandleDeptManagers)
	mux.HandleFunc("/dept_managers/", deptManagerHandler.HandleDeptManager)	

	return mux
}

func main() {
	// Configuración de conexión MySQL
	usuario := "root"
	password := "mi_password_secreto"  // <-- funciona para la BD MySQL del contenedor Docker
	//password := "paraPHPescolar" //<-- funciona para mi BD MySQL local
	host := "db" // <-- funciona para la BD MySQL del contenedor Docker
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

	//PASAMOS LAS RUTAS AL SERVIDOR
	router := SetupRoutes(db)

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
	
	// Empezar y escuchar al servidor usando nuestro enrutador configurado
	log.Fatal(http.ListenAndServe(":8080", router))
}
