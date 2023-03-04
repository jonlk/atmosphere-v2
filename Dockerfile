FROM golang:1.20.1-alpine3.17 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /atmosphere-v2-api

FROM alpine:3.17 as publish

COPY --from=build /atmosphere-v2-api ./

CMD [ "/atmosphere-v2-api"]