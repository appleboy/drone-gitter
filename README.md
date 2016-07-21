# drone-slack

Drone plugin for sending Gitter notifications

## Build

Build the binary with the following commands:

```
export GO15VENDOREXPERIMENT=1
go build
go test
```

## Docker

Build the docker image with the following commands:

```
export GO15VENDOREXPERIMENT=1
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo
```

Please note incorrectly building the image for the correct x64 linux and with GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-gitter' not found or does not exist..
```

## Usage

Post the build status to a room:

```
docker run --rm \
    -e GITTER_WEBHOOK=https://gitter.im/... \
    -e DRONE_REPO_OWNER=octocat \
    -e DRONE_REPO_NAME=hello-world \
    -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
    -e DRONE_COMMIT_REF=refs/heads/master \
    -e DRONE_COMMIT_BRANCH=master \
    -e DRONE_COMMIT_AUTHOR=octocat \
    -e DRONE_COMMIT_LINK=http://github.com/octocat/hello-world \
    -e DRONE_BUILD_NUMBER=1 \
    -e DRONE_BUILD_EVENT=push \
    -e DRONE_BUILD_STATUS=success \
    -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
    plugins/gitter
```
