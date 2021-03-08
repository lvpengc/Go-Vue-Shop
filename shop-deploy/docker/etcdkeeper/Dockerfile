FROM golang:latest as build

# 设置环境变量
# 可配置 https://goproxy.io、https://goproxy.cn
ARG GOPROXY=https://goproxy.cn
ENV GOPROXY ${GOPROXY}
ENV GO111MODULE on

WORKDIR /opt
RUN mkdir etcdkeeper
ADD . /opt/etcdkeeper
WORKDIR /opt/etcdkeeper/src/etcdkeeper

RUN go build -o etcdkeeper.bin main.go


FROM alpine:latest

# RUN apk add --no-cache ca-certificates

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /opt/etcdkeeper
COPY --from=build /opt/etcdkeeper/src/etcdkeeper/etcdkeeper.bin .
ADD assets assets

# 设置时区
RUN apk update && apk add -U --no-cache tzdata
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 8080

ENTRYPOINT ./etcdkeeper.bin -h 0.0.0.0 -p 8080
