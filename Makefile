DOCKER := docker-compose -f deployments/development/docker-compose.yml

.PHONY: all
all: stop start

.PHONY: start
start:
	$(DOCKER) up -d

.PHONY: stop
stop:
	$(DOCKER) down

.PHONY: logs
logs:
	$(DOCKER) logs -f

.PHONY: test
test:
	$(DOCKER) run payment go test ./...