FROM alpine:latest

LABEL maintainer="lyupeng <lyupeng@189.cn>"

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY main /main

ENTRYPOINT [ "/main" ]

EXPOSE 80 443
