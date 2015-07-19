TLS;DR
===
> Transport Layer Security; Didn't Read

That's the general philosophy of most people having to support apps that need secure connections.

This project aims to fix that by making TLS more human readable.

Installing
---
Assuming you've got golang correctly installed and configured,

```bash
go get github.com/rahulsom/tlsdr
```

Usage
---
To analyze a file and write text to STDOUT, run

```bash
tlsdr -i file.pcap
```

For more help, run `tlsdr` with no arguments.

For developing with this project, look at DEVELOPMENT.md
