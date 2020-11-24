FROM ubuntu:18.04

ENV TZ 'Asia/Jakarta'

WORKDIR /app

RUN echo $TZ > /etc/timezone && \
    apt-get update && apt-get install -y tzdata ca-certificates && \
    rm /etc/localtime && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    apt-get clean

EXPOSE 8080

ADD altar-app /app/altar-app

RUN chmod +x /app/altar-app && \
    mkdir -p /app/log

CMD cd /app && ./altar-app