package main

import (
	"log"
	"net"
	"os"
	"runtime"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func outboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func main() {
	switch runtime.GOOS {
	case "darwin":
	default:
		log.Fatal("test only in darwin environment")
	}

	c, err := icmp.ListenPacket("udp4", outboundIP())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP("8.8.8.8"), Zone: "en0"}); err != nil {
		log.Fatal(err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		m, _ := rm.Body.Marshal(1)
		log.Println(string(m))
		log.Printf("got reflection from %v", peer)
	default:
		log.Printf("got %+v; want echo reply", rm)
	}
}
