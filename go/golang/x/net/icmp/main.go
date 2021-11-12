package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
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

type request struct {
	pid    int
	seq    int
	sentAt time.Time
}

// EmbeddedMessage ...
type embeddedMessage struct {
	PID int64
}

func (r *request) embeddedMessage() *embeddedMessage {
	return &embeddedMessage{
		PID: int64(r.pid), // int is not propert to embed because a fixed size is necessary.
	}
}

func (r *request) encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, r.embeddedMessage())
	return buf.Bytes()
}

func decodeEmbeddedMessage(s []byte) (*embeddedMessage, error) {
	buf := bytes.NewReader(s)
	var em embeddedMessage
	err := binary.Read(buf, binary.LittleEndian, &em)
	if err != nil {
		return nil, err
	}
	return &em, nil
}

type stat struct {
	seq                 int
	elapsedMicroseconds time.Duration
}

func (r *request) calcStat(s []byte, seq int, recvAt time.Time) (*stat, error) {
	em, err := decodeEmbeddedMessage(s)
	if err != nil {
		return nil, fmt.Errorf("decode failure: %s", err)
	}
	if int(em.PID) != r.pid {
		return nil, fmt.Errorf("others (pid:%d)", em.PID)
	}
	return &stat{
		seq:                 seq,
		elapsedMicroseconds: recvAt.Sub(r.sentAt),
	}, nil
}

func newRequest(pid int, seq int) *request {
	return &request{
		pid:    pid,
		seq:    seq,
		sentAt: time.Now(),
	}
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

	var wg sync.WaitGroup
	var requests sync.Map
	pid := os.Getpid()

	// sender
	wg.Add(1)
	go func() {
		defer c.Close()
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			select {
			case <-time.After(time.Second):
				r := newRequest(pid, i)
				requests.Store(i, r)
				wm := icmp.Message{
					Type: ipv4.ICMPTypeEcho, Code: 0,
					Body: &icmp.Echo{
						ID: pid & 0xffff, Seq: i,
						Data: r.encode(),
					},
				}
				wb, err := wm.Marshal(nil)
				if err != nil {
					log.Printf("failed to prepare icmp message: %s", err)
					continue
				}
				if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP("8.8.8.8"), Zone: "en0"}); err != nil {
					log.Printf("failed to send icmp message: %s", err)
					continue
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
				log.Printf("failed to read: %s\n", err)
				continue
			}
			recvAt := time.Now()
			rm, err := icmp.ParseMessage(1, rb[:n])
			if err != nil {
				log.Printf("failed to parse icmp message: %s\n", err)
				continue
			}
			switch rm.Type {
			case ipv4.ICMPTypeEchoReply:
				m := (rm.Body).(*icmp.Echo)
				var r *request
				if rawRequest, found := requests.LoadAndDelete(m.Seq); !found {
					log.Printf("got reflection from %v, but ignore with unexpected seq %d\n", peer, m.Seq)
					continue
				} else {
					r = rawRequest.(*request)
				}
				stat, err := r.calcStat(m.Data, m.Seq, recvAt)
				if err != nil {
					log.Printf("got reflection from %v, but ignore with %s\n", peer, err)
					continue
				}
				log.Printf("got reflection from %v with %v\n", peer, stat.elapsedMicroseconds)
			default:
				log.Printf("got %+v; want echo reply\n", rm)
			}
		}
	}()

	wg.Wait()
}
