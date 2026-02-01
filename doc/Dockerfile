FROM node:18-slim AS build

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:alpine
LABEL MAINTAINER="lingcetech"
COPY --from=build /app/.vuepress/dist /usr/share/nginx/html
