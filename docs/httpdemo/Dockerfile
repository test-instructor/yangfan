# 基础镜像选择 golang 官方提供的 alpine 版本
FROM golang:alpine

# 设置工作目录为 /app
WORKDIR /app

# 将代码复制到容器中
COPY . .

# 编译 Go 代码
RUN go build -o main .

# 暴露容器的端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]
