# Mi Amorito 快速启动指南

Mi Amorito 是一个跨平台的角色聊天应用，支持与AI驱动的浪漫角色进行对话。

## 🚀 技术栈

- **后端**: Golang + Gin + MongoDB + Redis
- **AI**: OpenAI GPT-4.1-mini
- **跨平台共享**: Kotlin Multiplatform (KMP)
- **Android**: Kotlin + Jetpack Compose
- **Web**: React + TypeScript + Vite
- **数据库**: MongoDB (存储) + Redis (缓存)

## 📦 预设角色

应用内置5个迎合女性喜好的AI角色：

1. **陆庭深** - 霸道总裁：外冷内热的商界精英
2. **维克多** - 神秘王子：优雅的欧洲贵族
3. **顾云珩** - 落魄贵族：低调内敛的隐藏富豪
4. **江慕白** - 温柔医生：体贴可靠的暖男医生
5. **叶流风** - 浪漫艺术家：才华横溢的音乐家画家

## 🛠️ 快速开始

### 前置要求

- Docker & Docker Compose
- Go 1.21+
- Node.js 18+
- Android Studio (Android开发)
- OpenAI API Key

### 1. 克隆并初始化

```bash
# 克隆项目
git clone <your-repo>
cd mi-amorito

# 初始化开发环境
make setup
```

### 2. 配置环境变量

```bash
# 复制环境配置文件
cp backend/.env.example backend/.env

# 编辑配置文件，添加你的OpenAI API Key
nano backend/.env
```

### 3. 启动服务

```bash
# 终端1: 启动后端服务
make dev-backend

# 终端2: 启动Web前端
make dev-web

# 可选: 构建Android应用
make dev-android
```

### 4. 访问应用

- **Web应用**: http://localhost:3000
- **API文档**: http://localhost:8080/health
- **MongoDB**: localhost:27017
- **Redis**: localhost:6379

## 📱 项目结构

```
.
├── backend/              # Golang后端服务
│   ├── internal/
│   │   ├── config/      # 配置管理
│   │   ├── database/    # 数据库连接
│   │   ├── models/      # 数据模型
│   │   ├── services/    # 业务逻辑
│   │   ├── handlers/    # HTTP处理器
│   │   └── middleware/  # 中间件
│   └── main.go
├── shared/              # KMP共享代码
│   └── src/commonMain/kotlin/
├── androidApp/          # Android应用
│   └── src/main/java/
├── webApp/              # React Web应用
│   └── src/
├── docker-compose.yml   # 开发环境
├── Makefile            # 开发命令
└── README.md
```

## 🔧 开发命令

```bash
make help           # 显示所有可用命令
make setup          # 初始化开发环境
make dev-backend    # 启动后端服务
make dev-web        # 启动Web前端
make dev-android    # 构建Android应用
make build-all      # 构建所有应用
make clean          # 清理构建文件
make logs           # 查看环境日志
make test-api       # 测试API连接
```

## 🌟 核心功能

### 后端API

- `GET /api/v1/characters` - 获取角色列表
- `POST /api/v1/users` - 创建用户
- `POST /api/v1/chat/send` - 发送消息
- `GET /api/v1/chat/history/:session_id` - 获取聊天历史
- `GET /api/v1/chat/sessions` - 获取会话列表

### 前端特性

- 响应式设计，支持移动端和桌面端
- 实时聊天体验
- 角色选择界面
- 聊天历史管理
- 会话持久化

### KMP共享

- 共享数据模型
- 共享API客户端
- 跨平台业务逻辑

## 🔐 环境变量

```bash
# 服务器配置
PORT=8080
GIN_MODE=debug

# 数据库配置
MONGODB_URI=mongodb://admin:password123@localhost:27017/mi_amorito?authSource=admin
REDIS_ADDR=localhost:6379

# AI配置
OPENAI_API_KEY=your_openai_api_key_here
OPENAI_MODEL=gpt-4-1106-preview

# CORS配置
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

## 🐛 故障排除

### 常见问题

1. **MongoDB连接失败**
   ```bash
   make dev-env  # 确保Docker容器运行
   make logs     # 查看详细日志
   ```

2. **OpenAI API错误**
   - 检查API Key是否正确
   - 确认账户有足够余额
   - 验证模型名称是否正确

3. **端口冲突**
   - 修改`.env`文件中的端口配置
   - 或停止占用端口的其他服务

## 📝 API使用示例

### 发送消息

```bash
curl -X POST http://localhost:8080/api/v1/chat/send \
  -H "Content-Type: application/json" \
  -H "X-Device-ID: your-device-id" \
  -d '{
    "character_id": "domineering_ceo",
    "content": "你好，我是新来的秘书"
  }'
```

### 获取角色列表

```bash
curl -X GET http://localhost:8080/api/v1/characters
```

## 🚀 部署

### Docker部署

```bash
# 构建生产镜像
docker build -t mi-amorito-backend ./backend
docker build -t mi-amorito-web ./webApp

# 使用docker-compose部署
docker-compose -f docker-compose.prod.yml up -d
```

### 云平台部署

1. **后端**: 部署到任何支持Go的云平台
2. **Web**: 部署到Vercel、Netlify等静态网站服务
3. **数据库**: 使用MongoDB Atlas等云数据库服务

## 📞 技术支持

如果遇到问题，请检查：

1. 环境变量配置是否正确
2. 依赖是否安装完成
3. Docker容器是否正常运行
4. OpenAI API Key是否有效

祝您使用愉快！💕