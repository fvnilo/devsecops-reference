FROM node:23-alpine

WORKDIR /workspace
RUN mkdir libs apps

COPY package.json ./
COPY package-lock.json ./

COPY libs/echo libs/echo
COPY apps/service-c apps/service-c

RUN npm install

EXPOSE 3000

CMD ["node", "/workspace/apps/service-c/index.js"]

