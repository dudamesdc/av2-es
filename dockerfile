# Etapa 1: Construir o binário Go
FROM golang:1.23.4 AS builder

# Diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos Go para o contêiner
COPY . .

# Instalar dependências do Go
RUN go mod tidy

# Compilar o binário da aplicação
RUN go build -o main .

# Etapa 2: Construção da imagem final
FROM alpine:latest

# Instalar dependências necessárias (se houver alguma)
RUN apk --no-cache add ca-certificates

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/main /main

# Copiar o arquivo .env para o contêiner
COPY --from=builder /app/.env /app/.env

# Expor a porta que o aplicativo irá rodar
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["/main"]

