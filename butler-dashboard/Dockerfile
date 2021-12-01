#Stage 1
#######################################
# Use node image to build app
FROM node:16 AS build
# set working directory
WORKDIR /app

# install dependencies
COPY package.json .
RUN yarn install

# Build
COPY . .
RUN yarn build


#Stage 2
#######################################
# Use nginx to serve app
FROM nginx:1.21
# Set working directory to nginx resources directory
WORKDIR /usr/share/nginx/html

# Remove default nginx static resources
RUN rm -rf ./*
# Copies static resources from build stage
COPY --from=build /app/dist .
# Copies nginx config
COPY --from=build /app/config/nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]