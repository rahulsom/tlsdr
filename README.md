TLS;DR
===
> Transport Layer Security; Didn't Read

That's the general philosophy of most people having to support apps that need secure connections.

This project aims to fix that by making TLS more human readable.

Running the Vagrant machine
---

1. Install Vagrant (and VirtualBox)
2. Run this to setup your box
```bash
vagrant up && vagrant ssh
```
3. Profit!

Creating PCAPs for testing
---
Once SSHed into the vagrant box, run this

```bash
sudo su - 
cd /vagrant/test
./bootstrap.sh
```

The Certificate Authorities
---
There are 2 certificate authorities created using carb

    /vagrant/test/goodca
    /vagrant/test/badca

Both have a password of `tlsdr`.

The Certificates
---
When we say certificate in this section, we actually mean certificate+key pair.

```
goodca
  - trusted    (CN:trusted.demo.com)
  - mutual     (CN:mutual.demo.com)
  - goodclient (CN:goodclient)
badca
  - bad        (CN:bad.demo.com)
  - badclient  (CN:badclient)
```

The development environment
---
This project uses [gom](https://github.com/mattn/gom) for dependency management.
 
1. If you haven't already, set your $GOPATH
2. Make sure $GOPATH/bin is in your $PATH
3. Go Get gom
```bash
go get github.com/mattn/gom
```
4. Say `gom` when you mean `go`. E.g.
```
gom run tlsdr.go
```
5. When you need new dependencies, modify `Gomfile`