FROM golang:alpine AS builder
WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./run ./run


FROM yangfanz/python:3.10.0.chrome

WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/run .
COPY --from=builder /go/src/yangfan/get-pip.py .
COPY --from=builder /go/src/yangfan/hrp/pkg/boomer/data/x509 /root/.hrp/x509
COPY --from=builder /go/src/yangfan/deploy/docker-compose/config/docker.config.yaml .

ENTRYPOINT ./run -c docker.config.yaml

ENTRYPOINT ./run -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]