package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	logger := log.New()
	entry := logger.WithField("zookeeper", "test")
	c, ch, err := zk.Connect(
		[]string{"0.0.0.0:2181", "0.0.0.0:2182", "0.0.0.0:2183"},
		time.Second, zk.WithLogger(entry), zk.WithLogInfo(true))
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	for {
		e := <-ch
		fmt.Printf("New event: %+v\n", e)
		switch e.Type {
		case zk.EventSession:
			switch e.State {
			// https://zookeeper.apache.org/doc/r3.3.6/zookeeperProgrammers.html#ch_zkSessions
			case zk.StateConnecting:
				fmt.Println("Connecting")
			case zk.StateConnected:
				fmt.Println("Connected")
			case zk.StateHasSession:
				// make sure still having session
				if c.State() == zk.StateHasSession {
					fmt.Printf("Established with %+v\n", c.SessionID())
				} else {
					fmt.Println("Established, but no more session")
				}
			case zk.StateDisconnected:
				fmt.Println("Disconnected")
			case zk.StateExpired:
				fmt.Println("Expired")
			}
		}
	}
}
