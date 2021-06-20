#
# Dockerfile for pandoc
#

from pandoc/latex:2.14.0.2

ENV PATH $PATH:/root/.local/bin:/usr/local/bin

# Reset the work dir
WORKDIR /data

ENTRYPOINT ["pandoc"]