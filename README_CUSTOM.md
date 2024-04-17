# Developer Information

## API Shapes

The basepath for each API differs... documenting here for easy reference:

- harvester: /v1/harvester/<noun>/<namespace>/<name>
- harvester beta: /apis/harvesterhci.io/v1beta1/namespaces/<namespace string>/virtualmachineimages/<name>
- 

## Generating

```shell
$ openapi-generator generate -i api/openapi.json \
--generator-name go \
--package-name harvester \
--git-user-id drewmullen \
--git-repo-id harvester-go-sdk
```