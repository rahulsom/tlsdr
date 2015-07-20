[![Build Status](https://travis-ci.org/rahulsom/tlsdr.svg?branch=master)](https://travis-ci.org/rahulsom/tlsdr)

TLS;DR
===
> Transport Layer Security; Didn't Read

That's the general philosophy of most people having to support apps that need secure connections.

This project aims to fix that by making TLS more human readable.

Pre Install
---
If you need to install golang and wireshark, this is what you need to do. Instructions are for centos. You need to
find equivalent for your OS

```bash
# Install wireshark
yum install -y wireshark

# Install tools required for golang to work
yum install -y git gcc libpcap-devel

# Install golang (depending on your os+architecture)
wget --no-check-certificate https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz

# Ensure golang works correctly for the current user
echo 'export GOPATH=$HOME/golang' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc

# Ensure golang works in current shell
export GOPATH=$HOME/golang
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Installing
---
Assuming you've got golang correctly installed and configured,

```bash
go get github.com/rahulsom/tlsdr/tlsdr
```

Usage
---
To capture to a file, e.g., run
```bash
tshark -i eth0 -f "port 443" -w file.pcap 2>/dev/null 1>/dev/null
```
Then ^C to stop

To analyze a file and write text to STDOUT, run

```bash
tlsdr -i file.pcap
```

For more help, run `tlsdr` with no arguments.

For developing with this project, look at DEVELOPMENT.md
