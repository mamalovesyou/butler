# Base Image
FROM node:16-alpine AS base
WORKDIR /app

RUN apk --update --no-cache add automake git alpine-sdk nasm autoconf build-base \
		zlib \
		zlib-dev \
		libpng \
		libpng-dev\
		libwebp \
		libwebp-dev \
		libjpeg-turbo \
		libjpeg-turbo-dev

# Install dependencies
COPY package.json .
RUN yarn install

# Build image
FROM base AS build

ENV BUILD_TARGET prod
ENV APP_VERSION 0.0.1
ENV API_BASE_URL api.heybutler.local
ENV APP_BASE_URL app.heybutler.local

COPY . .
RUN yarn build


# Prod image
FROM nginx:1.21 as prod
# Set working directory to nginx resources directory
WORKDIR /usr/share/nginx/html
# Remove default nginx static resources
RUN rm -rf ./*
# Copies static resources from build stage
COPY --from=build /app/dist .
# Copies nginx config
COPY --from=build /app/config/nginx/nginx.conf /etc/nginx/conf.d/default.conf
# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]