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
func main() {
	if handle, err := pcap.OpenOffline("data/goodca-goodclient-goodclient-trusted.pcap"); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		payloadPackets := producePayloadPackets(packetSource.Packets())
		produceHandshakePackets(payloadPackets)

	}
}

//func handlePacket(p gopacket.PacketSource) {
//
//}

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
		log.Println("Payload data:", e)
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
				handShakePacs.PushBack(p)
			}
		}
	}
	for e := handShakePacs.Front(); e != nil; e = e.Next() {
		log.Println("Handshake data only:", e)
	}
	return handShakePacs
}


