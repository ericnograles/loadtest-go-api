FROM alpine:latest

MAINTAINER Eric Nograles <grales@gmail.com>

WORKDIR "/opt"

ADD .docker_build/loadtest-go-api /opt/bin/loadtest-go-api

CMD ["/opt/bin/loadtest-go-api"]

