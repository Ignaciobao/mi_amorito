.PHONY: help dev-backend dev-web dev-android install-deps clean build

help: ## 显示帮助信息
	@echo "Mi Amorito - 角色聊天应用开发命令"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-deps: ## 安装所有依赖
	@echo "Installing backend dependencies..."
	cd backend && go mod tidy
	@echo "Installing web dependencies..."
	cd webApp && npm install
	@echo "Dependencies installed!"

dev-env: ## 启动开发环境(MongoDB + Redis)
	@echo "Starting development environment..."
	docker-compose up -d
	@echo "Development environment started!"

dev-backend: ## 启动后端服务
	@echo "Starting backend server..."
	cd backend && go run main.go

dev-web: ## 启动Web前端
	@echo "Starting web frontend..."
	cd webApp && npm run dev

dev-android: ## 构建Android应用
	@echo "Building Android app..."
	./gradlew androidApp:assembleDebug

build-all: ## 构建所有应用
	@echo "Building all applications..."
	@echo "Building backend..."
	cd backend && go build -o bin/mi-amorito-server main.go
	@echo "Building web..."
	cd webApp && npm run build
	@echo "Building Android..."
	./gradlew androidApp:assembleRelease
	@echo "All applications built!"

clean: ## 清理构建文件
	@echo "Cleaning build files..."
	cd backend && rm -rf bin/
	cd webApp && rm -rf dist/
	./gradlew clean
	@echo "Clean completed!"

stop-env: ## 停止开发环境
	@echo "Stopping development environment..."
	docker-compose down
	@echo "Development environment stopped!"

logs: ## 查看开发环境日志
	docker-compose logs -f

test-api: ## 测试API连接
	@echo "Testing API connection..."
	curl -X GET http://localhost:8080/health || echo "Backend not running"

setup: install-deps dev-env ## 初始化开发环境
	@echo "Setup completed! You can now start developing."
	@echo "Run 'make dev-backend' in one terminal and 'make dev-web' in another."