package main

import (
	"fmt"
	"time"

	"github.com/miekg/dns"
)

func main() {
	m := new(dns.Msg)
	// m.SetQuestion(dns.Fqdn("miek.nl"), dns.TypeMX)
	// m.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
	m.SetQuestion(dns.Fqdn("google.com"), dns.TypeANY)
	// m.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
	// m.SetQuestion(dns.Fqdn("google.com"), dns.TypeMX)
	// m.SetQuestion(dns.Fqdn("netbox.ntti3.lan"), dns.TypeA)
	c := new(dns.Client)
	// // conn, err := dns.DialTimeout("udp", "127.0.0.1:53", 3*time.Second)
	// // conn, err := dns.DialTimeout("tcp", "127.0.0.1:53", 3*time.Second)
	// // conn, err := dns.DialTimeout("tcp", "127.0.0.1:53", 3*time.Second)
	// conn, err := dns.DialTimeout("tcp", "8.8.8.8:53", 3*time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// in, _, err := c.ExchangeWithConn(m, conn)
	// // in, _, err := c.Exchange(m, "172.18.17.17:53")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", in)
	// // fmt.Printf("%v %v\n", in, rtt)
	// // fmt.Printf("%v %v\n", in.Answer, rtt)
	// for _, e := range in.Answer {
	// 	// fmt.Printf("%#v\n", e)
	// 	fmt.Printf("%s %d %s %s\n", e.Header().Name, e.Header().Ttl, dns.ClassToString[e.Header().Class], dns.TypeToString[e.Header().Rrtype])
	// 	fmt.Printf("%d", dns.NumField(e))
	// 	for i := 1; i <= dns.NumField(e); i++ {
	// 		fmt.Printf("`%s`", dns.Field(e, i))
	// 	}
	// 	fmt.Printf("\n")
	// 	// fmt.Printf("%s\n", e.Header().Name)
	// 	// fmt.Printf("%s\n", e.Header().Rtype)
	// 	// fmt.Printf("%s\n", e.Header().Rtype)
	// }
	// // fmt.Printf("%v v\n", in, rtt)
	conn, err := dns.DialTimeout("tcp", "l.gtld-servers.net.:53", 3*time.Second)
	// conn, err := dns.DialTimeout("tcp", "8.8.8.8:53", 3*time.Second)
	if err != nil {
		panic(err)
	}
	in, _, err := c.ExchangeWithConn(m, conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", in)
	fmt.Println()
	// fmt.Printf("%#v\n", in.Answer[0])
	fmt.Printf("%#v\n", in.Ns[0])
	fmt.Printf("%#v\n", in.Extra[0])
}
