#!/usr/bin/env bash
if [ $(echo $PATH | grep /usr/local/go/bin -c) = 0 ]; then
    export GOPATH=/root/.go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
fi

export PS1="\[$(tput bold)\]\[$(tput setaf 6)\]\t \[$(tput setaf 2)\][\[$(tput setaf 3)\]\u\[$(tput setaf 1)\]@\[$(tput setaf 3)\]\h \[$(tput setaf 6)\]\W\[$(tput setaf 2)\]]\[$(tput setaf 4)\]\\$ \[$(tput sgr0)\]"

export CA_GOOD=/vagrant/test/goodca/caroot
export  CA_BAD=/vagrant/test/badca/caroot

alias tlsdr="go run /vagrant/src/tlsdr.go"
alias _curl="curl -v"
alias g_curl="_curl --cacert $CA_GOOD/cacert.pem"
alias b_curl="_curl --cacert $CA_BAD/cacert.pem"
alias ggcurl="g_curl -E $CA_GOOD/intrinsic/goodclient/cert.pem --key $CA_GOOD/intrinsic/goodclient/private.pem"
alias bbcurl="b_curl -E $CA_BAD/intrinsic/badclient/cert.pem --key $CA_BAD/intrinsic/badclient/private.pem"