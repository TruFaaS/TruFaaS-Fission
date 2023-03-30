## Build fission-cli


#### Windows
```go build -o fission.exe cmd/fission-cli/main.go```

#### Ubuntu
``` GOOS=linux GOARCH=amd64 go build -o fission cmd/fission-cli/main.go```

```sudo mv ./fission /usr/local/bin/fission```


## Create fn Route
```fission route create --name {fnName} --function {fnName} --url {fnName}```

## Start router

```kubectl port-forward svc/router 31314:80 -n fission```
```curl http://localhost:31314/{fn_route}```
