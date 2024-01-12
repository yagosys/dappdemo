#!/bin/bash -xe
sudo apt update
sudo apt install curl git
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo corepack enable
node --version
yarn --version
