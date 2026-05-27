package proxyproto

import (
	"bufio"
	"encoding/binary"
	"errors"
	"net"
)

// maxV2HeaderSize is the maximum acceptable size of a V2 header.
//
// A V2 header may be at most 16 bytes + 64KiB large. We enforce a lower limit
// to mitigate memory allocation DoS while allowing real-world legitimate
// headers. PP2_SUBTYPE_SSL_CLIENT_CERT is typically between 1 and 2KiB, so we
// use a 4KiB limit to leave some room for other TLVs.
const maxV2HeaderSize = 4096

var (
	lengthUnspec      = uint16(0)
	lengthV4          = uint16(12)
	lengthV6          = uint16(36)
	lengthUnix        = uint16(216)
	lengthUnspecBytes = func() []byte {
		a := make([]byte, 2)
		binary.BigEndian.PutUint16(a, lengthUnspec)
		return a
	}()
	lengthV4Bytes = func() []byte {
		a := make([]byte, 2)
		binary.BigEndian.PutUint16(a, lengthV4)
		return a
	}()
	lengthV6Bytes = func() []byte {
		a := make([]byte, 2)
		binary.BigEndian.PutUint16(a, lengthV6)
		return a
	}()
	lengthUnixBytes = func() []byte {
		a := make([]byte, 2)
		binary.BigEndian.PutUint16(a, lengthUnix)
		return a
	}()
	errUint16Overflow = errors.New("proxyproto: uint16 overflow")
)

type _ports struct {
	SrcPort uint16
	DstPort uint16
}

type _addr4 struct {
	Src     [4]byte
	Dst     [4]byte
	SrcPort uint16
	DstPort uint16
}

type _addr6 struct {
	Src [16]byte
	Dst [16]byte
	_ports
}

type _addrUnix struct {
	Src [108]byte
	Dst [108]byte
}

func parseVersion2(reader *bufio.Reader) (header *Header, err error) {
	_ = "STUB: not implemented"
	// Skip first 12 bytes (signature)
	return nil, nil
}

// Read the 13th byte, protocol version and command

// Read the 14th byte, address family and protocol

// UNSPEC is only supported when LOCAL is set.

// Make sure there are bytes available as specified in length

// Return early if the length is zero, which means that
// there's no address information and TLVs present for UNSPEC.

// Length-limited reader for payload section

// Read addresses and ports for protocols other than UNSPEC.
// Ignore address information for UNSPEC, and skip straight to read TLVs,
// since the length is greater than zero.

// Copy bytes for optional Type-Length-Value vector
// Allocate minimum size slice

func (header *Header) formatVersion2() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// For UNSPEC, write no addresses and ports but only TLVs if they are present

//nolint:gosec // Bounds are checked above.

//nolint:gosec // Bounds are checked above.

func (header *Header) validateLength(length uint16) bool { _ = "STUB: not implemented"; return false }

// addTLVLen adds the length of the TLV to the header length or errors on uint16 overflow.
func addTLVLen(cur []byte, tlvLen int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // newLen bounds are validated above.

func newIPAddr(transport AddressFamilyAndProtocol, ip net.IP, port uint16) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func parseUnixName(b []byte) string { _ = "STUB: not implemented"; return "" }

func formatUnixName(name string) []byte { _ = "STUB: not implemented"; return nil }
