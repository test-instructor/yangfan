FROM golang:1.23-alpine AS builder

WORKDIR /go/src/github.com/lingcetech/lingce
COPY go.work go.work.sum ./
COPY httprunner/ httprunner/
COPY server/ server/
COPY run/ run/

WORKDIR /go/src/github.com/test-instructor/yangfan/run
RUN CGO_ENABLED=0 go build -o run .

FROM alpine:latest

LABEL MAINTAINER="lingcetech"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /go/src/github.com/test-instructor/yangfan/run/run ./

EXPOSE 8889
ENTRYPOINT ./run -c config.docker.yaml
