FROM golang:1.20.2-bullseye as base

WORKDIR $GOPATH/src/app/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /atmosphere-v2-api .

FROM scratch

COPY --from=base /atmosphere-v2-api ./

CMD [ "/atmosphere-v2-api"]
