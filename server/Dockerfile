FROM golang:alpine AS builder

WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN  go build -o ./server ./server


FROM yangfanz/python:3.10.0
ENV TZ=Asia/Shanghai
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/server .
COPY --from=builder /go/src/yangfan/deploy/docker-compose/config/docker.config.yaml .

EXPOSE 8888
ENTRYPOINT ./server -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]
