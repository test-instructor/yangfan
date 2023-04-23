# 声明镜像来源为golang:alpine
FROM golang:alpine AS builder

# 声明工作目录
WORKDIR /go/src/yangfan/server

# 拷贝整个server项目到工作目录
COPY . .

# 设置go env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# go generate 编译前自动执行代码
# go env 查看go的环境变量
# 打包项目生成文件名为server的二进制文件
RUN  go build -o server .
# RUN go generate && go env && go build -o server .
RUN echo "===================="
RUN ls -al
RUN echo "===================="

# Path: docs/docker/Dockerfile
FROM python:3.10.10-slim
# Copy the binary
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/server .
RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip \
    && pip3 install  funppy -i  https://pypi.tuna.tsinghua.edu.cn/simple

EXPOSE 8888
# 配置文件一定需要写对了,否则会报错
ENTRYPOINT ./server -c config.yaml

CMD ["tail", "-f", "/dev/null"]












