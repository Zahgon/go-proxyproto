package proxyproto

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// readBufferSize is the size used for bufio.Reader's internal buffer.
//
// This is kept low to reduce per-connection memory overhead. If the header is
// larger than readBufferSize, the header will be decoded with multiple Read
// calls. For v1 the header length is at most 108 bytes. For v2 the header
// length is at most 52 bytes plus the length of the TLVs. We use 256 bytes to
// accommodate for the most common cases.
const readBufferSize = 256

var (
	// DefaultReadHeaderTimeout is how long header processing waits for header to
	// be read from the wire, if Listener.ReaderHeaderTimeout is not set.
	// It's kept as a global variable so to make it easier to find and override,
	// e.g. go build -ldflags -X "github.com/pires/go-proxyproto.DefaultReadHeaderTimeout=1s".
	DefaultReadHeaderTimeout = 10 * time.Second

	// ErrInvalidUpstream should be returned when an upstream connection address
	// is not trusted, and therefore is invalid.
	ErrInvalidUpstream = fmt.Errorf("proxyproto: upstream connection address not trusted for PROXY information")
)

// Listener is used to wrap an underlying listener,
// whose connections may be using the HAProxy Proxy Protocol.
// If the connection is using the protocol, the RemoteAddr() will return
// the correct client address. ReadHeaderTimeout will be applied to all
// connections in order to prevent blocking operations. If no ReadHeaderTimeout
// is set, a default of 10s will be used. This can be disabled by setting the
// timeout to < 0.
//
// Only one of Policy or ConnPolicy should be provided. If both are provided then
// a panic would occur during accept.
type Listener struct {
	// Listener is the underlying listener.
	Listener net.Listener
	// Deprecated: use ConnPolicyFunc instead. This will be removed in future release.
	Policy PolicyFunc
	// ConnPolicy is the policy function for accepted connections.
	ConnPolicy ConnPolicyFunc
	// ValidateHeader is the validator function for the proxy header.
	ValidateHeader Validator
	// ReadHeaderTimeout is the timeout for reading the proxy header.
	ReadHeaderTimeout time.Duration
	// ReadBufferSize is the read buffer size for accepted connections. When > 0,
	// each accepted connection uses this size for proxy header detection; 0 means default.
	ReadBufferSize int
}

// Conn is used to wrap and underlying connection which
// may be speaking the Proxy Protocol. If it is, the RemoteAddr() will
// return the address of the client instead of the proxy address. Each connection
// will have its own readHeaderTimeout and readDeadline set by the Accept() call.
type Conn struct {
	readDeadline atomic.Value // time.Time
	once         sync.Once
	readErr      error
	conn         net.Conn
	bufReader    *bufio.Reader
	// bufferSize is set when the client overrides via WithBufferSize; nil means use default.
	bufferSize        *int
	header            *Header
	ProxyHeaderPolicy Policy
	Validate          Validator
	readHeaderTimeout time.Duration
}

// Validator receives a header and decides whether it is a valid one
// In case the header is not deemed valid it should return an error.
type Validator func(*Header) error

// ValidateHeader adds given validator for proxy headers to a connection when passed as option to NewConn().
func ValidateHeader(v Validator) func(*Conn) { _ = "STUB: not implemented"; return nil }

// SetReadHeaderTimeout sets the readHeaderTimeout for a connection when passed as option to NewConn().
func SetReadHeaderTimeout(t time.Duration) func(*Conn) { _ = "STUB: not implemented"; return nil }

// WithBufferSize sets the size of the read buffer used for proxy header detection.
// Values <= 0 are ignored and the default (256 bytes) is used. Values < 16 are
// effectively 16 due to bufio's minimum. The default is tuned for typical proxy
// protocol header lengths.
func WithBufferSize(length int) func(*Conn) { _ = "STUB: not implemented"; return nil }

// Accept waits for and returns the next valid connection to the listener.
func (p *Listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"

	// Get the underlying connection.
	return *new(net.Conn), nil
}

// can't decide the policy, we can't accept the connection.

// keep listening for other connections.

// Handle a connection as a regular one.

// If the ReadHeaderTimeout for the listener is unset, use the default timeout.

// Set the readHeaderTimeout of the new conn to the value of the listener

// Close closes the underlying listener.
func (p *Listener) Close() error { _ = "STUB: not implemented"; return nil }

// Addr returns the underlying listener's network address.
func (p *Listener) Addr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// NewConn is used to wrap a net.Conn that may be speaking the PROXY protocol
	// into a proxyproto.Conn.
	//
	// NOTE: NewConn may interfere with previously set ReadDeadline on the provided net.Conn,
	// because it sets a temporary deadline when detecting and reading the PROXY protocol header.
	// If you need to enforce a specific ReadDeadline on the connection, be sure to call Conn.SetReadDeadline
	// again after NewConn returns, to restore your desired deadline.
	new(net.Addr)
}

