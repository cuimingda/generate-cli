# generate-cli

## Install

```sh
go install ./cmd/gen
go install github.com/cuimingda/generate-cli/cmd/gen@latest
```

## Test

```sh
go test -cover ./internal/...
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```
