FROM golang:1.23-alpine AS builder

# For private modules and HTTPS
RUN apk add --no-cache git ca-certificates openssh-client

# Match existing project style
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0

# Buildx will set these automatically (linux/amd64, linux/arm64, ...)
ARG TARGETOS
ARG TARGETARCH

# Clone sibling repo lingce to satisfy local replace directives (../lingce/*).
# ARG LINGCE_REF allows pinning a specific branch/tag/commit; defaults to main.
ARG LINGCE_REF=main
RUN git clone --depth 1 --branch ${LINGCE_REF} https://github.com/lingcetech/lingce.git /workspace/lingce

# Copy dataWarehouse source into /workspace/dataWarehouse (mirrors local dev layout)
WORKDIR /workspace/dataWarehouse
COPY . .

RUN go mod download

RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags "-s -w" -o /out/data-warehouse .


FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /out/data-warehouse ./data-warehouse

EXPOSE 18080
ENTRYPOINT ./data-warehouse -c config.docker.yaml
