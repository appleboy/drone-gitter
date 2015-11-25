# Docker image for the Drone Gitter plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-gitter .

FROM alpine:3.2
RUN apk add -U ca-certificates && rm -rf /var/cache/apk/*
ADD drone-gitter /bin/
ENTRYPOINT ["/bin/drone-gitter"]