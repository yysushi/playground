package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

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

type customMessage struct {
	pid    int
	sentAt time.Time
}

func newCustomMessage(pid int) *customMessage {
	return &customMessage{
		pid:    pid,
		sentAt: time.Now(),
	}
}

type EmbeddedMessage struct {
	PID    int64
	SentAt int64
}

func (c *customMessage) toEmbed() *EmbeddedMessage {
	return &EmbeddedMessage{
		PID:    int64(c.pid), // fixed size is necessary
		SentAt: c.sentAt.UnixMicro(),
	}
}

func newCustomMessageFromBinary(s []byte, expectPID int) (*customMessage, error) {
	buf := bytes.NewReader(s)
	var em EmbeddedMessage
	err := binary.Read(buf, binary.LittleEndian, &em)
	if err != nil {
		return nil, err
	}
	if int(em.PID) != expectPID {
		return nil, errors.New("other sender")
	}
	return &customMessage{
		pid:    int(em.PID),
		sentAt: time.UnixMicro(em.SentAt),
	}, nil
}

func (c *customMessage) toBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, c.toEmbed())
	return buf.Bytes()
}

func main() {
	switch runtime.GOOS {
	case "darwin":
	default:
		log.Fatal("test only in darwin environment")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := icmp.ListenPacket("udp4", outboundIP())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var wg sync.WaitGroup
	pid := os.Getpid()

	// sender
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Second):
				custom := newCustomMessage(pid)
				wm := icmp.Message{
					Type: ipv4.ICMPTypeEcho, Code: 0,
					Body: &icmp.Echo{
						ID: pid & 0xffff, Seq: 1,
						Data: custom.toBytes(),
					},
				}
				wb, err := wm.Marshal(nil)
				if err != nil {
					log.Fatal(err)
				}
				if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP("8.8.8.8"), Zone: "en0"}); err != nil {
					log.Fatal(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// receiver
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
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
				m := (rm.Body).(*icmp.Echo)
				if c, err := newCustomMessageFromBinary(m.Data, pid); err != nil {
					log.Printf("got reflection from %v, but ignore with %s", peer, err)
				} else {
					log.Printf("got reflection from %v with (id:%d, seq:%d, data:%#v)", peer, m.ID, m.Seq, c)
				}
			default:
				log.Printf("got %+v; want echo reply", rm)
			}
		}
	}()

	wg.Wait()
}
