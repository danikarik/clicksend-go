# clicksend-go

ClickSend Go SDK

## Getting dependencies

```sh
go get -u github.com/go-openapi/errors \
    github.com/go-openapi/loads \
    github.com/go-openapi/runtime \
    github.com/go-openapi/spec \
    github.com/go-openapi/strfmt \
    github.com/go-openapi/swag \
    github.com/go-openapi/validate \
    golang.org/x/net/context
```

## Generate client

```sh
swagger generate client \
    --additional-initialism=SMS \
    --additional-initialism=MMS \
    --additional-initialism=Fax \
    --skip-models \
    --client-package=api \
    --name=clicksend \
    --spec=./spec/swagger.yaml

go get -u -f ./...
```
