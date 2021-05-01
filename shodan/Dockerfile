#
# Dockerfile for shodan
#

FROM alpine:3.13.2

RUN set -xe \
    && apk add --no-cache ca-certificates \
        openssl \
        python3 \
        py3-pip \
    && pip3 install --no-cache-dir shodan

WORKDIR /data

ENTRYPOINT ["shodan"]

CMD ["--help"]
