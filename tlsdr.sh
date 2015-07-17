#!/usr/bin/env bash

cd /root/tlsdr
gom build
ls -1tr data/*.pcap | sed -e "s/.pcap$//g" | sed -s "s/^data.//g" \
    | xargs -n 1 -I {} ./tlsdr -i /vagrant/data/{}.pcap -o /usr/share/nginx/html/{}/html -f html
ls -1tr data/*.pcap | sed -e "s/.pcap$//g" | sed -s "s/^data.//g" \
    | xargs -n 1 -I {} ./tlsdr -i /vagrant/data/{}.pcap -o /usr/share/nginx/html/{}/txt -f txt