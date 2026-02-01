FROM node:18-slim AS build

ARG NODE_MAX_OLD_SPACE=4096
ENV NODE_OPTIONS="--max-old-space-size=${NODE_MAX_OLD_SPACE} --openssl-legacy-provider"

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build || (echo "vuepress build failed; fallback to existing .vuepress/dist" && test -d .vuepress/dist)

FROM nginx:alpine
LABEL MAINTAINER="lingcetech"
COPY --from=build /app/.vuepress/dist /usr/share/nginx/html
