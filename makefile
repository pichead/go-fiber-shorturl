.PHONY: start deploy clean test
.IGNORE: deploy deploy-dev deploy-prod run-server

comment ?= update

define setup_env
	$(eval ENV_FILE := $(1))
	@echo " - setup env $(ENV_FILE)"
	$(eval include $(1))
	$(eval export sed 's/=.*//' $(1))
endef

push-code:
	git add .
	git commit -m "$(comment)"
	git push

run-server:
	go run ./cmd/main.go


deploy:
	$(call setup_env,.env)
	docker network create $(CONTAINER_NETWORK)
	docker-compose -p e-wallet-prod --env-file .env -f docker-compose.prod.yml up -d --build


deploy-prod:
	$(call setup_env,.env.prod)
	docker network create $(CONTAINER_NETWORK)
	docker-compose -p e-wallet-prod --env-file .env.prod -f docker-compose.prod.yml up -d --build
