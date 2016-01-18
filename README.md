# drone-gitter

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-gitter/status.svg)](http://beta.drone.io/drone-plugins/drone-gitter)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-gitter/coverage.svg)](https://aircover.co/drone-plugins/drone-gitter)
[![](https://badge.imagelayers.io/plugins/drone-gitter:latest.svg)](https://imagelayers.io/?images=plugins/drone-gitter:latest 'Get your own badge on imagelayers.io')

Drone plugin to send build status notifications via Gitter

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-gitter <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "webhook": "https://webhooks.gitter.im/e/91e06797227ae5dbe6ec"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-gitter <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "webhook": "https://webhooks.gitter.im/e/91e06797227ae5dbe6ec"
    },
    "vargs": {
    }
}
EOF
```
