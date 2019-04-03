FROM node:10-alpine

RUN npm install -g parcel-bundler

WORKDIR /frontend0
EXPOSE 1234

USER node


CMD npm init -y && parcel index.html