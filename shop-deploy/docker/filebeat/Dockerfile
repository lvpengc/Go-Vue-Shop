ARG ELK_VERSION

# https://hub.docker.com/_/filebeat
FROM elastic/filebeat:${ELK_VERSION}

USER root
# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
