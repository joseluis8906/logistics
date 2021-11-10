.PHONY: build
build:
	@go build -o build/logistics cmd/logistics/*.go

.PHONY: clean
clean:
	@go clean && rm -rf build/* 

.PHONY: infra-up
infra-up:
	@cd scripts/ && docker-compose -p logistics up -d

.PHONY: infra-down
infra-down:
	@cd scripts/ && docker-compose -p logistics down --remove-orphans

.PHONY: configure
configure: infra-up
	@crypt set -endpoint="http://127.0.0.1:2379" -plaintext /config/logistics.json ./configs/runtime.json

.PHONY: reset
reset: clean infra-down
	@echo "cleaned and reseted :)"
