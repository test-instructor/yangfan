FROM golang:1.23-alpine AS builder
WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on

RUN go build -o ./data/data-warehouse ./data


FROM yangfanz/python:3.10.0
ENV TZ=Asia/Shanghai
WORKDIR /yangfan
RUN apk add --no-cache tzdata bash
COPY --from=builder /go/src/yangfan/data/data-warehouse .
COPY --from=builder /go/src/yangfan/get-pip.py .
COPY --from=builder /go/src/yangfan/httprunner/pkg/boomer/data/x509 /root/.lc/x509
COPY --from=builder /go/src/yangfan/deploy/docker-compose/config/docker.config.yaml .

EXPOSE 9000
ENTRYPOINT ./data-warehouse -c docker.config.yaml
CMD ["tail", "-f", "/dev/null"]
