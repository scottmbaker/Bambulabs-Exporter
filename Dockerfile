FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./
COPY .env .env

RUN go build -o /bambulabs-exporter

EXPOSE 9101

CMD [ "/bambulabs-exporter" ]
