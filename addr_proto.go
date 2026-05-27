package proxyproto

// AddressFamilyAndProtocol represents address family and transport protocol.
type AddressFamilyAndProtocol byte

// AddressFamilyAndProtocol enum values.
const (
	UNSPEC       AddressFamilyAndProtocol = '\x00'
	TCPv4        AddressFamilyAndProtocol = '\x11'
	UDPv4        AddressFamilyAndProtocol = '\x12'
	TCPv6        AddressFamilyAndProtocol = '\x21'
	UDPv6        AddressFamilyAndProtocol = '\x22'
	UnixStream   AddressFamilyAndProtocol = '\x31'
	UnixDatagram AddressFamilyAndProtocol = '\x32'
)

// IsIPv4 returns true if the address family is IPv4 (AF_INET4), false otherwise.
func (ap AddressFamilyAndProtocol) IsIPv4() bool { _ = "STUB: not implemented"; return false }

// IsIPv6 returns true if the address family is IPv6 (AF_INET6), false otherwise.
func (ap AddressFamilyAndProtocol) IsIPv6() bool { _ = "STUB: not implemented"; return false }

// IsUnix returns true if the address family is UNIX (AF_UNIX), false otherwise.
func (ap AddressFamilyAndProtocol) IsUnix() bool { _ = "STUB: not implemented"; return false }

// IsStream returns true if the transport protocol is TCP or STREAM (SOCK_STREAM), false otherwise.
func (ap AddressFamilyAndProtocol) IsStream() bool { _ = "STUB: not implemented"; return false }

// IsDatagram returns true if the transport protocol is UDP or DGRAM (SOCK_DGRAM), false otherwise.
func (ap AddressFamilyAndProtocol) IsDatagram() bool { _ = "STUB: not implemented"; return false }

// IsUnspec returns true if the transport protocol or address family is unspecified, false otherwise.
func (ap AddressFamilyAndProtocol) IsUnspec() bool { _ = "STUB: not implemented"; return false }

func (ap AddressFamilyAndProtocol) toByte() byte { _ = "STUB: not implemented"; return 0 }
