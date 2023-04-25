## Build TrustedFission


#### Windows
 1. Run ```./trufaas/scripts/windows-build-script.bat ```
 2. Move the created ```fission.exe``` to ```/C:/Program Files (x86)/fission```

#### Ubuntu
 1. Run ```chmod +x ./trufaas/scripts/ubuntu-build-script.sh```
 2. Execute ```sudo ./trufaas/scripts/ubuntu-build-script.sh```

## Function Test

 1. Windows --> Run ```./trufaas/scripts/fn-test.bat {fnName}```

## Others

 1. Create Fn Route ```fission route create --name {fnName} --function {fnName} --url {fnName}```
 2. Start router ```kubectl port-forward svc/router 31314:80 -n fission```
 3. Test using fn url```curl http://localhost:31314/{fn_route}```
