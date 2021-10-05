echo "start vagrant provioning..."
sudo apt-get update

echo "configure lxd user and group"
sudo adduser vagrant lxd
newgrp lxd

echo "install lxd tools"
sudo apt install zfsutils-linux

echo "setup lxd manager"
lxd init --preseed

echo "install curl pkg..."
sudo apt-get install -y curl zfsutils-linux

echo "install golang pkg"
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update -y
sudo apt install -y golang-go 

echo "Install dlv pkg"
 git clone https://github.com/go-delve/delve.git $GOPATH/src/github.com/go-delve/delve
 cd $GOPATH/src/github.com/go-delve/delve
 make install

### export dlv bin path
export PATH=$PATH:/home/vagrant/go/bin

echo "Finished provisioning."
