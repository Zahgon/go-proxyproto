package tlvparse

import (
	"github.com/pires/go-proxyproto"
)

// pp2_tlv_ssl.client bit fields.
//
//nolint:revive // Names follow the PROXY protocol spec.
const (
	// PP2_BITFIELD_CLIENT_SSL indicates the client used SSL/TLS.
	PP2_BITFIELD_CLIENT_SSL uint8 = 0x01
	// PP2_BITFIELD_CLIENT_CERT_CONN indicates cert on the connection.
	PP2_BITFIELD_CLIENT_CERT_CONN uint8 = 0x02
	// PP2_BITFIELD_CLIENT_CERT_SESS indicates cert in the session.
	PP2_BITFIELD_CLIENT_CERT_SESS uint8 = 0x04
)

const (
	// tlvSSLMinLen is the minimum length of a SSL TLV.
	tlvSSLMinLen = 5 // len(pp2_tlv_ssl.client) + len(pp2_tlv_ssl.verify)
)

// PP2SSL represents the PP2_TYPE_SSL TLV and its subtypes.
//
// See section 2.2.5 of the PROXY protocol spec.
/*
   struct pp2_tlv_ssl {
           uint8_t  client;
           uint32_t verify;
           struct pp2_tlv sub_tlv[0];
   };
*/
type PP2SSL struct {
	// The Client field is made of a bit field from the following values,
	// indicating which element is present: PP2_BITFIELD_CLIENT_SSL,
	// PP2_BITFIELD_CLIENT_CERT_CONN, PP2_BITFIELD_CLIENT_CERT_SESS
	Client uint8
	// Verify will be zero if the client presented a certificate
	// and it was successfully verified, and non-zero otherwise.
	Verify uint32
	TLV    []proxyproto.TLV
}

// Verified is true if the client presented a certificate and it was successfully verified.
func (s PP2SSL) Verified() bool { _ = "STUB: not implemented"; return false }

// ClientSSL indicates that the client connected over SSL/TLS.  When true, SSLVersion will return the version.
func (s PP2SSL) ClientSSL() bool { _ = "STUB: not implemented"; return false }

// ClientCertConn indicates that the client provided a certificate over the current connection.
func (s PP2SSL) ClientCertConn() bool { _ = "STUB: not implemented"; return false }

// ClientCertSess indicates that the client provided a certificate at least once over the TLS session this
// connection belongs to.
func (s PP2SSL) ClientCertSess() bool { _ = "STUB: not implemented"; return false }

// SSLVersion returns the US-ASCII string representation of the TLS version and whether that extension exists.
func (s PP2SSL) SSLVersion() (string, bool) { _ = "STUB: not implemented"; return "", false }

// SSLCipher returns the US-ASCII string representation of the used TLS cipher and whether that extension exists.
func (s PP2SSL) SSLCipher() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Marshal formats the PP2SSL structure as a TLV.
func (s PP2SSL) Marshal() (proxyproto.TLV, error) {
	_ = "STUB: not implemented"
	return *new(proxyproto.TLV), nil
}

// ClientCN returns the string representation (in UTF8) of the Common Name field (OID: 2.5.4.3) of the client
// certificate's Distinguished Name and whether that extension exists.
func (s PP2SSL) ClientCN() (string, bool) { _ = "STUB: not implemented"; return "", false }

// ClientCert returns the raw X.509 client certificate encoded in ASN.1 DER and
// whether that extension exists.
func (s PP2SSL) ClientCert() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

// IsSSL reports whether the TLV is of SSL type.
func IsSSL(t proxyproto.TLV) bool { _ = "STUB: not implemented"; return false }

// SSL returns the pp2_tlv_ssl from section 2.2.5 or errors with ErrIncompatibleTLV or ErrMalformedTLV.
func SSL(t proxyproto.TLV) (PP2SSL, error) { _ = "STUB: not implemented"; return *new(PP2SSL), nil }

/*
	The PP2_CLIENT_SSL flag indicates that the client connected over SSL/TLS. When
	this field is present, the US-ASCII string representation of the TLS version is
	appended at the end of the field in the TLV format using the type
	PP2_SUBTYPE_SSL_VERSION.
*/

/*
	In all cases, the string representation (in UTF8) of the Common Name field
	(OID: 2.5.4.3) of the client certificate's Distinguished Name, is appended
	using the TLV format and the type PP2_SUBTYPE_SSL_CN. E.g. "example.com".
*/

/*
	The second level TLV PP2_SUBTYPE_SSL_CIPHER provides the US-ASCII string name
	of the used cipher, for example "ECDHE-RSA-AES128-GCM-SHA256".
*/

// FindSSL returns the first PP2SSL if it exists and is well formed.
func FindSSL(tlvs []proxyproto.TLV) (PP2SSL, bool) {
	_ = "STUB: not implemented"
	return *new(PP2SSL), false
}

// isASCII checks whether a byte slice has all characters that fit in the ascii character set, including the null byte.
func isASCII(b []byte) bool { _ = "STUB: not implemented"; return false }
