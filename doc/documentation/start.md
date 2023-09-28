
## 环境准备
* 安装docker环境

## 执行脚本
```shell
# 进入目录
mkdir -p /home/yangfan && cd /home/yangfan
# 执行安装
bash -c "$(curl -fsSLk http://docs.yangfan.gd.cn/install/install.sh)"
```

## 配置修改
1. 脚本默认安装在/data/yangfan目录下
2. 如需修改配置请到安装目录下修改`./config/docker.config.yaml`文件
3. `./config/docker.config.yaml`中的`yang-fan.grafana-host`需要手动修改为`http://IP:3000`，否则无法访问grafana
4. `./config/docker.config.yaml`中的`yang-fan.front`需要手动修改为前端`http://IP:8080`，否则无法通过测试报告通知跳转到测试报告详情
5. `./config/docker.config.yaml`中的`fs`需要手动修改为飞书登录相关配置，否则无法登录
6. 修改配置后重新执行安装脚本即可