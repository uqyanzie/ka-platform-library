FROM alpine:3.18

ENV NODE_VERSION 20.11.1

WORKDIR /app

COPY public/ /app/public
COPY src/ /app/src
COPY package.json /app/

RUN apk add --no-cache nodejs npm && \
    npm install

CMD ["npm", "start"]

