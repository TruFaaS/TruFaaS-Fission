## Build fission-cli

### For windows

```go build -o fission.exe cmd/fission-cli/main.go```


### Create fn Route
```fission route create --name {fnName} --function {fnName} --url {fnName}```

### Start router

```kubectl port-forward svc/router 31314:80 -n fission```
```curl http://localhost:31314/{fn_route}```
