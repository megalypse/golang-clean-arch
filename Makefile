run-server:
	docker-compose up -d
	go run cmd/server/main.go

build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/server ./cmd/server/main.go

IMAGE_ID := ${shell docker images 'cleanarch-server' -a -q}
clean:
	docker compose down
	docker rmi ${IMAGE_ID} -f

DB_IMAGE_ID := ${shell docker images 'cleanarch-db' -a -q}
clean-db:
	docker rmi ${DB_IMAGE_ID}


DOC_REGISTRY_PATH=./internal/presentation/phttp/controllers/controller_doc_registry.go
DOC_DEPS_PATH=./internal/domain/models
OTHER_PATH=./internal/main/factory/router_factory.go
generate-docs:
	./swag init -g ${OTHER_PATH} --pd ${DOC_DEPS_PATH}

run-compose: generate-docs build
	docker compose up -d

run-compose-clean: generate-docs clean build
	docker compose up -d


