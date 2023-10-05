# Etapa de construção
FROM golang:1.17 as builder

# Definir a pasta de trabalho no container
WORKDIR /app

# Copiar go mod e sum files
COPY go.mod go.sum ./

# Baixar todas as dependências
RUN go mod download

# Copiar o código fonte do aplicativo para o container
COPY . .

# Construir o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Etapa de execução
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário pré-construído do estágio anterior
COPY --from=builder /app/main .

# Expor a porta na qual o app vai rodar
EXPOSE 8080

# Comando para executar o binário
CMD ["./main"]
