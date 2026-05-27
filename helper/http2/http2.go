// Package http2 provides helpers for HTTP/2.
package http2

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

const listenerRetryBaseDelay = 5 * time.Millisecond

// Server is an HTTP server accepting both regular and proxied, both HTTP/1 and
// HTTP/2 connections.
//
// HTTP/2 is negotiated using TLS ALPN, either directly via a tls.Conn, either
// indirectly via the PROXY protocol. When the PROXY protocol is used, the
// TLS-terminating proxy in front of the server must be configured to accept
// the "h2" TLS ALPN protocol.
//
// The server is closed when the http.Server is.
type Server struct {
	h1         *http.Server  // regular HTTP/1 server
	h2         *http2.Server // HTTP/2 server
	h2Err      error         // HTTP/2 server setup error, if any
	h1Listener h1Listener    // pipe listener for the HTTP/1 server

	// The following fields are protected by the mutex
	mu        sync.Mutex
	closed    bool
	listeners map[net.Listener]struct{}
}

// NewServer creates a new HTTP server.
//
// A nil h2 is equivalent to a zero http2.Server.
func NewServer(h1 *http.Server, h2 *http2.Server) *Server { _ = "STUB: not implemented"; return nil }

// proxyListener.Accept never fails

func (srv *Server) errorLog() *log.Logger { _ = "STUB: not implemented"; return nil }

// Serve accepts incoming connections on the listener ln.
func (srv *Server) Serve(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

// net.Listener.Accept can fail for temporary failures, e.g. too many open
// files or other timeout conditions. In that case, wait and retry later.
// This mirrors what the net/http package does.

func (srv *Server) serveConn(baseCtx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids

// Mirror net/http.Server ConnContext behavior.

func (srv *Server) closeListeners() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) isClosed() bool { _ = "STUB: not implemented"; return false }

// h1Listener is used to signal back http.Server's Close and Shutdown to the
// HTTP/2 server.
type h1Listener struct {
	*pipeListener
	srv *Server
}

func (ln h1Listener) Close() error {
	_ = "STUB: not implemented"
	// pipeListener.Close never fails
	return nil
}
