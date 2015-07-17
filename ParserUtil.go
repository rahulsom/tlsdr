package main

import (
	_ "fmt"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"github.com/certifyTian/TLSHandshakeDecoder"
	_ "github.com/davecgh/go-spew/spew"
	"log"
	"container/list"
)


//TODO obviously not a main function, rename it to the caller
func parseFile(fileName string) list.List{
	connects := list.List{}
	if handle, err := pcap.OpenOffline(fileName); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		payloadPackets := producePayloadPackets(packetSource.Packets())
		handshakes := produceHandshakePackets(payloadPackets)
		//ProduceAlertPackets(payloadPackets)
		events := CreateEventsFromHSPackets(handshakes)
		for e := events.Front(); e != nil; e = e.Next() {
			log.Println("Events data:", e)
		}

	}
	return connects
}


// chanPacs: raw data as channel of gopacket.Packet from pcap file
// return: a list of packets that has payload([]byte)
func producePayloadPackets(chanPacs chan gopacket.Packet) list.List {
	var payloadPacs list.List
	for packet := range chanPacs {
		if (packet.ApplicationLayer() != nil) {
			payloadPacs.PushBack(packet.ApplicationLayer().Payload())
		}
	}
	for e := payloadPacs.Front(); e != nil; e = e.Next() {
		//log.Println("Payload data:", e)
	}
	return payloadPacs
}

// payloadPacs: a list of raw packets([]byte)
// return a list of TLSRecordLayer that only contains handshake packets
func produceHandshakePackets(payloadPacs list.List) list.List {
	var handShakePacs list.List
	for e := payloadPacs.Front(); e != nil; e = e.Next() {
		var p TLSHandshakeDecoder.TLSRecordLayer
		pl := e.Value.([]byte)
		err := TLSHandshakeDecoder.DecodeRecord(&p, pl); if err != nil {
			panic(err)
		} else {
			if (len(p.Fragment) > 4 && p.ContentType == TLSHandshakeDecoder.TypeHandshake) {
				//handShakePacs
				handShakePacs.PushBack(p)
			}
		}
	}

	//log.Printf("%04x", TLSHandshakeDecoder.VersionTLS10)
	return handShakePacs
}

func getHandShakeSegment(p TLSHandshakeDecoder.TLSRecordLayer) TLSHandshakeDecoder.TLSHandshake {
	var ph TLSHandshakeDecoder.TLSHandshake
	err := TLSHandshakeDecoder.TLSDecodeHandshake(&ph, p.Fragment); if err != nil {
		panic(err)
	} else {
		//log.Println("Parsed Handshake data:", ph)
		return ph
	}
}

//parse a handshake to a client hello struct
func parseClientHello(hsp TLSHandshakeDecoder.TLSHandshake) TLSHandshakeDecoder.TLSClientHello {
	var pch TLSHandshakeDecoder.TLSClientHello
	err := TLSHandshakeDecoder.TLSDecodeClientHello(&pch, hsp.Body); if err != nil {
		panic(err)
	} else {
		log.Println("Parsed Client Hello data: ", pch)
		return pch
	}
}


func CreateEventsFromHSPackets(handShakePacs list.List) list.List {
	var events list.List
	for e := handShakePacs.Front(); e != nil; e = e.Next() {
		log.Printf("Handshake data only:", e)
		handshake := getHandShakeSegment(e.Value.(TLSHandshakeDecoder.TLSRecordLayer))
		event := NewEvent(int(handshake.HandshakeType))
		events.PushBack(event)
	}
	return events
}











