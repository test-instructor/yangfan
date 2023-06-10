FROM golang:alpine AS builder
WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./run ./run


FROM yangfanz/python:3.10.0
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/run .
COPY --from=builder /go/src/yangfan/docker.config.yaml .

ENTRYPOINT ./run -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]

