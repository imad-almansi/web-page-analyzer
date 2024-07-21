FROM golang:1

RUN apt-get update && apt-get -y upgrade && apt-get -y install build-essential ca-certificates

ENV GOOS=linux
ENV GOARCH=amd64

COPY . /web-page-analyzer

WORKDIR /web-page-analyzer

RUN go build -o bin/web-page-analyzer ./main.go

FROM ubuntu

RUN apt-get update && apt-get -y upgrade && apt-get -y install build-essential ca-certificates

WORKDIR /web-page-analyzer

EXPOSE 8080

COPY --from=0 /web-page-analyzer /web-page-analyzer

CMD /web-page-analyzer/bin/web-page-analyzer