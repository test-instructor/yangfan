FROM golang:alpine AS builder
WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./master ./master


FROM yangfanz/python:3.10.0
ENV TZ=Asia/Shanghai
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/master .
COPY --from=builder /go/src/yangfan/docker.config.yaml .

EXPOSE 7966
EXPOSE 9093
EXPOSE 9092

ENTRYPOINT ./master -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]
