#!/bin/bash
#
# This is the one time setup of the machine to get all necessary packages intalled
#

# Add Nginx as a repo
cat > /etc/yum.repos.d/nginx.repo << EOF
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/
gpgcheck=0
enabled=1
EOF

# Install all packages required
yum install -y wireshark nginx openssl git ruby rubygems man wget ntp #java-1.8.0-openjdk

# Install thor for carb to work
gem install thor

# Softlink files necessary for working
rm -rf /etc/nginx && cd /etc && ln -s /vagrant/test/etc/nginx
rm /etc/hosts     && cd /etc && ln -s /vagrant/test/etc/hosts
cd /etc/profile.d &&            ln -s /vagrant/test/etc/profile.d/aliases.sh

# Start nginx
service nginx start

# Install go
wget -q --no-check-certificate https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz
