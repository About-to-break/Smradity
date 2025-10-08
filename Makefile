DOCKER_COMPOSE = docker compose
COMPOSE_FILE = docker-compose.yml

.PHONY: run stop logs rebuild build config drain

run:
	$(DOCKER_COMPOSE) up -d

stop:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f

rebuild:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build --no-cache
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down
	$(DOCKER_COMPOSE)  -f $(COMPOSE_FILE) up -d

build:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

config:
	$(DOCKER_COMPOSE) config

drain:
	$(DOCKER_COMPOSE) down -v