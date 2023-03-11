build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/server ./cmd/server/main.go

IMAGE_ID := ${shell docker images 'cleanarch-server' -a -q}
clean:
	docker compose stop server
	docker rmi ${IMAGE_ID} -f

DOC_DEPS_PATH=./internal/domain/models
SWAGGER_ENTRYPOINT=./internal/main/factory/router_factory.go
generate-docs:
	swag init -g ${SWAGGER_ENTRYPOINT} --pd --quiet

run-compose: generate-docs build
	docker compose up -d

run-compose-clean: generate-docs clean build
	docker compose up -d

run-compose-clean-all: generate-docs build
	docker compose down
	docker compose up -d


