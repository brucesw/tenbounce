FROM golang:alpine

WORKDIR /app

COPY . .

RUN mv tenbounce-prod.yaml tenbounce.yaml

RUN go build

ENTRYPOINT [ "./tenbounce", "start" ]
