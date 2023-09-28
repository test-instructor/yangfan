FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache python3 py3-pip
RUN python -m pip install --upgrade pip
RUN python -m venv ~/.hrp/venv
RUN ~/.hrp/venv/bin/python -m pip install --upgrade pip