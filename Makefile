# Variáveis
APP_NAME=resenha-api
PORT=8080

# Comandos
.PHONY: run build stop clean debug

# Constrói a imagem Docker
build:
	docker build -t $(APP_NAME) .

# Remove a imagem antiga e reconstrói
clean:
	docker rmi -f $(APP_NAME)
	docker build -t $(APP_NAME) .

# Executa a aplicação
run:
	docker run --rm -p $(PORT):$(PORT) --env-file .env $(APP_NAME)

# Executa com modo interativo para debug
debug:
	docker run --rm -it -p $(PORT):$(PORT) --env-file .env $(APP_NAME)

# Para todos os containers da aplicação
stop:
	docker stop $$(docker ps -q --filter ancestor=$(APP_NAME)) || true

# Constrói e executa a aplicação
up: clean run

# Comandos Docker
docker-compose:
	docker-compose build

docker-compose-up:
	docker-compose up

docker-compose-down:
	docker-compose down

docker-compose-logs:
	docker-compose logs -f

docker-compose-clean:
	docker-compose down -v
	docker system prune -f

# Comandos de desenvolvimento
test:
	go test ./...

# Comandos de banco de dados
mongo-shell:
	docker-compose exec mongodb mongosh

# Comandos de ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make build    - Constrói as imagens Docker"
	@echo "  make up       - Inicia os containers"
	@echo "  make down     - Para os containers"
	@echo "  make logs     - Mostra os logs dos containers"
	@echo "  make clean    - Remove containers, volumes e imagens não utilizadas"
	@echo "  make test     - Executa os testes"
	@echo "  make mongo-shell - Acessa o shell do MongoDB"