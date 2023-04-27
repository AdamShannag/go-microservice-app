BROKER_BINARY=brokerApp
USER_BINARY=userApp
AUTH_BINARY=authApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_user build_auth
	@echo "Stopping docker images (if running...)"
	sudo docker-compose down
	@echo "Building (when required) and starting docker images..."
	sudo docker-compose up --build
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_user: builds the user binary as a linux executable
build_user:
	@echo "Building user binary..."
	cd ../user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../auth-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${AUTH_BINARY} ./cmd/api
	@echo "Done!"