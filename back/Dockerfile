# FROM node:12.14.0-alpine3.11
FROM node:18.13-alpine3.17

RUN apk add --no-cache bash

RUN npm install -g @nestjs/cli@7.5.6
# RUN npm install -g @nestjs/cli@9.2.0

USER node

WORKDIR /home/node/app