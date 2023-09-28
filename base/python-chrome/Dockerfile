# 使用Alpine Linux作为基础映像
FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache python3 py3-pip
RUN python -m pip install --upgrade pip
RUN python -m venv ~/.hrp/venv
RUN ~/.hrp/venv/bin/python -m pip install --upgrade pip

# 安装必要的依赖
RUN apk update && apk add --no-cache curl && \
    apk add --no-cache xvfb dbus ttf-freefont fluxbox udev && \
    apk add --no-cache --virtual .build-deps bash xvfb-run && \
    rm -rf /var/cache/apk/*

# 下载并安装Google Chrome
RUN curl -O https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apk add --no-cache dpkg && \
    dpkg -i google-chrome-stable_current_amd64.deb || true && \
    apk --fix-broken --no-cache add && \
    rm google-chrome-stable_current_amd64.deb
# 运行一个Python示例脚本




#FROM centos:8
#
#ENV TZ=Asia/Shanghai
#
#RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
#RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
#
#RUN dnf install -y python39 vim
#RUN python3 -m pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple \
#    && pip3 install funppy -i https://pypi.tuna.tsinghua.edu.cn/simple
#
#RUN python3 -m venv ~/.hrp/venv
#
## 安装必要的依赖
#RUN dnf install -y epel-release && \
#    dnf install -y wget unzip which && \
#    dnf clean all
#
## 安装Chrome浏览器的依赖
#RUN dnf install -y fontconfig libX11 libXcomposite libXcursor libXdamage libXext libXi libXtst libXrandr libXScrnSaver libXss libXxf86vm libdrm mesa-libGL
#
## 下载并安装Chrome浏览器
#RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
#    dnf localinstall -y google-chrome-stable_current_x86_64.rpm && \
#    rm google-chrome-stable_current_x86_64.rpm
#
## 重新配置Chrome启动器以使用中文区域设置
#RUN sed -i 's|exec -a "$0" "$HERE/chrome" "$@"|exec -a "$0" "$HERE/chrome" "$@" --lang=zh-CN|g' /opt/google/chrome/google-chrome
#
#CMD ["bash"]
