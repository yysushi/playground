package main

import (
	"fmt"
	"os"
	// "strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	handle, err := pcap.OpenOffline(os.Args[1])
	defer handle.Close()
	if err != nil {
		fmt.Printf("failed open %s with %s", os.Args[1], err)
		os.Exit(1)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println(handle.LinkType())
	for packet := range packetSource.Packets() {
		handlePacket(packet) // Do something with a packet here.
	}
}

func handlePacket(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Printf("%b, %b, %b\n", tcp.SYN, tcp.ACK, tcp.FIN)
	}
}
