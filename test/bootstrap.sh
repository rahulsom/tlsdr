#!/usr/bin/env bash

OUTDIR=/vagrant/data

if [ ! -d $OUTDIR ]; then
    mkdir -p $OUTDIR
fi

rm -rf $OUTDIR/*

function capture() {
    CAP_NAME=$1
    if [ "" = "${CAP_NAME}" ]; then
        echo "Can't call capture without CAP_NAME"
        exit 1
    fi
    tshark -i lo -f "port 443" -w ${OUTDIR}/${CAP_NAME}.pcap &
    LAST_CAPTURE=$!
    sleep 2
}

function stop() {
    kill ${LAST_CAPTURE}
}

function run() {
    echo "Testing $2 $4/$6 with $7"
    capture $2-$4-$6-$1
    curl --cacert /vagrant/test/$2/caroot/cacert.pem \
        -E /vagrant/test/$3/caroot/intrinsic/$4/cert.pem --key /vagrant/test/$5/caroot/intrinsic/$6/private.pem \
        $7 > /dev/null 2>/dev/null
    if [ $? = 0 ]; then
        echo "Connected"
    else
        echo "Failed"
    fi
    stop
}

#run  SUFFIX      TRUSTCA    CERT_CA    CERT_DIR       PRKEYCA    KEY_DIR        URL
 run  bad         goodca     goodca     goodclient     goodca     goodclient     https://bad.demo.com/
 run  trusted     goodca     goodca     goodclient     goodca     goodclient     https://trusted.demo.com/
 run  mutual      goodca     goodca     goodclient     goodca     goodclient     https://mutual.demo.com/
