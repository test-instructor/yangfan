FROM alpine:latest
RUN apk add --no-cache tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    date && \
    rm -rf /var/cache/apk/*
ENV TZ=Asia/Shanghai
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache python3 py3-pip vim
RUN python -m pip install --upgrade pip
RUN python -m venv ~/.hrp/venv
RUN ~/.hrp/venv/bin/python -m pip install --upgrade pip