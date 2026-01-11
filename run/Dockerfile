FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.work go.work.sum ./
COPY httprunner/ httprunner/
COPY parsing/ parsing/
COPY server/ server/
COPY run/ run/
COPY data/ data/

WORKDIR /app/run
RUN CGO_ENABLED=0 go build -o run .

FROM alpine:latest
LABEL MAINTAINER="test-instructor"
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /app/run/run ./

EXPOSE 8889
ENTRYPOINT ["./run", "-c", "config.docker.yaml"]
