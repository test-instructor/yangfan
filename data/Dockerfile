FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.work go.work.sum ./
COPY httprunner/ httprunner/
COPY parsing/ parsing/
COPY server/ server/
COPY run/ run/
COPY data/ data/

WORKDIR /app/data
RUN CGO_ENABLED=0 go build -o data-warehouse .

FROM alpine:latest
LABEL MAINTAINER="test-instructor"
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /app/data/data-warehouse ./

EXPOSE 18080
ENTRYPOINT ["./data-warehouse", "-c", "config.docker.yaml"]
