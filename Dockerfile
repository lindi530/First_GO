# -------- build stage ----------
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

# 使用国内镜像，加速依赖安装
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache tzdata ca-certificates git

WORKDIR /app

# 复制依赖并下载
COPY go.mod go.sum ./
RUN go mod download

# 复制项目源码
COPY . .

# 编译 Go 程序
RUN go build -ldflags="-s -w" -o app main.go

# -------- run stage ----------
FROM alpine:3.20

# 使用国内源，安装证书和时区数据
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache tzdata ca-certificates \
    && rm -rf /var/cache/apk/*

RUN apk add --no-cache docker-cli

WORKDIR /app

# 拷贝可执行文件和配置文件
COPY --from=builder /app/app /app/app
COPY --from=builder /app/settings.yaml /app/settings.yaml

# 设置时区
ENV TZ=Asia/Shanghai

EXPOSE 8080

CMD ["./app"]
