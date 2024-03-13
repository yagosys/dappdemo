#!/bin/bash

# Update package information
sudo apt-get update

# Install necessary packages
sudo apt-get install -y ca-certificates curl gnupg

# Create directory for keyrings
sudo install -m 0755 -d /etc/apt/keyrings

# Download and add Docker GPG key, redirect "y" to the overwrite question
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg && echo "y" | sudo gpg --yes --import --no-default-keyring --keyring /etc/apt/keyrings/docker.gpg

# Add Docker repository to sources.list
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Update package information again
sudo apt-get update
# Install jq
sudo apt-get install -y jq
# Install Docker packages
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Create docker group if it doesn't exist
if ! getent group docker >/dev/null; then
    sudo groupadd docker
fi

# Add the user to the docker group
sudo usermod -aG docker $USER

# Activate changes in the current shell session without starting a new one
sudo su - $USER -c "docker info"

