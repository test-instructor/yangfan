# 声明镜像来源为golang:alpine
FROM golang:alpine AS builder

# 声明工作目录
WORKDIR /go/src/yangfan/master

COPY ./master .
COPY docker.config.yaml .

# 设置go env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN  go build -o master .
RUN echo "===================="
RUN ls -al
RUN echo "===================="

# Path: docs/docker/Dockerfile
FROM python:3.10.10-slim
# Copy the binary
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/master .
RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip \
    && pip3 install  funppy -i  https://pypi.tuna.tsinghua.edu.cn/simple

EXPOSE 7966
EXPOSE 9093
EXPOSE 9092
# 配置文件一定需要写对了,否则会报错
ENTRYPOINT ./master -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]












