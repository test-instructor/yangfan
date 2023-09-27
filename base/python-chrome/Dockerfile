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


CMD ["bash"]
