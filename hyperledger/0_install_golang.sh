# Install Go (1.20.0)
sudo rm -r /usr/local/go
wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz 
tar -xvf go1.20.2.linux-amd64.tar.gz

# move extracted go directory into /usr/local
sudo mv go /usr/local

# Create and redirect to the new working directory
mkdir -p $HOME/go/src/github.com/$USER

echo export GOROOT=/usr/local/go >> ~/.bashrc
export GOROOT=/usr/local/go
echo export GOPATH=$HOME/go/src/github.com/$USER >> ~/.bashrc
export GOPATH=$HOME/go/src/github.com/$USER
echo export PATH=$GOPATH/bin:$GOROOT/bin:$PATH >> ~/.bashrc
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
source $HOME/.bashrc
# Verify set go paths
echo $GOROOT $GOPATH $PATH
# Verify installation version
go version
