package main

import (
	"container/list"
	_ "fmt"
	"github.com/certifyTian/TLSHandshakeDecoder"
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/google/gopacket"
	_ "github.com/google/gopacket/pcap"
	"log"
)

type Alert struct {
	ContentType uint8
	Version     uint16
	Length      uint16
	Level       uint8
	Description uint8
}

func ProduceAlertPackets(payloadPacs list.List) list.List {
	var alertPacs list.List
	for e := payloadPacs.Front(); e != nil; e = e.Next() {
		var p TLSHandshakeDecoder.TLSRecordLayer
		pl := e.Value.([]byte)
		err := TLSHandshakeDecoder.DecodeRecord(&p, pl)
		if err != nil {
			panic(err)
		} else {
			if len(p.Fragment) > 1 && p.ContentType == TLSHandshakeDecoder.TypeAlert {
				var alert Alert
				DecodeAlert(&alert, p)
				alertPacs.PushBack(alert)
			}
		}
	}
	for e := alertPacs.Front(); e != nil; e = e.Next() {
		log.Println("Alert data:", e)
	}
	return alertPacs
}

func DecodeAlert(a *Alert, data TLSHandshakeDecoder.TLSRecordLayer) {
	a.ContentType = data.ContentType
	a.Version = data.Version
	a.Length = data.Length
	a.Level = data.Fragment[0]
	if a.Level != 1 && a.Level != 2 && a.Level != 255 {
		a.Description = 255
		a.Level = 255
	} else {
		a.Description = data.Fragment[1]
	}
}
