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
