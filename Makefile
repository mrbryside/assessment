
integration-test-up:
	DOCKER_BUILDKIT=0 docker-compose -f Docker-compose-test.yaml up --build --abort-on-container-exit --exit-code-from it_tests

integration-test-down:
	DOCKER_BUILDKIT=0 docker-compose -f Docker-compose-test.yaml down

unit-test:
	go clean -testcache && go test -v -tags=unit ./... -cover

unit-test-only:
	go clean -testcache && go test -v -tags=only ./... -cover

docker-build:
	docker build -t assessment/expenses:latest .

# usage => make docker-run DATABASE_URL="url" PORT=2565
docker-run:
ifneq ($(and $(DATABASE_URL),$(PORT)),)
	docker run -d -p $(PORT):$(PORT) -e DATABASE_URL=$(DATABASE_URL) -e PORT=$(PORT) --name assessment assessment/expenses:latest
else
		@echo 'no DATABASE_URL, PORT'
endif

build:
	go build -o "app" ./cmd/server

run: build unit-test
	./app

start:
	go run ./cmd/server/server.go
