# clicksend-go

ClickSend Go SDK

## Generate client

```sh
swagger generate client \
    --additional-initialism=SMS \
    --additional-initialism=MMS \
    --additional-initialism=FAX \
    --skip-models \
    --client-package=api \
    --name=clicksend \
    --principal=Auth \
    --spec=./spec/swagger.yaml

go get -u -f ./...
```
