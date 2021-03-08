#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "DIR：${DIR}"

# 构建后台前端
echo "打包编译后台前端"
cd  $DIR/easy-mall-admin
npm run build:prod
cd $DIR/easy-mall-micro/admin
GOOS=linux go build -o $DIR/easy-mall-docker/admin/main *.go

# 启动 docker
echo "部署 docker 环境，启动项目..."
cd $DIR/easy-mall-docker
docker-compose down
docker-compose up -d --build admin

echo "项目部署完成"
