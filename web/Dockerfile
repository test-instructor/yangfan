FROM node:20-slim AS build

# Build-time toggles
ARG NODE_MAX_OLD_SPACE=4096
ARG VITE_LEGACY=false

ENV NODE_OPTIONS="--max-old-space-size=${NODE_MAX_OLD_SPACE}"
ENV VITE_LEGACY=${VITE_LEGACY}

WORKDIR /app
COPY package.json ./
RUN corepack enable yarn && yarn install
COPY . .
RUN yarn build

FROM nginx:alpine
LABEL MAINTAINER="lingcetech"
COPY --from=build /app/.docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=build /app/dist /usr/share/nginx/html
