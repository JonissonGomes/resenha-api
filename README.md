# Resenha API

API para gerenciamento de resenhas e jogadores.

## 🚀 Tecnologias

- Go 1.20
- Fiber (Framework Web)
- MongoDB
- Docker
- Docker Compose

## 📋 Pré-requisitos

- Docker
- Docker Compose
- Go 1.20 (opcional, se quiser rodar localmente)

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/JonissonGomes/resenha-api.git
cd resenha-api
```

2. Crie o arquivo .env:
```bash
cp .env.example .env
```

3. Configure as variáveis de ambiente no arquivo .env:
```env
MONGODB_URI=mongodb://mongodb:27017
PORT=8080
```

## 🏃‍♂️ Executando o Projeto

### Usando Docker Compose (Recomendado)

```bash
# Construir e iniciar os containers
make docker-compose
make docker-compose-up

# Ver logs
make docker-compose-logs

# Parar os containers
make docker-compose-down
```

### Comandos Disponíveis

```bash
# Ver todos os comandos disponíveis
make help

# Docker
make build          # Constrói a imagem Docker
make up            # Inicia a aplicação
make down          # Para a aplicação
make logs          # Mostra os logs
make clean         # Limpa containers e imagens

# Desenvolvimento
make test          # Executa os testes
make mongo-shell   # Acessa o shell do MongoDB

# Git
make git-status    # Mostra status do repositório
make git-add       # Adiciona alterações
make git-commit    # Faz commit (use: make git-commit m='sua mensagem')
make git-push      # Envia alterações
make git-pull      # Atualiza repositório
```

## 📡 Endpoints da API

### Health Check
- **GET** `/health`
  - Retorna o status da API
  - Resposta:
    ```json
    {
      "status": "ok",
      "message": "API is running"
    }
    ```

### Players
- **GET** `/api/players`
  - Retorna a lista de jogadores

## 📁 Estrutura do Projeto

```
resenha-api/
├── internal/
│   ├── database/     # Configuração do banco de dados
│   ├── handlers/     # Handlers da API
│   └── routes/       # Definição das rotas
├── main.go           # Ponto de entrada da aplicação
├── Dockerfile        # Configuração do Docker
├── docker-compose.yml # Configuração do Docker Compose
├── Makefile          # Comandos automatizados
└── README.md         # Documentação
```

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Faça commit das suas alterações (`make git-commit m='Add some AmazingFeature'`)
4. Faça push para a branch (`make git-push`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
