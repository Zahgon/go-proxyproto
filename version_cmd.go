package proxyproto

// ProtocolVersionAndCommand represents the command in proxy protocol v2.
// Command doesn't exist in v1 but it should be set since other parts of
// this library may rely on it for determining connection details.
type ProtocolVersionAndCommand byte

const (
	// LOCAL represents the LOCAL command in v2 or UNKNOWN transport in v1,
	// in which case no address information is expected.
	LOCAL ProtocolVersionAndCommand = '\x20'
	// PROXY represents the PROXY command in v2 or transport is not UNKNOWN in v1,
	// in which case valid local/remote address and port information is expected.
	PROXY ProtocolVersionAndCommand = '\x21'
)

var supportedCommand = map[ProtocolVersionAndCommand]bool{
	LOCAL: true,
	PROXY: true,
}

// IsLocal returns true if the command in v2 is LOCAL or the transport in v1 is UNKNOWN,
// i.e. when no address information is expected, false otherwise.
func (pvc ProtocolVersionAndCommand) IsLocal() bool { _ = "STUB: not implemented"; return false }

// IsProxy returns true if the command in v2 is PROXY or the transport in v1 is not UNKNOWN,
// i.e. when valid local/remote address and port information is expected, false otherwise.
func (pvc ProtocolVersionAndCommand) IsProxy() bool { _ = "STUB: not implemented"; return false }

// IsUnspec returns true if the command is unspecified, false otherwise.
func (pvc ProtocolVersionAndCommand) IsUnspec() bool {
	_ = "STUB: not implemented"
	// Must be LOCAL or PROXY.
	return false
}

func (pvc ProtocolVersionAndCommand) toByte() byte { _ = "STUB: not implemented"; return 0 }
