#
# Dockerfile for youtube-dl
#

FROM alpine:3.13.2

RUN set -xe \
    && apk add --no-cache ca-certificates \
        ffmpeg \
        openssl \
        python3 \
        py3-pip \
        aria2 \
    && pip3 install --no-cache-dir youtube-dl

WORKDIR /data

ENTRYPOINT ["youtube-dl"]
CMD ["--help"]
