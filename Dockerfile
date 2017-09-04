FROM alpine:latest

MAINTAINER Eric Nograles <grales@gmail.com>

WORKDIR "/opt"

ADD .docker_build/go-api-example /opt/bin/go-api-example

CMD ["/opt/bin/go-getting-started"]

