FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache python3 py3-pip vim
RUN python -m venv ~/.lc/venv
RUN ~/.lc/venv/bin/python -m pip install --upgrade pip