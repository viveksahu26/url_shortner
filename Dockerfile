FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /url_shortner

EXPOSE 8080

CMD [ "/url_shortner" ]