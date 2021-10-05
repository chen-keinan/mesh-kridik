# LXD-vagrantfile

vagrant file to be used for lxd associated  programs developments, file include :
- buntu/bionic64
- lxd cluster 
- dlv for remote debug

### Quick Start

```
 git clone git@github.com:chen-keinan/lxd-vagrantfile.git
 cd lxd-vagrantfile
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
