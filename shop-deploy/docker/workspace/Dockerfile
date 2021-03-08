FROM golang:latest

LABEL maintainer="lyupeng <lyupeng@189.cn>"

# 添加 alias
COPY ./aliases.sh /root/aliases.sh

RUN sed -i 's/\r//' /root/aliases.sh && \
    echo "" >> ~/.bashrc && \
    echo "# Load Custom Aliases" >> ~/.bashrc && \
    echo "source ~/aliases.sh" >> ~/.bashrc && \
    echo "" >> ~/.bashrc

# 安装更新
RUN apt-get update -y && \
    apt-get upgrade -y
RUN apt-get install -y apt-utils autoconf automake libtool curl make g++ musl-dev vim net-tools zip git

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置环境变量
# 可配置 https://goproxy.io、https://goproxy.cn
ARG GOPROXY=https://goproxy.cn
ENV GOPROXY ${GOPROXY}
ENV GO111MODULE on

# 安装 micro、go-micro
RUN go get github.com/micro/micro && \
    go get github.com/micro/go-micro && \
    go get github.com/micro/protoc-gen-micro && \
    go get -u github.com/golang/protobuf/protoc-gen-go

# protobuf
WORKDIR /usr
RUN git config --global http.postBuffer 524288000 && git clone https://github.com/protocolbuffers/protobuf.git
# ADD https://github.com/protocolbuffers/protobuf/archive/master.zip ./protobuf.zip
# RUN apt-get install -y zip && unzip protobuf.zip

WORKDIR /usr/protobuf
RUN git submodule update --init --recursive && \
    # WORKDIR /usr/protobuf-master
    ./autogen.sh && ./configure && make && make check && make install

# 避免报错 protoc: error while loading shared libraries: libprotoc.so
ARG LD_LIBRARY_PATH=/usr/local/lib
ENV LD_LIBRARY_PATH ${LD_LIBRARY_PATH}

# 设置micro的registry环境变量，可etcd、mdns
ENV MICRO_REGISTRY etcd
ENV MICRO_REGISTRY_ADDRESS etcd:2379

WORKDIR /var/www

ENTRYPOINT [ "micro", "web" ]

EXPOSE 8082 80
