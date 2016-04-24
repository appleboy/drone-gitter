# Docker image for Drone's slack notification plugin
#
#     docker build --rm=true -t plugins/drone-gitter .

FROM alpine:3.2
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD drone-gitter /bin/
ENTRYPOINT ["/bin/drone-gitter"]
