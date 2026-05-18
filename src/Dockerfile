# Estágio 1: Build
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Compila o binário principal (ajuste o path se o main.go estiver em outro lugar)
RUN CGO_ENABLED=0 go build -o main_app main.go

# Estágio 2: Execução
FROM alpine:latest
WORKDIR /root/
# Copia o binário do estágio de build
COPY --from=builder /app/main_app .
# Se o seu init.sql for necessário para o Go ler, copie também:
# COPY --from=builder /app/init.sql .

EXPOSE 8080 50051
CMD ["./main_app"]
