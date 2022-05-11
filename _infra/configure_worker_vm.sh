#!/bin/bash

VM_USER=alcides
VM_IP=34.95.160.141

# setup
ssh $VM_USER@$VM_IP -- sudo apt update
ssh $VM_USER@$VM_IP -- sudo apt upgrade -y

# install docker & docker-compose
ssh $VM_USER@$VM_IP -- sudo apt install --yes apt-transport-https ca-certificates curl gnupg2 software-properties-common && curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add - && sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" && sudo apt update && sudo apt install --yes docker-ce && sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

# setup my user
ssh $VM_USER@$VM_IP -- sudo groupadd docker
ssh $VM_USER@$VM_IP -- sudo usermod -aG docker $USER

ssh $VM_USER@$VM_IP -- git clone https://github.com/alcidesmig/tokens-overhead