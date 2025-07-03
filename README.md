# Mi Amorito - 角色聊天应用

一个支持多平台的角色聊天应用，用户可以与预设的浪漫角色进行AI驱动的对话。

## 技术栈

- **后端**: Golang + Gin + MongoDB
- **跨平台共享**: Kotlin Multiplatform (KMP)
- **Android**: Kotlin + Jetpack Compose
- **Web**: React + TypeScript
- **AI**: OpenAI GPT-4.1-mini

## 角色设定

应用内置多个迎合女性喜好的角色：
- 霸道总裁：成功商界精英，外冷内热
- 神秘王子：来自异国的贵族，神秘优雅  
- 落魄贵族：隐藏身份的富有贵族，低调内敛
- 温柔医生：救死扶伤的暖男医生
- 艺术家：才华横溢的音乐家/画家

## 项目结构

```
.
├── backend/           # Golang 后端服务
├── shared/           # KMP 共享代码
├── androidApp/       # Android 应用
├── webApp/           # React Web 应用
├── docker-compose.yml # 开发环境
└── README.md
```

## 快速开始

### 后端服务
```bash
cd backend
go mod tidy
go run main.go
```

### Android 应用
```bash
./gradlew androidApp:installDebug
```

### Web 应用
```bash
cd webApp
npm install
npm start
```

## 功能特性

- ✅ 多角色选择
- ✅ AI 驱动对话
- ✅ 聊天记录持久化
- ✅ 跨平台支持
- ✅ 实时聊天体验