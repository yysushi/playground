package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"

	"github.com/hanaugai/playground/go/golang/x/net/icmp/request"
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
				r := request.NewRequest(pid, i)
				// garbage because receiver may miss the request
				requests.Store(i, r)
				wm := icmp.Message{
					Type: ipv4.ICMPTypeEcho, Code: 0,
					Body: &icmp.Echo{
						ID: pid & 0xffff, Seq: i,
						Data: r.Encode(),
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
				var r *request.Request
				if rawRequest, found := requests.LoadAndDelete(m.Seq); !found {
					log.Printf("got reflection from %v, but ignore with unexpected seq %d\n", peer, m.Seq)
					continue
				} else {
					r = rawRequest.(*request.Request)
				}
				stat, err := r.CalcStat(m.Data, m.Seq, recvAt)
				if err != nil {
					log.Printf("got reflection from %v, but ignore with %s\n", peer, err)
					continue
				}
				log.Printf("got reflection from %v with %v\n", peer, stat.ElapsedMicroseconds)
			default:
				log.Printf("got %+v; want echo reply\n", rm)
			}
		}
	}()

	wg.Wait()
}
