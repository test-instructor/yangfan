FROM ubuntu:20.04

ENV TZ=Asia/Shanghai

RUN apt-get update -y && \
    apt-get install -y python3.10 python3.10-venv python3-pip


RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple \
&& pip3 install funppy -i https://pypi.tuna.tsinghua.edu.cn/simple


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
    xdg-utils \
    vim


# 设置编码格式
RUN wget -N --no-check-certificate https://raw.githubusercontent.com/notetoday/LocaleCN/master/LocaleCN.sh && bash LocaleCN.sh

# 下载和安装Chrome
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN dpkg -i google-chrome-stable_current_amd64.deb
RUN rm google-chrome-stable_current_amd64.deb


