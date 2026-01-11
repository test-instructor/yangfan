FROM golang:1.23-alpine AS builder

WORKDIR /go/src/github.com/test-instructor/yangfan/server/v2
COPY . .

RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/github.com/test-instructor/yangfan/server/v2

COPY --from=0 /go/src/github.com/test-instructor/yangfan/server/v2/server ./
COPY --from=0 /go/src/github.com/test-instructor/yangfan/server/v2/resource ./resource/
COPY --from=0 /go/src/github.com/test-instructor/yangfan/server/v2/config.docker.yaml ./

# 挂载目录：如果使用了sqlite数据库，容器命令示例：docker run -d -v /宿主机路径/gva.db:/go/src/github.com/test-instructor/yangfan/server/v2/gva.db -p 8888:8888 --name gva-server-v1 gva-server:1.0
# VOLUME ["/go/src/github.com/test-instructor/yangfan/server/v2"]

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
