# Docker image for the Drone Gitter plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-gitter
#     make deps build
#     docker build --rm=true -t plugins/drone-gitter .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-gitter /bin/
ENTRYPOINT ["/bin/drone-gitter"]