func NewConn(conn net.Conn, opts ...func(*Conn)) *Conn { _ = "STUB: not implemented"; return nil }

// Read is check for the proxy protocol header when doing
// the initial scan. If there is an error parsing the header,
// it is returned and the socket is closed.
func (p *Conn) Read(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// Ensure header processing runs at most once and surface any errors.
	return 0, nil
}

// Drain the buffer if it exists and has data.

// Did we empty the buffer?
// Buffering a net.Conn means the buffer doesn't return io.EOF until the connection returns io.EOF.
// Therefore, we use Buffered() == 0 to detect if we are done with the buffer.

// Garbage collect the buffer.

// Return immediately. Do not touch p.conn.
// If err is EOF here, it means the connection is actually closed,
// so we should return that error to the user anyway.

// If buffer was empty to begin with (shouldn't happen with the >0 check
// but good for safety), clear it.

// From now on, read directly from the underlying connection.

// Write wraps original conn.Write.
func (p *Conn) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// Ensure header processing has completed before writing.
	return 0, nil
}

// Close wraps original conn.Close.
func (p *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// ProxyHeader returns the proxy protocol header, if any. If an error occurs
// while reading the proxy header, nil is returned.
func (p *Conn) ProxyHeader() *Header {
	_ = "STUB: not implemented"
	// Ensure header processing runs at most once.
	return nil
}

// LocalAddr returns the address of the server if the proxy
// protocol is being used, otherwise just returns the address of
// the socket server. In case an error happens on reading the
// proxy header the original LocalAddr is returned, not the one
// from the proxy header even if the proxy header itself is
// syntactically correct.
func (p *Conn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	// Ensure header processing runs at most once.
	return *new(net.Addr)
}

// RemoteAddr returns the address of the client if the proxy
// protocol is being used, otherwise just returns the address of
// the socket peer. In case an error happens on reading the
// proxy header the original RemoteAddr is returned, not the one
// from the proxy header even if the proxy header itself is
// syntactically correct.
func (p *Conn) RemoteAddr() net.Addr {
	_ = "STUB: not implemented"
	// Ensure header processing runs at most once.
	return *new(net.Addr)
}

// Raw returns the underlying connection which can be casted to
// a concrete type, allowing access to specialized functions.
//
// Use this ONLY if you know exactly what you are doing.
func (p *Conn) Raw() net.Conn {
	_ = "STUB: not implemented"

	// TCPConn returns the underlying TCP connection,
	// allowing access to specialized functions.
	//
	// Use this ONLY if you know exactly what you are doing.
	return *new(net.Conn)
}

func (p *Conn) TCPConn() (conn *net.TCPConn, ok bool) { _ = "STUB: not implemented"; return nil, false }

// UnixConn returns the underlying Unix socket connection,
// allowing access to specialized functions.
//
// Use this ONLY if you know exactly what you are doing.
func (p *Conn) UnixConn() (conn *net.UnixConn, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// UDPConn returns the underlying UDP connection,
// allowing access to specialized functions.
//
// Use this ONLY if you know exactly what you are doing.
func (p *Conn) UDPConn() (conn *net.UDPConn, ok bool) { _ = "STUB: not implemented"; return nil, false }

// SetDeadline wraps original conn.SetDeadline.
func (p *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline wraps original conn.SetReadDeadline.
func (p *Conn) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	// Set a local var that tells us the desired deadline. This is
	// needed in order to reset the read deadline to the one that is
	// desired by the user, rather than an empty deadline.
	return nil
}

// SetWriteDeadline wraps original conn.SetWriteDeadline.
func (p *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// readHeader reads the proxy protocol header from the connection.
func (p *Conn) readHeader() error {
	_ = "STUB: not implemented"
	// If the connection's readHeaderTimeout is more than 0,
	// apply a temporary deadline without extending a user-configured
	// deadline. If the user has no deadline, we use now + timeout.
	return nil
}

// Clamp to the user's earlier deadline to avoid extending it.

// If the connection's readHeaderTimeout is more than 0, undo the change to the
// deadline that we made above. Because we retain the readDeadline as part of our
// SetReadDeadline override, we can restore the user's deadline (if any).
// Therefore, we check whether the error is a net.Timeout and if it is, we decide
// the proxy proto does not exist and set the error accordingly.

// For the purpose of this wrapper shamefully stolen from armon/go-proxyproto
// let's act as if there was no error when PROXY protocol is not present.

// but not if it is required that the connection has one

// proxy protocol header was found

// this connection is not allowed to send one

// ensureHeaderProcessed runs header processing once.
func (p *Conn) ensureHeaderProcessed() error { _ = "STUB: not implemented"; return nil }

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (p *Conn) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	// Ensure header processing has completed before reading/writing.
	return 0, nil
}

// WriteTo implements io.WriterTo.
func (p *Conn) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	// Ensure header processing has completed before reading/writing.
	return 0, nil
}

// If the buffer has been drained (or cleared), copy directly from conn.

// this should never happen as we read buffered data.
