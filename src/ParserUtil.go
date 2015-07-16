package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	_"github.com/certifyTian/TLSHandshakeDecoder"
	_ "github.com/davecgh/go-spew/spew"
	_"log"
	"container/list"
)

//TODO obviously not a main function, rename it to the caller
func main() {
	if handle, err := pcap.OpenOffline("data/goodca-goodclient-goodclient-trusted.pcap"); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		producePayloadPackets(packetSource.Packets())

	}
}

//func handlePacket(p gopacket.PacketSource) {
//
//}

func producePayloadPackets(chanPacs chan gopacket.Packet) list.List {
	var payloadPacs list.List
	for packet := range chanPacs {
		if (packet.ApplicationLayer() != nil) {
			payloadPacs.PushBack(packet.ApplicationLayer().Payload())
		}
	}
	for e := payloadPacs.Front(); e != nil; e=e.Next() {
		fmt.Println(e)
	}
	return payloadPacs
}

func produceHandshakePackets(payloadPacs list.List) list.List {
	var handShakePacs list.List
	for e := payloadPacs.Front(); e != nil; e = e.Next() {
		//pl := e.Value.([]byte)
		//continue implement
	}

	return handShakePacs
}
