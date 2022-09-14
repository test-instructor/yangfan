FROM node:16

WORKDIR /cheetah_web/
COPY . .

RUN yarn && yarn build

FROM nginx:alpine
LABEL MAINTAINER="185541429@qq.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /cheetah_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
