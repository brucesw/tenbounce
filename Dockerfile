FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build

ENTRYPOINT [ "./tenbounce", "start" ]