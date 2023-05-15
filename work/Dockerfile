# 声明镜像来源为golang:alpine
FROM golang:alpine AS builder

WORKDIR /go/src/yangfan/work

COPY ./work .
COPY docker-compose.yml .

# 设置go env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN  go build -o work .

RUN echo "===================="
RUN ls -al
RUN echo "===================="

FROM python:3.10.10-slim
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/work .
RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip \
    && pip3 install  funppy -i  https://pypi.tuna.tsinghua.edu.cn/simple

EXPOSE 8888
# 配置文件一定需要写对了,否则会报错
ENTRYPOINT ./work -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]












