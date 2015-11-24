Use the Gitter plugin to send build status updates to [Gitter](https://gitter.im) when a build completes. The status updates are displayed in a room's activity feed.

Example configuration:

```yaml
notify:
  gitter:
    webhook: https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
```

Example configuration to notify multiple rooms:

```yaml
notify:
  gitter:
    webhook:
     - https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
     - https://webhooks.gitter.im/e/27a2e6ece5db91e06797
```
