# ==============================================================================
# Makefile helper functions for dagger
# ==============================================================================

# 变量定义
# ==============================================================================
SHELL := /bin/bash
BINARY_NAME := dagger

# 目录相关
ROOT_DIR := $(shell pwd)
BIN_DIR := $(ROOT_DIR)/bin
CONFIG_DIR := $(ROOT_DIR)/config
LOG_DIR := $(ROOT_DIR)/logs

# 配置文件
CONFIG_FILE := $(CONFIG_DIR)/config.toml
DEBUG_CONFIG := $(CONFIG_DIR)/debug.toml
RELEASE_CONFIG := $(CONFIG_DIR)/release.toml

# 颜色定义
SUCCESS := \033[32m
ERROR := \033[31m
WARNING := \033[33m
RESET := \033[0m

# Go 相关环境变量
export GO111MODULE := on
export GOPROXY := https://goproxy.cn
export GOSUMDB := sum.golang.org

# 帮助信息
# ==============================================================================
.DEFAULT_GOAL := help

.PHONY: help
help: ## 显示帮助信息
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# 开发相关命令
# ==============================================================================
.PHONY: mod
mod: ## 更新依赖
	@go mod tidy -v
	@go mod download

.PHONY: swag
swag: ## 生成 Swagger 文档
	@swag init

.PHONY: cover
cover: ## 运行测试覆盖率
	@go vet $(ROOT_DIR)
	@go test -coverpkg="./..." -cover $(ROOT_DIR)/... -gcflags='all=-N -l'

# 构建相关命令
# ==============================================================================
.PHONY: clean
clean: ## 清理构建产物
	@printf "==> Cleaning build artifacts...\n"
	@rm -rf $(BIN_DIR)/* $(CONFIG_FILE) $(LOG_DIR)/* ./nohup.out
	@printf "$(SUCCESS)Clean completed$(RESET)\n"

.PHONY: debug
debug: clean ## 构建调试版本
	@printf "==> Building debug version...\n"
	@cp $(DEBUG_CONFIG) $(CONFIG_FILE)
	@go build -o $(BIN_DIR)/$(BINARY_NAME)
	@printf "$(SUCCESS)[debug] Build completed$(RESET)\n"

.PHONY: release
release: clean ## 构建发布版本
	@printf "==> Building release version...\n"
	@cp $(RELEASE_CONFIG) $(CONFIG_FILE)
	@go build -o $(BIN_DIR)/$(BINARY_NAME)
	@printf "$(SUCCESS)[release] Build completed$(RESET)\n"

# 运行相关命令
# ==============================================================================
.PHONY: run
run: stop ## 运行服务
	@printf "==> Starting service...\n"
	@nohup $(BIN_DIR)/$(BINARY_NAME) -c $(CONFIG_FILE) > /dev/null 2>&1 &
	@$(MAKE) -s check-process

.PHONY: stop
stop: ## 停止服务
	@printf "==> Stopping service...\n"
	@ps -ef | grep bin/$(BINARY_NAME) | grep -v grep | awk '{print $$2}' | xargs -r kill -9
	@printf "$(WARNING)Service stopped$(RESET)\n"

.PHONY: check-process
check-process: ## 检查服务状态
	@if ps aux | grep -v grep | grep bin/$(BINARY_NAME); then \
		printf "$(SUCCESS)Service is running$(RESET)\n"; \
	else \
		printf "$(ERROR)Service is not running$(RESET)\n"; \
	fi
