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
	return TCPStream{tcpID, make([]TCPSegment, 0)}
}

func (s *TCPStream) getTimeRange() (time.Time, time.Time) {
	return s.segments[0].timestamp, s.segments[len(s.segments)-1].timestamp
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
	start := time.Date(2999, time.December, 12, 31, 23, 59, 59, time.UTC)
	end := time.Date(999, time.December, 1, 1, 0, 0, 0, time.UTC)
	for packet := range packets {
		tcp, err := obtainTCPLayer(packet)
		if err != nil {
			fmt.Printf("failed read %s", packet)
			os.Exit(1)
		}
		timestamp := packet.Metadata().Timestamp.UTC()
		if timestamp.Before(start) {
			start = timestamp
		}
		if timestamp.After(end) {
			end = timestamp
		}
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
	normalizedStart := start.Round(time.Minute)
	for baseTime := normalizedStart; baseTime.Before(end); baseTime = baseTime.Add(time.Minute) {
		counts := 0
		for _, stream := range tcpStreams {
			begin, finish := stream.getTimeRange()
			if !(finish.Before(baseTime) || begin.After(baseTime.Add(time.Minute))) {
				counts = counts + 1
			}
		}
		fmt.Printf("%s,%d\n", baseTime, counts)
	}
}
