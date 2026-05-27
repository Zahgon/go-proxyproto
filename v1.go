package proxyproto

import (
	"bufio"
	"net"
)

const (
	crlf      = "\r\n"
	separator = " "
)

func initVersion1() *Header { _ = "STUB: not implemented"; return nil }

// Command doesn't exist in v1

func parseVersion1(reader *bufio.Reader) (*Header, error) {
	_ = "STUB: not implemented"
	//The header cannot be more than 107 bytes long. Per spec:
	//
	//   (...)
	//   - worst case (optional fields set to 0xff) :
	//     "PROXY UNKNOWN ffff:f...f:ffff ffff:f...f:ffff 65535 65535\r\n"
	//     => 5 + 1 + 7 + 1 + 39 + 1 + 39 + 1 + 5 + 1 + 5 + 2 = 107 chars
	//
	//   So a 108-byte buffer is always enough to store all the line and a
	//   trailing zero for string processing.
	//
	// It must also be CRLF terminated, as above. The header does not otherwise
	// contain a CR or LF byte.
	//
	// ISSUE #69
	// We can't use Peek here as it will block trying to fill the buffer, which
	// will never happen if the header is TCP4 or TCP6 (max. 56 and 104 bytes
	// respectively) and the server is expected to speak first.
	//
	// Similarly, we can't use ReadString or ReadBytes as these will keep reading
	// until the delimiter is found; an abusive client could easily disrupt a
	// server by sending a large amount of data that do not contain a LF byte.
	// Another means of attack would be to start connections and simply not send
	// data after the initial PROXY signature bytes, accumulating a large
	// number of blocked goroutines on the server. ReadSlice will also block for
	// a delimiter when the internal buffer does not fill up.
	//
	// A plain Read is also problematic since we risk reading past the end of the
	// header without being able to easily put the excess bytes back into the reader's
	// buffer (with the current implementation's design).
	//
	// So we use a ReadByte loop, which solves the overflow problem and avoids
	// reading beyond the end of the header. However, we need one more trick to harden
	// against partial header attacks (slow loris) - per spec:
	//
	//    (..) The sender must always ensure that the header is sent at once, so that
	//    the transport layer maintains atomicity along the path to the receiver. The
	//    receiver may be tolerant to partial headers or may simply drop the connection
	//    when receiving a partial header. Recommendation is to be tolerant, but
	//    implementation constraints may not always easily permit this.
	//
	// We are subject to such implementation constraints. So we return an error if
	// the header cannot be fully extracted with a single read of the underlying
	// reader.
	return nil, nil
}

// End of header found

// No delimiter in first 107 bytes

// Header was not buffered in a single read. Since we can't
// differentiate between genuine slow writers and DoS agents,
// we abort. On healthy networks, this should never happen.

// Check for CR before LF.

// Check full signature.

// Expect at least 2 tokens: "PROXY" and the transport protocol.

// Read address family and protocol

// doesn't exist in v1 but fits UNKNOWN

// Expect 6 tokens only when UNKNOWN is not present.

// When a signature is found, allocate a v1 header with Command set to PROXY.
// Command doesn't exist in v1 but set it for other parts of this library
// to rely on it for determining connection details.

// Transport protocol has been processed already.

// When UNKNOWN, set the command to LOCAL and return early

// Otherwise, continue to read addresses and ports

func (header *Header) formatVersion1() ([]byte, error) {
	_ = "STUB: not implemented"
	// As of version 1, only "TCP4" ( \x54 \x43 \x50 \x34 ) for TCP over IPv4,
	// and "TCP6" ( \x54 \x43 \x50 \x36 ) for TCP over IPv6 are allowed.
	return nil, nil
}

// Unknown connection (short form)

func parseV1PortNumber(portStr string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func parseV1IPAddress(protocol AddressFamilyAndProtocol, addrStr string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// Some proxies (notably nginx OSS stream module) emit plain IPv4
// addresses in TCP6 headers when the backend is IPv4 but the client
// is IPv6. Promote to IPv4-mapped IPv6 for interoperability.
//
// This is an intentional departure from the PROXY protocol v1 spec,
// which states that addresses in a TCP6 line must be in IPv6 format.

// ATTENTION: this is a lossy conversion — round-trip serialization will
// render the address as "::ffff:x.x.x.x" rather than the original "x.x.x.x".
