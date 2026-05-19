# Paso 1: Compilar la aplicación
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o mi-api .

# Paso 2: Crear una imagen ligera para ejecutarla
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mi-api .
EXPOSE 8080
CMD ["./mi-api"]