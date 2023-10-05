# Etapa base para compilar o código Go
FROM golang:1.18.2-alpine3.16 as base

# Atualizar o index do pacote e definir o diretório de trabalho
RUN apk update
WORKDIR /src/api

# Copiar os arquivos go.mod e go.sum e instalar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante dos arquivos para o container
COPY . .

# Compilar o aplicativo
RUN go build -o api ./main.go

# Etapa para criar a imagem final
FROM alpine:3.16 as binary

# Copiar o binário compilado e outros arquivos necessários da etapa anterior
COPY --from=base /src/api/api .
# Se você tiver outros arquivos ou pastas que precisam ser copiados para a imagem, inclua-os aqui.
# Por exemplo, se você tiver uma pasta 'web', descomente a próxima linha.
# COPY --from=base /src/api/web ./web

# Expor a porta na qual o app vai rodar
EXPOSE 3000

# Comando para executar o binário
CMD ["./api"]
