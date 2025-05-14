#!/bin/sh

# 确保日志目录存在
mkdir -p /app/logs/nginx

# 启动 Nginx
nginx

# 启动后端服务
./main 