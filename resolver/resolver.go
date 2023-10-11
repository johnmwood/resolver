package resolver

import "net"

type Resolver struct {
	net.Dialer
}
