package desktopApp;

import com.fasterxml.jackson.databind.ObjectMapper;
import desktopApp.controller.DepartmentClient;
import desktopApp.controller.DeptManagerClient;
import desktopApp.controller.EmployeeClient;
import desktopApp.view.Home;
import java.net.http.HttpClient;

public class Main {

    public static void main(String[] args) {
        ObjectMapper mapper = new ObjectMapper();
        HttpClient client = HttpClient.newHttpClient();
        String pathServer = "http://localhost:8080/";
        String headerName = "Content-Type";
        String headerValue = "application/json";

        EmployeeClient ec = new EmployeeClient(mapper, client, pathServer, headerName, headerValue);
        DepartmentClient dc = new DepartmentClient(mapper, client, pathServer, headerName, headerValue);
        DeptManagerClient dmc = new DeptManagerClient(mapper, client, pathServer, headerName, headerValue);
        
        new Home(ec, dc, dmc);        
    }

}