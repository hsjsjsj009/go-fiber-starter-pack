FROM golang:alpine as builder

WORKDIR /app

COPY . .
RUN go mod download && \
    go build -o app ./server/.

FROM alpine:latest as runner

ARG BUILD_ENV=staging

WORKDIR /opt/main/

ENV APP_ENV=$BUILD_ENV

RUN apk add --no-cache tzdata

COPY --from=builder /app/app .
COPY --from=builder /app/server/.env .env

ENTRYPOINT [ "/bin/sh","-c","/opt/main/app"]
