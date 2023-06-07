FROM golang:alpine AS builder
WORKDIR /go/src/yangfan

COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./work ./work


FROM yangfanz/python:3.10.0
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/work .
COPY --from=builder /go/src/yangfan/docker.config.yaml .

EXPOSE 8888
ENTRYPOINT ./work -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]