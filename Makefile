SHELL=/bin/bash
GO:=go

DEV_ENV=. dev-env.sh;
DC=$(DEV_ENV) docker-compose -p ginapi



app-build:
	$(GO) build -o ginapi cmd/ginapi/main.go



env-dev-up:
	$(DC) up -d

env-dev-down:
	$(DC) down

env-dev-setup: env-dev-down env-dev-up