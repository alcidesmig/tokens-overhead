#!/bin/bash

VM_USER ?= alcidesmig
VM_IP ?= 34.125.56.135

# setup
sudo apt update
sudo apt upgrade -y

# install docker & docker-compose
sudo apt install --yes apt-transport-https ca-certificates curl gnupg2 software-properties-common && curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add - && sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" && sudo apt update && sudo apt install --yes docker-ce && sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

# setup my user
sudo groupadd docker
sudo usermod -aG docker $USER

ssh $VM_USER@$VM_IP -- git clone https://github.com/alcidesmig/tokens-overhead