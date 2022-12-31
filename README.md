# Assessment project

This is the post-test project call assessment from go software engineer course.

## Installation
    go mod download

## Run with environment variable 
    DATABASE_URL=postgres://hgopfnjn:iaPYJl23WqduL57tfmotm0MmtyCAMfsm@tiny.db.elephantsql.com/hgopfnjn?sslmode=disable PORT=2565 go run ./cmd/server/server.go

## Run with configuration file
    touch ./internal/config/application.yaml
    
    echo 'port: 2565' >> application.yaml && echo 'database:' >> application.yaml && \
    echo '  url: postgres://hgopfnjn:iaPYJl23WqduL57tfmotm0MmtyCAMfsm@tiny.db.elephantsql.com/hgopfnjn?sslmode=disable' >> application.yaml

    mv application.yaml ./internal/config/application.yaml

    make start


## Run with docker
    make docker-build

    make docker-run DATABASE_URL="postgres://hgopfnjn:iaPYJl23WqduL57tfmotm0MmtyCAMfsm@tiny.db.elephantsql.com/hgopfnjn?sslmode=disable" PORT=2565

## Unit-test and Integration-test
    make unit-test
    
    make integration-test-up

## Down the integration-test sandbox
    make integration-test-down