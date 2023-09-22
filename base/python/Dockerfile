FROM python:3.10-slim

ENV TZ=Asia/Shanghai
RUN python3 -m venv ~/.hrp/venv
RUN python3 -m pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple \
&& pip3 install funppy -i https://pypi.tuna.tsinghua.edu.cn/simple


# 安装Chrome的依赖
RUN apt-get update && apt-get install -y \
    vim


