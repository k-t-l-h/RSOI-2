FROM golang:1.14 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app/warehouse
COPY . .
RUN go build cmd/warehouse/main.go

FROM ubuntu:latest

RUN apt-get -y update && apt-get install -y tzdata

ENV TZ=Russina/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root

WORKDIR /usr/src/app/warehouse

COPY . .
COPY --from=builder /opt/app/warehouse/main .

EXPOSE 8280
CMD ./main