# Imagem base leve do Go
FROM golang:1.20-alpine

# Define o diretório da aplicação
WORKDIR /app

# Copiar o código fonte
COPY . .

# Instalar dependências
RUN go get github.com/gofiber/fiber/v2 && \
    go get github.com/joho/godotenv && \
    go get go.mongodb.org/mongo-driver/mongo && \
    go get go.mongodb.org/mongo-driver/mongo/options

# Build da aplicação
RUN go build -o main .

# Expor a porta
EXPOSE 8080

# Adiciona um script de inicialização para debug
COPY <<'EOF' /app/start.sh
#!/bin/sh
echo "Conteúdo do arquivo .env:"
cat /app/.env
echo "\nVariáveis de ambiente:"
env
echo "\nIniciando aplicação..."
./main
EOF

RUN chmod +x /app/start.sh

# Comando para executar a aplicação
CMD ["./main"]