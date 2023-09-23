FROM centos:8

ENV TZ=Asia/Shanghai

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN dnf install -y python39
RUN python3 -m venv ~/.hrp/venv

RUN python3 -m pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple \
&& pip3 install funppy -i https://pypi.tuna.tsinghua.edu.cn/simple

RUN dnf install -y \
    libXScrnSaver \
    alsa-lib \
    at-spi2-atk \
    gtk3 \
    nspr \
    nss \
    libX11-xcb \
    libxkbcommon \
    libxcb-dri3 \
    libXcomposite \
    libXdamage \
    libXfixes \
    libXrandr \
    mesa-libgbm \
    wget \
    liberation-fonts-common \
    liberation-sans-fonts \
    liberation-serif-fonts \
    curl \
    libu2f-host \
    vulkan \
    xdg-utils \
    vim \
    && dnf clean all

RUN wget -N --no-check-certificate https://raw.githubusercontent.com/notetoday/LocaleCN/master/LocaleCN.sh && bash LocaleCN.sh

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
RUN dnf localinstall -y google-chrome-stable_current_x86_64.rpm
RUN rm google-chrome-stable_current_x86_64.rpm


