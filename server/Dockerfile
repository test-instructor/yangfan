FROM golang:alpine AS builder

WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN  go build -o ./server ./server


FROM python:3.10.10-slim
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/server .
RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip \
    && pip3 install  funppy

EXPOSE 8888
ENTRYPOINT ./server -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]
