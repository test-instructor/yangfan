FROM golang:alpine AS builder
WORKDIR /go/src/yangfan

COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./run ./run


FROM yangfanz/python:3.10.0
ENV TZ=Asia/Shanghai
WORKDIR /yangfan
COPY --from=builder /go/src/yangfan/run .
COPY --from=builder /go/src/yangfan/get-pip.py .
COPY --from=builder /go/src/yangfan/hrp/pkg/boomer/data/x509 /root/.hrp/x509
COPY --from=builder /go/src/yangfan/deploy/docker-compose/config/docker.config.yaml .

ENTRYPOINT ./run -c docker.config.yaml

# 安装Chrome的依赖
RUN apt-get update && apt-get install -y \
    libxss1 \
    libasound2 \
    libatk-bridge2.0-0 \
    libgtk-3-0 \
    libnspr4 \
    libnss3 \
    libx11-xcb1 \
    libxkbcommon0 \
    libxcb-dri3-0 \
    libxcomposite1 \
    libxdamage1 \
    libxfixes3 \
    libxrandr2 \
    libgbm1 \
    libasound2 \
    wget \
    --fix-broken \
    fonts-liberation  \
    libcurl4  \
    libu2f-udev  \
    libvulkan1  \
    xdg-utils


# 设置编码格式
RUN wget -N --no-check-certificate https://raw.githubusercontent.com/notetoday/LocaleCN/master/LocaleCN.sh && bash LocaleCN.sh

# 下载和安装Chrome
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN dpkg -i google-chrome-stable_current_amd64.deb
RUN rm google-chrome-stable_current_amd64.deb

USER root
RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf

ENTRYPOINT ./run -c docker.config.yaml

CMD ["tail", "-f", "/dev/null"]