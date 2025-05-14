# Go Backend Project

## 项目结构

```
backend/
├── data/               # 数据文件目录
│   └── app.db         # SQLite 数据库文件
├── internal/           # 内部包
│   ├── config/        # 配置管理
│   ├── database/      # 数据库初始化和管理
│   ├── handlers/      # HTTP 请求处理器
│   ├── middleware/    # 中间件
│   ├── models/        # 数据模型
│   └── router/        # 路由配置
├── go.mod             # Go 模块文件
├── go.sum             # Go 依赖校验文件
└── main.go            # 应用程序入口点
```

## 开发指南

1. 配置文件：在项目根目录创建 `config.json` 来自定义配置，否则将使用默认配置
2. 数据库：使用 SQLite，数据文件位于 `data/app.db`
3. 默认管理员账户：
   - 用户名：admin
   - 密码：admin

## 启动服务

```bash
go run main.go
```

默认服务器地址为 `:8080` 