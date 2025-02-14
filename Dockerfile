FROM golang:1.23-alpine AS build

WORKDIR /build

COPY ./app/ /build/app
COPY go.mod go.sum /build/

RUN  apk update && apk add build-base
RUN CGO_ENABLED=1 go build -o forumbin ./app/main.go

FROM alpine:latest AS deploy

WORKDIR /app

COPY --from=build /build/forumbin /app/forumbin
COPY --from=build /build/app/config/schema.sql /app/app/config/schema.sql
COPY ./static /app/static
EXPOSE 8080
CMD [ "/app/forumbin" ]