.PHONY: mod,run,cover,clean,debug,test,release,stop,swag

CUR_PATH:=$(shell pwd)
APP_PATH:=$(CUR_PATH)
BINARY_NAME = dagger

export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOSUMDB=off
export GO111MODULE=on

default: debug

clean: stop
	@rm -rf $(APP_PATH)/bin/*
	@rm -rf $(APP_PATH)/config/config.toml
	@rm -rf ./logs/*
	@rm -rf ./nohup.out
	@echo "$(BINARY_NAME) clean up completed"

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
	@nohup $(APP_PATH)/bin/$(BINARY_NAME) -c $(CUR_PATH)/config/config.toml 2>&1 &

stop:
	@ps -ef | grep bin/$(BINARY_NAME) | grep -v grep | awk '{print $$2}' | xargs kill -9
	@echo "$(BINARY_NAME) service is shutdown"

swag:
	@swag init

cover:
	@go vet $(APP_PATH)
	@go test -coverpkg="./..." -cover $(APP_PATH)/... -gcflags='all=-N -l'
