FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.work go.work.sum ./
COPY httprunner/ httprunner/
COPY parsing/ parsing/
COPY server/ server/
COPY run/ run/
COPY data/ data/

WORKDIR /app/server
RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest
LABEL MAINTAINER="test-instructor"
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /app/server/server ./
COPY --from=builder /app/server/resource ./resource/
COPY --from=builder /app/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ["./server", "-c", "config.docker.yaml"]
