package http2

import (
	"net"
	"sync"
)

// pipeListener is a hack to workaround the lack of http.Server.ServeConn.
// See: https://github.com/golang/go/issues/36673
type pipeListener struct {
	ch     chan net.Conn
	closed bool
	mu     sync.Mutex
}

func newPipeListener() *pipeListener { _ = "STUB: not implemented"; return nil }

func (ln *pipeListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (ln *pipeListener) Close() error { _ = "STUB: not implemented"; return nil }

// ServeConn enqueues a new connection. The connection will be returned in the
// next Accept call.
func (ln *pipeListener) ServeConn(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (ln *pipeListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

type pipeAddr struct{}

func (pipeAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (pipeAddr) String() string { _ = "STUB: not implemented"; return "" }
