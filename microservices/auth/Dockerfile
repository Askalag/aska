
FROM golang:alpine3.15

RUN mkdir /auth_app
WORKDIR /auth_app
COPY ../../protolib/Makefile .

RUN go mod download
RUN go build -o /auth-app

EXPOSE 9091

CMD ["/auth-app"]