# 声明镜像来源为node:12.16.1
FROM node:16

# 声明工作目录
WORKDIR /go/src/cheetah/web
# 限制内存，否则报错
ENV NODE_OPTIONS=--max_old_space_size=4096

# 拷贝整个web项目到当前工作目录
COPY . .

# 通过npm下载cnpm, 用于加速npm下载,我这边是在GitHub actions中构建的,使用官方,国内使用则使用淘宝镜像
# RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

## 使用cnpm进行安装依赖
#RUN cnpm install
# 使用npm进行安装依赖
RUN ls -al
RUN  npm install --force

# 使用npm run build命令打包web项目
RUN npm run build
# ===================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为nginx:alpine, alpine 镜像小
FROM nginx:alpine

WORKDIR /cheetah/web/

# 从.docker-compose/nginx/conf.d/目录拷贝my.conf到容器内的/etc/nginx/conf.d/my.conf
COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf

# 从第一阶段进行拷贝文件
COPY --from=0 /cheetah/web/dist .

# 查看/etc/nginx/nginx.conf文件
RUN cat /etc/nginx/nginx.conf

# 查看 /etc/nginx/conf.d/my.conf
RUN cat /etc/nginx/conf.d/my.conf

# 查看 文件是否拷贝成功
RUN ls -al /usr/share/nginx/html
