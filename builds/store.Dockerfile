FROM golang:1.14 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app/store
COPY . .
RUN go build cmd/store/main.go

FROM ubuntu:latest

RUN apt-get -y update && apt-get install -y tzdata

ENV TZ=Russina/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER root

WORKDIR /usr/src/app/store

COPY . .
COPY --from=builder /opt/app/store/main .

EXPOSE 8480

CMD ./main