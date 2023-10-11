package main

import (
	"fmt"

	message "github.com/johnmwood/resolver/dns"
	"github.com/johnmwood/resolver/resolver"
)

func main() {
	ip := "8.8.8.8:53"
	r := resolver.Resolver{}
	msg := message.NewMessage(
		message.WithDefaultHeader(),
		message.WithQuestion("dns.google.com", 1, 1),
	)

	conn, err := r.Dial("udp", ip)
	if err != nil {
		fmt.Printf("err dialing ip: %s\nfailed with err: %v", ip, err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(msg.ByteEncode())
	if err != nil {
		fmt.Printf("writing to connection failed with err: %v", err)
		return
	}

	buf := make([]byte, 2048)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("reading from connection failed with err: %v", err)
		return
	}

	fmt.Println(string(buf))
}
