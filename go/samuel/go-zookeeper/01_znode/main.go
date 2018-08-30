package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect(
		[]string{"0.0.0.0:2181", "0.0.0.0:2182", "0.0.0.0:2183"},
		time.Second, zk.WithLogInfo(false))
	if err != nil {
		panic(err)
	}
	znode, err := c.Create("/koketani", []byte{}, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Created znode %s\n", znode)
	}
	eph, err := c.Create("/koketani", []byte{}, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Created ephemeral znode %s\n", eph)
	}
	_, _, ch, err := c.GetW(znode)
	if err != nil {
		panic(err)
	}
	for {
		e, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("New event: %+v\n", e)
	}
}
