package main

import (
	"fmt"
	"os"
	"time"
	// "strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type TCPID struct {
	clientPort layers.TCPPort
	serverPort layers.TCPPort
}

type TCPSegment struct {
	timestamp time.Time
	tcp       layers.TCP
}

type TCPStream struct {
	id       TCPID
	segments []TCPSegment
}

func NewTCPStream(tcpID TCPID) TCPStream {
	return TCPStream{tcpID, make([]TCPSegment, 100)}
}

func getDirectional(srcPort, destPort layers.TCPPort) (clientPort, serverPort layers.TCPPort) {
	if srcPort < destPort {
		clientPort = destPort
		serverPort = srcPort
	} else {
		clientPort = srcPort
		serverPort = destPort
	}
	return
}

func obtainTCPLayer(packet gopacket.Packet) (*layers.TCP, error) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		if tcp, ok := tcpLayer.(*layers.TCP); ok {
			return tcp, nil
		} else {
			return nil, fmt.Errorf("")
		}
	}
	return nil, fmt.Errorf("")
}

func main() {
	handle, err := pcap.OpenOffline(os.Args[1])
	defer handle.Close()
	if err != nil {
		fmt.Printf("failed open %s with %s", os.Args[1], err)
		os.Exit(1)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	var tcpStreams = make(map[TCPID]*TCPStream, len(packets))
	for packet := range packets {
		tcp, err := obtainTCPLayer(packet)
		if err != nil {
			fmt.Printf("failed read %s", packet)
			os.Exit(1)
		}
		timestamp := packet.Metadata().Timestamp.UTC()
		clientPort, serverPort := getDirectional(tcp.SrcPort, tcp.DstPort)
		tcpID := TCPID{clientPort, serverPort}
		if stream, ok := tcpStreams[tcpID]; ok {
			(*stream).segments = append((*stream).segments, TCPSegment{timestamp, *tcp})
		} else {
			newstream := NewTCPStream(tcpID)
			newstream.segments = append(newstream.segments, TCPSegment{timestamp, *tcp})
			tcpStreams[tcpID] = &newstream
		}
	}
	for id, _ := range tcpStreams {
		fmt.Println(id)
	}
}
