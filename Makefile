DOCKER_REGISTRY=docker.io/darkcrux
APP=webhook-manager
VERSION=latest
IMAGE=${DOCKER_REGISTRY}/${APP}:${VERSION}


all: clean lint build test package

clean:
	@echo '========================================'
	@echo ' Cleaning project'
	@echo '========================================'
	@go clean
	@rm -rf build | true
	@rm -rf volumes | true
	@docker-compose down
	@echo 'Done.'

deps:
	@echo '========================================'
	@echo ' Getting Dependencies'
	@echo '========================================'
	@echo 'Cleaning up dependency list...'
	@go mod tidy
	@echo 'Vendoring dependencies...'
	go mod vendor

gen:
	@echo '========================================'
	@echo ' Generating dependencies'
	@echo '========================================'
	@go generate ./cmd

build: deps gen
	@echo '========================================'
	@echo ' Building project'
	@echo '========================================'
	@go fmt ./...
	@go build -o build/bin/${APP} -mod vendor -ldflags "-X main.version=${VERSION} -w -s" .
	@echo 'Done.'

test:
	@echo '========================================'
	@echo ' Running tests'
	@echo '========================================'
	@go test ./...
	@echo 'Done.'

lint:
	@echo '========================================'
	@echo ' Running lint'
	@echo '========================================'
	@golint ./...
	@echo 'Done.'

run: build
	@echo '========================================'
	@echo ' Running application'
	@echo '========================================'
	@./build/bin/webhook-manager serve ${ARGS}
	@echo 'Done.'

docker-stop:
	@echo '========================================'
	@echo ' Stopping application'
	@echo '========================================'
	@@docker-compose stop
	@echo 'Done.'

package-image:
	@echo '========================================'
	@echo ' Packaging docker image'
	@echo '========================================'
	docker build -t ${IMAGE} .
	@echo 'Done.'

package: package-image
	
publish-image: package-image
	@echo '========================================'
	@echo ' Publishing image'
	@echo '========================================'
	docker push ${IMAGE}
	@echo 'Done.'

publish: publish-image
