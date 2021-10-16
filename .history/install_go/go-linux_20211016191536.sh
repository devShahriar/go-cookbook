#!/bin/sh

wget -c https://golang.org/dl/go1.15.2.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz

echo "export GOROOT=/usr/local/go" >> ~/.profile
echo "export GOPATH=$HOME" >> ~/.profile
echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.profile
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=direct" >> ~/.profile
echo "export GOSUMDB=off" >> ~/.profile

echo "Installion successful"

echo $(go version)