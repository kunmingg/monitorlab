ARG GOLANG_VERSION=1.13.1
FROM golang:$GOLANG_VERSION

ENV GO111MODULE=on

RUN mkdir -p /opt/client
WORKDIR /opt/client

COPY . .
RUN go build main.go

EXPOSE 8080

CMD ["/opt/client/main"]