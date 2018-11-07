package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	time.Sleep(time.Now().UTC().Add(2 * time.Second).Truncate(2 * time.Second).Add(1 * time.Second).Sub(time.Now()))
	t := <-ticker.C
	fmt.Printf("now:%v, delivered time:%v, \n", time.Now().UTC(), t.UTC())
	time.Sleep(4 * time.Second)
	t = <-ticker.C
	fmt.Printf("now:%v, delivered time:%v, \n", time.Now().UTC(), t.UTC())
	t = <-ticker.C
	fmt.Printf("now:%v, delivered time:%v, \n", time.Now().UTC(), t.UTC())
}
