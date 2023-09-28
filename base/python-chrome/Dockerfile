FROM centos:8

ENV TZ=Asia/Shanghai

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN dnf install -y python39 vim
RUN python3 -m pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple \
    && pip3 install funppy -i https://pypi.tuna.tsinghua.edu.cn/simple

RUN python3 -m venv ~/.hrp/venv

RUN dnf install -y wget
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
RUN dnf localinstall -y google-chrome-stable_current_x86_64.rpm
RUN rm google-chrome-stable_current_x86_64.rpm

RUN wget -N --no-check-certificate https://raw.githubusercontent.com/notetoday/LocaleCN/master/LocaleCN.sh && bash LocaleCN.sh

# 安装中文字体和语言包
RUN dnf install -y wqy-zenhei-fonts langpacks-zh_CN

# 设置区域设置为中文
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.UTF-8
ENV LANG=zh_CN.UTF-8

# 重新配置Chrome启动器以使用中文区域设置
RUN sed -i 's|exec -a "$0" "$HERE/chrome" "$@"|exec -a "$0" "$HERE/chrome" "$@" --lang=zh-CN|g' /opt/google/chrome/google-chrome

CMD ["bash"]
