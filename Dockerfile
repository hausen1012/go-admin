# 构建前端
FROM node:18-alpine as frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# 构建后端
FROM golang:1.23-alpine as backend-builder
WORKDIR /app/backend

# 安装编译工具
RUN apk add --no-cache gcc musl-dev

COPY backend/go.* ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# 最终镜像
FROM alpine:latest
WORKDIR /app

# 安装必要的运行时依赖
RUN apk add --no-cache sqlite nginx

# 复制前端构建产物到 Nginx 目录
COPY --from=frontend-builder /app/frontend/dist /usr/share/nginx/html

# 复制后端构建产物
COPY --from=backend-builder /app/backend/main /app/main
RUN chmod +x /app/main

# 配置 Nginx
COPY deploy/default.conf /etc/nginx/http.d/default.conf

# 创建启动脚本
COPY deploy/start.sh /app/start.sh
RUN chmod +x /app/start.sh

# 暴露端口
EXPOSE 80

# 启动命令
CMD ["/app/start.sh"] 