# drone-gitter

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-gitter/status.svg)](http://beta.drone.io/drone-plugins/drone-gitter)
[![](https://badge.imagelayers.io/plugins/drone-gitter:latest.svg)](https://imagelayers.io/?images=plugins/drone-gitter:latest 'Get your own badge on imagelayers.io')

Drone plugin for sending build status notifications via Gitter

## Usage

```sh
./drone-gitter <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
        "webhook": "https://webhooks.gitter.im/e/91e06797227ae5dbe6ec"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-gitter .
```

### Example

```sh
docker run -i plugins/drone-gitter <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
        "webhook": "https://webhooks.gitter.im/e/91e06797227ae5dbe6ec"
    }
}
EOF
```
