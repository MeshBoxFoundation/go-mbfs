package tcpreuse

import (
	"errors"
	"sync"

	logging "mbfs/go-mbfs/gx/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"
)

var log = logging.Logger("reuseport-transport")

// ErrWrongProto is returned when dialing a protocol other than tcp.
var ErrWrongProto = errors.New("can only dial TCP over IPv4 or IPv6")

// Transport is a TCP reuse transport that reuses listener ports.
type Transport struct {
	v4 network
	v6 network
}

type network struct {
	mu        sync.RWMutex
	listeners map[*listener]struct{}
	dialer    dialer
}
