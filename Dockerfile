# Estágio de Build
FROM golang:1.26-alpine AS builder

# Instala dependências necessárias para compilação (se houver CGO)
RUN apk add --no-cache git

WORKDIR /app

# Copia arquivos de dependências primeiro (otimiza cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
# Certifique-se que o caminho para o main.go está correto
RUN CGO_ENABLED=0 GOOS=linux go build -o main_app main.go

# Estágio de Execução
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/main_app .

# Exponha as portas que seu app usa
EXPOSE 8080 50051

CMD ["./main_app"]
