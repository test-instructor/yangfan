#!/bin/bash

# 在容器启动时设置 DNS 配置
echo "nameserver 8.8.8.8" > /etc/resolv.conf

