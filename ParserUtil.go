package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"github.com/certifyTian/TLSHandshakeDecoder"
	_ "github.com/davecgh/go-spew/spew"
	"log"
	"container/list"
	_ "errors"
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
		//var p TLSHandshakeDecoder.TLSRecordLayer
		pl := e.Value.([]byte)
		packets := DecomposeRecordLayer(pl)
		for e := packets.Front(); e != nil; e = e.Next() {
			if(e.Value.(TLSHandshakeDecoder.TLSRecordLayer).ContentType == TLSHandshakeDecoder.TypeHandshake) {
				handShakePacs.PushBack(e.Value)
			}
			//log.Println(e)
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
	for el := handShakePacs.Front(); el != nil; el = el.Next() {
		tlsRecordLayer := el.Value.(TLSHandshakeDecoder.TLSRecordLayer)
		hsPacets := DecomposeHandshakes(tlsRecordLayer.Fragment)
		for e := hsPacets.Front(); e != nil; e = e.Next() {
			handshake := e.Value.(TLSHandshakeDecoder.TLSHandshake)
			event := NewEvent(handshake.HandshakeType)
			events.PushBack(event)
			log.Printf("Created Event:", event)
		}

	}
	return events
}



func DecomposeRecordLayer(data []byte) list.List {
	if len(data) < 5 {
		return list.List{}
	}
	log.Println("Parsing one packet......")
	var tlsLayerlist list.List
	total := uint16(len(data))
	var offset uint16 = 0

	for (offset < total) {
		var p TLSHandshakeDecoder.TLSRecordLayer
		p.ContentType = uint8(data[0+offset])
		p.Version = uint16(data[1+offset])<<8 | uint16(data[2+offset])
		p.Length = uint16(data[3+offset])<<8 | uint16(data[4+offset])
		p.Fragment = make([]byte, p.Length)
		l := copy(p.Fragment, data[5+offset:5+p.Length+offset])
		tlsLayerlist.PushBack(p)
		log.Println("Length: ", p.Length)
		offset += 5+p.Length
		log.Print("Type:  ", p.ContentType)
		if l < int(p.Length) {
			fmt.Errorf("Payload to short: copied %d, expected %d.", l, p.Length)
		}
	}
	return tlsLayerlist
}

func DecomposeHandshakes(data []byte) list.List {
	if len(data) < 4 {
		return list.List{}
	}
	log.Println("Parsing one TLSLayer.......")
	var handshakelist list.List
	total := uint32(len(data))
	var offset uint32 = 0

	for (offset < total) {
		var p TLSHandshakeDecoder.TLSHandshake
		p.HandshakeType = uint8(data[0+offset])
		p.Length = uint32(data[1+offset])<<16 | uint32(data[2+offset])<<8 | uint32(data[3+offset])
		p.Body = make([]byte, p.Length)
		if (p.Length < 2048) {
			l := copy(p.Body, data[4+offset : 4+p.Length+offset])

			if l < int(p.Length) {
				fmt.Errorf("Payload to short: copied %d, expected %d.", l, p.Length)
			}
			offset += 4+p.Length
		} else {
			p.HandshakeType = 99
			p.Length = 0
			offset = total
		}

		log.Printf("Handshake Type: %d, length: %d ", p.HandshakeType, p.Length)
		handshakelist.PushBack(p)
	}
	return handshakelist
}










