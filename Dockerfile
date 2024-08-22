FROM golang:alpine

WORKDIR /app

COPY . .

RUN go test ./...

RUN go build

ENTRYPOINT [ "./tenbounce", "start" ]
