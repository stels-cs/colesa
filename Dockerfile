FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/colesa /opt/bin/colesa

CMD ["/opt/bin/colesa"]

