Use this plugin for sending build status notifications via Gitter. The status
updates are displayed in a room's activity feed.

## Config

The following parameters are used to configure the plugin:

* **webhook** - a single or a list of webhooks

The following secret values can be set to configure the plugin.

* **GITTER_WEBHOOK** - corresponds to **webhook**

It is highly recommended to put the **GITTER_WEBHOOK** into a secret so it is
not exposed to users. This can be done using the drone-cli.

```bash
drone secret add --image=plugins/gitter \
    octocat/hello-world GITTER_WEBHOOK https://webhooks.gitter.im/...
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  gitter:
    image: plugins/gitter
    webhook: https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
```

Example configuration that sends multiple messages:

```yaml
pipeline:
  gitter:
    image: plugins/gitter
    webhook:
      - https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
      - https://webhooks.gitter.im/e/27a2e6ece5db91e06797
```
