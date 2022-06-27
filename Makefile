default:
	@echo "install:		Install go-related dependencies"
	@echo "image: 		Create docker image"
	@echo "compose-up:	Start service using docker-compose"
	@echo "compose-down:	Tear down service"
	@echo "mocks:		Regenerate mocks used in unit tests"
	@echo "run:		Creates image, then starts service"
	@echo "test:		Runs unit tests (and go vet)"
	@echo "vet:		Runs go vet"

install:
	@echo "==== Installing dependencies... ===="
	cat ./tools/tools.go | grep -oP '(?<=_ ).*' | xargs -tI % go install %

image:
	@echo "==== Building image ===="
	docker build -t powerflex -f ./Dockerfile ./

compose-up:
	@echo "==== Starting service... ===="
	docker-compose -f ./docker-compose.yml up

compose-down:
	@echo "==== Tearing down service... ===="
	docker-compose -f ./docker-compose.yml down

mocks:
	@echo "==== Regenerating mocks... ===="
	mockery --all --keeptree

run: image compose-up

test: vet
	@echo "==== Running unit tests... ===="
	go test ./... -cover

vet:
	@echo "==== Running go vet ===="
	go vet ./...
