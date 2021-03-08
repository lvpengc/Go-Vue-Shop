FROM redis:latest

LABEL maintainer="lyupeng <lyupeng@189.cn>"

ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
## For security settings uncomment, make the dir, copy conf, and also start with the conf, to use it
#RUN mkdir -p /usr/local/etc/redis
#COPY redis.conf /usr/local/etc/redis/redis.conf

VOLUME /data

EXPOSE 6379

#CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
CMD ["redis-server"]
