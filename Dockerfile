FROM golang:1.14-buster
LABEL maintainer="Benjamin Bollen <ben@ost.com>"

# based off of IPFS and textileio/go-threads Dockerfile,
# with thanks!

RUN apt-get update && apt-get install -y \
  ca-certificates

ENV SRC_DIR /go-mosaic

# Download packages for caching
COPY go.mod go.sum $SRC_DIR/
RUN cd $SRC_DIR \
  && go mod download

COPY . $SRC_DIR

RUN cd $SRC_DIR \
  && go install github.com/mosaicdao/go-mosaic/cmd/mosaic

# Get su-exec, a very minimal tool for dropping privileges,
# and tini, a very minimal init daemon for containers
ENV SUEXEC_VERSION v0.2
ENV TINI_VERSION v0.18.0
RUN set -x \
  && cd /tmp \
  && git clone https://github.com/ncopa/su-exec.git \
  && cd su-exec \
  && git checkout -q $SUEXEC_VERSION \
  && make \
  && cd /tmp \
  && wget -q -O tini https://github.com/krallin/tini/releases/download/$TINI_VERSION/tini \
  && chmod +x tini

# Now build the actual target image, which aims to be as small as possible.
FROM busybox:1.31.1-glibc
LABEL maintainer="Benjamin Bollen <ben@ost.com>"

ENV SRC_DIR /go-mosaic
COPY --from=0 /go/bin/mosaic /usr/local/bin/mosaic
COPY --from=0 /tmp/su-exec/su-exec /sbin/su-exec
COPY --from=0 /tmp/tini /sbin/tini
COPY --from=0 /etc/ssl/certs /etc/ssl/certs

# This shared lib (part of glibc) doesn't seem to be included with busybox.
COPY --from=0 /lib/x86_64-linux-gnu/libdl.so.2 /lib/libdl.so.2

# swarm TCP for IPFS/Libp2p
EXPOSE 4001

# hostAddr; should be exposed to the public
EXPOSE 4006

ENV MOSAIC_WORK_DIR /data/mosaic
RUN mkdir -p $MOSAIC_WORK_DIR \
  && adduser -D -h $MOSAIC_WORK_DIR -u 1000 -G users mosaic \
  && chown -R mosaic:users $MOSAIC_WORK_DIR

VOLUME $MOSAIC_WORK_DIR

ENTRYPOINT ["/sbin/tini", "--", "mosaic"]

CMD ["node"]
