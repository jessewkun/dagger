.PHONY: mod,run,cover,clean,debug,test,release

CUR_PATH:=$(shell pwd)
APP_PATH:=$(CUR_PATH)
CONFIG_NAME:=$(CUR_PATH)/config.toml
BINARY_NAME = dagger

export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOSUMDB=off
export GO111MODULE=on

default: debug

clean:
	@rm -rf $(APP_PATH)/bin/*
	@rm -rf $(APP_PATH)/config/config.toml
	@rm -rf ./logs/*
	@rm -rf ./nohup.out

mod:
	@go mod tidy -v
	@go mod download

debug: clean main.go go.sum go.mod
	@cp $(CUR_PATH)/config/debug.toml $(CUR_PATH)/config/config.toml
	@go env
	@go build -o $(APP_PATH)/bin/$(BINARY_NAME)
	@echo "[debug] $(BINARY_NAME) build success"

test: clean main.go go.sum go.mod
	@cp $(CUR_PATH)/config/test.toml $(CUR_PATH)/config/config.toml
	@go env
	@go build -o $(APP_PATH)/bin/$(BINARY_NAME)
	@echo "[test] $(BINARY_NAME) build success"

release: clean main.go go.sum go.mod
	@cp $(CUR_PATH)/config/release.toml $(CUR_PATH)/config/config.toml
	@go env
	@go build -o $(APP_PATH)/bin/$(BINARY_NAME)
	@echo "[release] $(BINARY_NAME) build success"

run:
	@nohup $(APP_PATH)/bin/$(BINARY_NAME) -c $(CUR_PATH)/config/config.toml

cover:
	@go env
	@go vet $(APP_PATH)
	@go test -coverpkg="./..." -c -cover -covermode=atomic $(APP_PATH) -o $(APP_PATH)/bin/$(BINARY_NAME)_cover -gcflags='all=-N -l'
