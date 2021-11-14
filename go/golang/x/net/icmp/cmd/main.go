package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"

	pingtest "github.com/hanaugai/playground/go/golang/x/net/icmp"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
		log.Printf("there may be a bug since we don't test in %s environment\n", runtime.GOOS)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := icmp.ListenPacket("udp4", outboundIP())
	if err != nil {
		log.Fatal(err)
	}
	err = c.IPv4PacketConn().SetTTL(100)
	// err = c.IPv4PacketConn().SetTTL(1)
	if err != nil {
		log.Fatal("failed to set ttl:", err)
	}

	var wg sync.WaitGroup
	var requests sync.Map
	pid := rand.Intn(0xffff) // 16 bit for ICMP Echo message compatibility
	log.Println(pid)

	// sender
	wg.Add(1)
	go func() {
		defer c.Close()
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			select {
			case <-time.After(time.Second):
				// case <-time.After(time.Millisecond * 100):
				r := pingtest.NewRequest(pid, i, 1472)
				// r := pingtest.NewRequest(pid, i, 1473)
				// r := pingtest.NewRequest(pid, i, 100)
				wm := icmp.Message{
					Type: ipv4.ICMPTypeEcho, Code: 0,
					Body: &icmp.Echo{
						ID: pid, Seq: i, // no effect via datagram socket
						Data: r.Encode(),
					},
				}
				wb, err := wm.Marshal(nil)
				if err != nil {
					log.Printf("failed to prepare icmp message: %s", err)
					continue
				}
				// garbage because receiver may miss the request
				requests.Store(i, r.MarkSentAt())
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
			isEchoReply := rm.Type == ipv4.ICMPTypeEchoReply
			if !isEchoReply {
				switch rm.Type {
				case ipv4.ICMPTypeDestinationUnreachable:
					original := (rm.Body).(*icmp.DstUnreach)
					rm, err = icmp.ParseMessage(1, original.Data[20:])
					if err != nil {
						log.Printf("failed to parse original request in icmp error in %#v: %s\n", rm, err)
						continue
					}
				case ipv4.ICMPTypeTimeExceeded:
					original := (rm.Body).(*icmp.TimeExceeded)
					rm, err = icmp.ParseMessage(1, original.Data[20:])
					if err != nil {
						log.Printf("failed to parse original request in icmp error in %#v: %s\n", rm, err)
						continue
					}
				case ipv4.ICMPTypeParameterProblem:
					original := (rm.Body).(*icmp.ParamProb)
					rm, err = icmp.ParseMessage(1, original.Data[20:])
					if err != nil {
						log.Printf("failed to parse original request in icmp error in %#v: %s\n", rm, err)
						continue
					}
				default:
					original := (rm.Body).(*icmp.RawBody)
					rm, err = icmp.ParseMessage(1, original.Data[20:])
					if err != nil {
						log.Printf("failed to parse original request in icmp error in %#v: %s\n", rm, err)
						continue
					}
				}
			}

			m := (rm.Body).(*icmp.Echo)
			// if err != nil {
			// 	log.Printf("failed to handle icmp echo. it is broken: %s\n", err)
			// 	continue
			// }
			var r *pingtest.Request
			if rawRequest, found := requests.LoadAndDelete(m.Seq); !found {
				log.Printf("got reflection from %v, but ignore with unexpected seq %d\n", peer, m.Seq)
				continue
			} else {
				r = rawRequest.(*pingtest.Request)
			}
			stat, err := r.CalcStat(m, recvAt)
			if err != nil {
				log.Printf("got reflection from %v, but ignore with %s\n", peer, err)
				continue
			}
			if isEchoReply {
				log.Printf("got reflection from %v with %v\n", peer, stat.ElapsedMicroseconds)
			} else {
				log.Printf("got error on %d", stat.Seq)
			}
		}
	}()

	wg.Wait()
}
