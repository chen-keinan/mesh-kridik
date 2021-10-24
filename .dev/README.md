# Mesg-vagrantfile

vagrant file to be used for mesh associated  programs developments, file include :
- jq
- istio 

### Quick Start

```
 git clone git@github.com:chen-keinan/mesh-vagrantfile.git
 cd mesh-vagrantfile
 vagrant up

```

### Compile binary with debug params
```
GOOS=linux GOARCH=amd64 go build -v -gcflags='-N -l' demo.go
```
### Run debug on remote machine
```
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./demo
```

### Tear down
```
 vagrant destroy
