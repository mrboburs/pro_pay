CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
IMG_NAME=${APP}
REGISTRY=${REGISTRY:-861701250313.dkr.ecr.us-east-1.amazonaws.com}
TAG=latest
ENV_TAG=latest
ifneq (,$(wildcard ./.env))
	include .env
endif
ifdef CI_COMMIT_BRANCH
        include .build_info
endif
run:
	swag init -g cmd/main.go  -o docs > /dev/null &&	go run  cmd/main.go
make create-env:
	cp ./.env.example ./.env
set-env:
	./scripts/set-env.sh ${CURRENT_DIR}
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go
schema-init:
	migrate create -ext sql -dir schema -seq init_schema - init go migration
clear:
	rm -rf ${CURRENT_DIR}/bin/*
lint:
	golint -set_exit_status ${PKG_LIST}
delete-branches:
	${CURRENT_DIR}/scripts/delete-branches.sh
swag-gen:
	echo ${REGISTRY}
	swag init -g command/main.go -o api/docs
migrate-down: set-env
	env POSTGRES_HOST=${POSTGRES_HOST} env POSTGRES_PORT=${POSTGRES_PORT} env POSTGRES_USER=${POSTGRES_USER} env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} env POSTGRES_DB=${POSTGRES_DB} ./scripts/migrate-jeyran.sh
minio:
	sudo MINIO_ROOT_USER=bakhodir MINIO_ROOT_PASSWORD=bakhodir0224 ./minio server /mnt/data --console-address ":7000"
daemon:
	cd ./command && compiledaemon -command="${CURRENT_DIR}/command"
ifneq (,$(wildcard vendor))
	go mod vendor
endif
.PHONY: vendor

.PHONY: sqlc
sqlc:
	sqlc generate
