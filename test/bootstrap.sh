#!/usr/bin/env bash
##
## Bootstrap.sh
##
## Prepares the test data for TLS;DR
##

# Configure and cleanup the data dir
OUTDIR=/vagrant/data
if [ ! -d $OUTDIR ]; then
    mkdir -p $OUTDIR
fi
rm -rf $OUTDIR/*

# Starts a capture using tshark
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

# Stops the last started capture
function stop() {
    kill ${LAST_CAPTURE}
}

# Runs a test
# Usage:
#      run SUFFIX TRUSTCA CERT_CA CERT_DIR PRKEYCA KEY_DIR URL
# if suffix is '-', new capture is not created
function run() {
    echo "Testing $2 $4/$6 with $7"
    CAPNAME=$2-$4-$6-$1

    if [ "$1" != "-" ]; then
        capture $CAPNAME
    fi

    curl --cacert /vagrant/test/$2/caroot/cacert.pem \
        -E /vagrant/test/$3/caroot/intrinsic/$4/cert.pem --key /vagrant/test/$5/caroot/intrinsic/$6/private.pem \
        $7 > $OUTDIR/$CAPNAME.out 2> $OUTDIR/$CAPNAME.err

    if [ $? = 0 ]; then
        echo "Connected"
    else
        echo "Failed"
    fi

    if [ "$1" != "-" ]; then
        stop
    fi
}

# Generate data for unit testing individual cases

#run  SUFFIX      TRUSTCA    CERT_CA    CERT_DIR       PRKEYCA    KEY_DIR        URL
 run  bad         goodca     goodca     goodclient     goodca     goodclient     https://bad.demo.com/
 run  trusted     goodca     goodca     goodclient     goodca     goodclient     https://trusted.demo.com/
 run  mutual      goodca     goodca     goodclient     goodca     goodclient     https://mutual.demo.com/

# Generate data for comprehensive reporting

capture integration
#run  SUFFIX      TRUSTCA    CERT_CA    CERT_DIR       PRKEYCA    KEY_DIR        URL
 run  -           goodca     goodca     goodclient     goodca     goodclient     https://bad.demo.com/
 run  -           goodca     goodca     goodclient     goodca     goodclient     https://trusted.demo.com/
 run  -           goodca     goodca     goodclient     goodca     goodclient     https://mutual.demo.com/
stop