
package desktopApp.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import desktopApp.model.Department;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class DepartmentClient {
    private final ObjectMapper mapper;
    private final HttpClient client;
    private final String pathServer;
    private final String headerName;
    private final String headerValue;
    private final String pathDepartments = "departments";

    public DepartmentClient(ObjectMapper mapper, HttpClient client, String pathServer, String headerName, String headerValue) {
        this.mapper = mapper;
        this.client = client;
        this.pathServer = pathServer;
        this.headerName = headerName;
        this.headerValue = headerValue;
    }

    public Department[] get() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathDepartments))
                .header(headerName, headerValue)
                .GET()
                .build();

        return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                .thenApply(response -> {
                    try {
                        return mapper.readValue(response.body(), Department[].class);
                    } catch (Exception e) {
                        e.printStackTrace();
                        return null; // Retorno por defecto en caso de error
                    }
                })
                .join(); // Espera a que la operación asíncrona termine y retorna el valor
    }

    public Department getById(String DeptNo) {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathDepartments + "/" + DeptNo))
                .header(headerName, headerValue)
                .GET()
                .build();

        return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                .thenApply(response -> {
                    try {
                        return mapper.readValue(response.body(), Department.class);
                    } catch (Exception e) {
                        e.printStackTrace();
                        return null;
                    }
                })
                .join();
    }

    public Department put(Department d) {
        try {
            String jsonBody = mapper.writeValueAsString(d);

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(pathServer + pathDepartments + "/" + d.deptNo))
                    .header(headerName, headerValue)
                    .PUT(HttpRequest.BodyPublishers.ofString(jsonBody))
                    .build();

            return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                    .thenApply(response -> {
                        try {
                            return mapper.readValue(response.body(), Department.class);
                        } catch (Exception ex) {
                            ex.printStackTrace();
                            return null;
                        }
                    })
                    .join();
            
        } catch (Exception ex) {
            ex.printStackTrace();
            return null;
        }

    }        
    
    public Department post(Department d) {
        try {
            String jsonBody = mapper.writeValueAsString(d);

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(pathServer + pathDepartments))
                    .header(headerName, headerValue)
                    .POST(HttpRequest.BodyPublishers.ofString(jsonBody))
                    .build();

            return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                    .thenApply(response -> {
                        try {
                            return mapper.readValue(response.body(), Department.class);
                        } catch (Exception ex) {
                            ex.printStackTrace();
                            return null;
                        }
                    })
                    .join();
            
        } catch (Exception ex) {
            ex.printStackTrace();
            return null;
        }
                
    }
    
    public boolean delete(String DeptNo) {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathDepartments + "/" + DeptNo))
                .header(headerName, headerValue)
                .DELETE()
                .build();

        return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                .thenApply(response -> {
                    try {
                        return response.statusCode() == 204; //204 significa: 'StatusNoContent', Se borró con éxito, no hay cuerpo que devolver --> se retorna true
                    } catch (Exception e) {
                        e.printStackTrace();
                        return false;
                    }
                })
                .join();                
    }   
    
}


