FROM golang:1.23-alpine AS build

WORKDIR /app


COPY . .

RUN go mod tidy



RUN go build -v -o /app/goback 

FROM alpine:latest

WORKDIR /app

COPY --from=build /app /app

ENV PATH="/app:${PATH}"

EXPOSE 8888

ENTRYPOINT [ "goback" ]
