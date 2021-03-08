ARG ELK_VERSION

# https://hub.docker.com/_/logstash
FROM logstash:${ELK_VERSION}

USER root
# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
USER logstash

# Add your logstash plugins setup here
# Example: RUN logstash-plugin install logstash-filter-json