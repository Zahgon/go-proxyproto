// Type-Length-Value splitting and parsing for proxy protocol V2
// See spec https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt sections 2.2 to 2.7 and

package proxyproto

import (
	"errors"
)

// TLV type constants defined by the PROXY protocol spec.
//
//nolint:revive // Names follow the spec.
const (
	// Section 2.2.
	PP2_TYPE_ALPN               PP2Type = 0x01
	PP2_TYPE_AUTHORITY          PP2Type = 0x02
	PP2_TYPE_CRC32C             PP2Type = 0x03
	PP2_TYPE_NOOP               PP2Type = 0x04
	PP2_TYPE_UNIQUE_ID          PP2Type = 0x05
	PP2_TYPE_SSL                PP2Type = 0x20
	PP2_SUBTYPE_SSL_VERSION     PP2Type = 0x21
	PP2_SUBTYPE_SSL_CN          PP2Type = 0x22
	PP2_SUBTYPE_SSL_CIPHER      PP2Type = 0x23
	PP2_SUBTYPE_SSL_SIG_ALG     PP2Type = 0x24
	PP2_SUBTYPE_SSL_KEY_ALG     PP2Type = 0x25
	PP2_SUBTYPE_SSL_GROUP       PP2Type = 0x26
	PP2_SUBTYPE_SSL_SIG_SCHEME  PP2Type = 0x27
	PP2_SUBTYPE_SSL_CLIENT_CERT PP2Type = 0x28
	PP2_TYPE_NETNS              PP2Type = 0x30

	// Section 2.2.7, reserved types.
	PP2_TYPE_MIN_CUSTOM     PP2Type = 0xE0
	PP2_TYPE_MAX_CUSTOM     PP2Type = 0xEF
	PP2_TYPE_MIN_EXPERIMENT PP2Type = 0xF0
	PP2_TYPE_MAX_EXPERIMENT PP2Type = 0xF7
	PP2_TYPE_MIN_FUTURE     PP2Type = 0xF8
	PP2_TYPE_MAX_FUTURE     PP2Type = 0xFF
)

var (
	// ErrTruncatedTLV indicates a TLV was truncated.
	ErrTruncatedTLV = errors.New("proxyproto: truncated TLV")
	// ErrMalformedTLV indicates a TLV has malformed data.
	ErrMalformedTLV = errors.New("proxyproto: malformed TLV Value")
	// ErrIncompatibleTLV indicates a TLV is of an unexpected type.
	ErrIncompatibleTLV = errors.New("proxyproto: incompatible TLV type")
)

// PP2Type is the proxy protocol v2 type.
type PP2Type byte

// TLV is a uninterpreted Type-Length-Value for V2 protocol, see section 2.2.
type TLV struct {
	Type  PP2Type
	Value []byte
}

// SplitTLVs splits the Type-Length-Value vector, returns the vector or an error.
func SplitTLVs(raw []byte) ([]TLV, error) { _ = "STUB: not implemented"; return nil, nil }

// Max length = 65K

// Ignore no-op padding

// JoinTLVs joins multiple Type-Length-Value records.
func JoinTLVs(tlvs []TLV) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // lengthValue is validated above.

// Registered is true if the type is registered in the spec, see section 2.2.
func (p PP2Type) Registered() bool { _ = "STUB: not implemented"; return false }

// App is true if the type is reserved for application specific data, see section 2.2.7.
func (p PP2Type) App() bool { _ = "STUB: not implemented"; return false }

// Experiment is true if the type is reserved for temporary experimental use by application
// developers, see section 2.2.7.
func (p PP2Type) Experiment() bool { _ = "STUB: not implemented"; return false }

// Future is true is the type is reserved for future use, see section 2.2.7.
func (p PP2Type) Future() bool { _ = "STUB: not implemented"; return false }

// Spec is true if the type is covered by the spec, see section 2.2 and 2.2.7.
func (p PP2Type) Spec() bool { _ = "STUB: not implemented"; return false }
