#
# Dockerfile for pandoc
#

FROM pandoc/latex:2.16.2

ARG VERSION=2.16.2

LABEL maintainer="Team Hadenlabs <hello@hadenlabs.com>" \
      org.label-schema.vcs-url="https://github.com/hadenlabs/dockerfiles" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

ENV PATH $PATH:/root/.local/bin:/usr/local/bin

RUN tlmgr list \
    && tlmgr update --self \
    && tlmgr install \
    merriweather \
    fontaxes \
    mweights \
    mdframed \
    needspace \
    sourcesanspro \
    sourcecodepro \
    titling \
    ly1 \
    pagecolor \
    adjustbox \
    collectbox \
    titlesec \
    fvextra \
    pdftexcmds \
    footnotebackref \
    zref \
    fontawesome5 \
    footmisc \
    sectsty

# Reset the work dir
WORKDIR /data

ENTRYPOINT ["pandoc"]