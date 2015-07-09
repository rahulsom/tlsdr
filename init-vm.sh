#!/bin/bash
#
# This is the one time setup of the machine to get all necessary packages intalled
#
cat > /etc/yum.repos.d/nginx.repo << EOF
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/
gpgcheck=0
enabled=1
EOF

yum install -y wireshark nginx java-1.8.0-openjdk openssl
service nginx start
rm -rf /etc/nginx && cd /etc && ln -s /vagrant/test/etc/nginx