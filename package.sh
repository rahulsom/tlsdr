#!/bin/bash
yum install -y gcc
go get github.com/mitchellh/gox
gox -build-toolchain
cd ~/tlsdr
cat Gomfile| cut -d " " -f 2 | xargs -n 1 go get
gox -output="bin/tlsdr_{{.OS}}_{{.Arch}}"
