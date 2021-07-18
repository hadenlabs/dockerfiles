#
# Dockerfile for dirsearch
#
FROM python:3.8.2-slim-buster
LABEL source="<github.com/maurosoria/dirsearch>"
LABEL maintainer="luis@hadenlabs.com"

ENV BUILD_DEPS \
    git

WORKDIR /root/dirsearch

RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
    $BUILD_DEPS \
    && git clone https://github.com/maurosoria/dirsearch.git ./ \
    && pip3 install --no-cache-dir -r requirements.txt \
    && pip3 install --no-cache-dir requests \
    && apt-get clean \
    && apt-get purge -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


ENTRYPOINT ["./dirsearch.py"]

CMD ["--help"]
