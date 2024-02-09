FROM golang:1.16-alpine

WORKDIR /app

ADD .  /app

RUN go build  -o url_shortner ./cmd/url_shortner

EXPOSE 8080

CMD [ "/url_shortner" ]
