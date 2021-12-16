# Base Image
FROM node:16 AS base
WORKDIR /app

# Install dependencies
COPY package.json .
RUN yarn install

# Dev image
FROM base AS dev
COPY . .

ENV PLATFORM_ENV dev
ENV APP_VERSION 0.0.1
ENV API_BASE_URL api.heybutler.local
ENV WEBAPP_BASE_URL app.heybutler.local

RUN yarn start

# Build image
FROM base AS build

ENV PLATFORM_ENV dev
ENV APP_VERSION 0.0.1
ENV API_BASE_URL api.heybutler.local
ENV WEBAPP_BASE_URL app.heybutler.local

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