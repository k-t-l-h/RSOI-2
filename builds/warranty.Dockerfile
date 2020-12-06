FROM golang:1.14 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app/warranty
COPY . .
RUN go build cmd/warranty/main.go

FROM ubuntu:latest

RUN apt-get -y update && apt-get install -y tzdata

ENV TZ=Russina/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root

WORKDIR /usr/src/app/warranty

COPY . .
COPY --from=builder /opt/app/warranty/main .

EXPOSE 8180

CMD ./main