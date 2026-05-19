package desktopApp.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import desktopApp.model.Employee;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class EmployeeClient {

    private final ObjectMapper mapper;
    private final HttpClient client;
    private final String pathServer;
    private final String headerName;
    private final String headerValue;
    private final String pathEmployees = "employees";

    public EmployeeClient(ObjectMapper mapper, HttpClient client, String pathServer, String headerName, String headerValue) {
        this.mapper = mapper;
        this.client = client;
        this.pathServer = pathServer;
        this.headerName = headerName;
        this.headerValue = headerValue;
    }

    public Employee[] get() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathEmployees))
                .header(headerName, headerValue)
                .GET()
                .build();

        return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                .thenApply(response -> {
                    try {
                        return mapper.readValue(response.body(), Employee[].class);
                    } catch (Exception e) {
                        e.printStackTrace();
                        return null; // Retorno por defecto en caso de error
                    }
                })
                .join(); // Espera a que la operación asíncrona termine y retorna el valor
    }

    public Employee getById(int empNo) {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathEmployees + "/" + empNo))
                .header(headerName, headerValue)
                .GET()
                .build();

        return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                .thenApply(response -> {
                    try {
                        return mapper.readValue(response.body(), Employee.class);
                    } catch (Exception e) {
                        System.out.println("No lo encontramos");
                        e.printStackTrace();                        
                        return null;
                    }
                })
                .join();
    }

    public Employee put(Employee e) {
        try {
            String jsonBody = mapper.writeValueAsString(e);

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(pathServer + pathEmployees + "/" + e.empNo))
                    .header(headerName, headerValue)
                    .PUT(HttpRequest.BodyPublishers.ofString(jsonBody))
                    .build();

            return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                    .thenApply(response -> {
                        try {
                            return mapper.readValue(response.body(), Employee.class);
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
    
    public Employee post(Employee e) {
        try {
            String jsonBody = mapper.writeValueAsString(e);

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(pathServer + pathEmployees))
                    .header(headerName, headerValue)
                    .POST(HttpRequest.BodyPublishers.ofString(jsonBody))
                    .build();

            return client.sendAsync(request, HttpResponse.BodyHandlers.ofString())
                    .thenApply(response -> {
                        try {
                            return mapper.readValue(response.body(), Employee.class);
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
    
    public boolean delete(int empNo) {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(pathServer + pathEmployees + "/" + empNo))
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
