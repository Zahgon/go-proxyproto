package proxyproto

import (
	"net"
)

// PolicyFunc can be used to decide whether to trust the PROXY info from
// upstream. If set, the connecting address is passed in as an argument.
//
// See below for the different policies.
//
// In case an error is returned the connection is denied.
//
// Deprecated: use ConnPolicyFunc instead.
type PolicyFunc func(upstream net.Addr) (Policy, error)

// ConnPolicyFunc can be used to decide whether to trust the PROXY info
// based on connection policy options. If set, the connecting addresses
// (remote and local) are passed in as argument.
//
// See below for the different policies.
//
// In case an error is returned the connection is denied.
type ConnPolicyFunc func(connPolicyOptions ConnPolicyOptions) (Policy, error)

// ConnPolicyOptions contains the remote and local addresses of a connection.
type ConnPolicyOptions struct {
	Upstream   net.Addr
	Downstream net.Addr
}

// Policy defines how a connection with a PROXY header address is treated.
type Policy int

const (
	// USE address from PROXY header.
	USE Policy = iota
	// IGNORE address from PROXY header, but accept connection.
	IGNORE
	// REJECT connection when PROXY header is sent
	// Note: even though the first read on the connection returns an error if
	// a PROXY header is present, subsequent reads do not. It is the task of
	// the code using the connection to handle that case properly.
	REJECT
	// REQUIRE connection to send PROXY header, reject if not present
	// Note: even though the first read on the connection returns an error if
	// a PROXY header is not present, subsequent reads do not. It is the task
	// of the code using the connection to handle that case properly.
	REQUIRE
	// SKIP accepts a connection without requiring the PROXY header.
	// Note: an example usage can be found in the SkipProxyHeaderForCIDR
	// function.
	SKIP
)

// ConnSkipProxyHeaderForCIDR returns a ConnPolicyFunc which can be used to accept
// a connection from a skipHeaderCIDR without requiring a PROXY header, e.g.
// Kubernetes pods local traffic. The def is a policy to use when an upstream
// address doesn't match the skipHeaderCIDR.
func ConnSkipProxyHeaderForCIDR(skipHeaderCIDR *net.IPNet, def Policy) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}

// SkipProxyHeaderForCIDR returns a PolicyFunc which can be used to accept a
// connection from a skipHeaderCIDR without requiring a PROXY header, e.g.
// Kubernetes pods local traffic. The def is a policy to use when an upstream
// address doesn't match the skipHeaderCIDR.
//
// Deprecated: use ConnSkipProxyHeaderForCIDR instead.
func SkipProxyHeaderForCIDR(skipHeaderCIDR *net.IPNet, def Policy) PolicyFunc {
	_ = "STUB: not implemented"
	return *new(PolicyFunc)
}

// WithPolicy adds given policy to a connection when passed as option to NewConn().
func WithPolicy(p Policy) func(*Conn) { _ = "STUB: not implemented"; return nil }

// ConnLaxWhiteListPolicy returns a ConnPolicyFunc which decides whether the
// upstream ip is allowed to send a proxy header based on a list of allowed
// IP addresses and IP ranges. In case upstream IP is not in list the proxy
// header will be ignored. If one of the provided IP addresses or IP ranges
// is invalid it will return an error instead of a ConnPolicyFunc.
func ConnLaxWhiteListPolicy(allowed []string) (ConnPolicyFunc, error) {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc), nil
}

// LaxWhiteListPolicy returns a PolicyFunc which decides whether the
// upstream ip is allowed to send a proxy header based on a list of allowed
// IP addresses and IP ranges. In case upstream IP is not in list the proxy
// header will be ignored. If one of the provided IP addresses or IP ranges
// is invalid it will return an error instead of a PolicyFunc.
//
// Deprecated: use ConnLaxWhiteListPolicy instead.
func LaxWhiteListPolicy(allowed []string) (PolicyFunc, error) {
	_ = "STUB: not implemented"
	return *new(PolicyFunc), nil
}

// ConnMustLaxWhiteListPolicy returns a ConnLaxWhiteListPolicy but will panic
// if one of the provided IP addresses or IP ranges is invalid.
func ConnMustLaxWhiteListPolicy(allowed []string) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}

// MustLaxWhiteListPolicy returns a LaxWhiteListPolicy but will panic if one
// of the provided IP addresses or IP ranges is invalid.
//
// Deprecated: use ConnMustLaxWhiteListPolicy instead.
func MustLaxWhiteListPolicy(allowed []string) PolicyFunc {
	_ = "STUB: not implemented"
	return *new(PolicyFunc)
}

// ConnStrictWhiteListPolicy returns a ConnPolicyFunc which decides whether the
// upstream ip is allowed to send a proxy header based on a list of allowed
// IP addresses and IP ranges. In case upstream IP is not in list reading on
// the connection will be refused on the first read. Please note: subsequent
// reads do not error. It is the task of the code using the connection to
// handle that case properly. If one of the provided IP addresses or IP
// ranges is invalid it will return an error instead of a ConnPolicyFunc.
func ConnStrictWhiteListPolicy(allowed []string) (ConnPolicyFunc, error) {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc), nil
}

// StrictWhiteListPolicy returns a PolicyFunc which decides whether the
// upstream ip is allowed to send a proxy header based on a list of allowed
// IP addresses and IP ranges. In case upstream IP is not in list reading on
// the connection will be refused on the first read. Please note: subsequent
// reads do not error. It is the task of the code using the connection to
// handle that case properly. If one of the provided IP addresses or IP
// ranges is invalid it will return an error instead of a PolicyFunc.
//
// Deprecated: use ConnStrictWhiteListPolicy instead.
func StrictWhiteListPolicy(allowed []string) (PolicyFunc, error) {
	_ = "STUB: not implemented"
	return *new(PolicyFunc), nil
}

// ConnMustStrictWhiteListPolicy returns a ConnStrictWhiteListPolicy but will panic
// if one of the provided IP addresses or IP ranges is invalid.
func ConnMustStrictWhiteListPolicy(allowed []string) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}

// MustStrictWhiteListPolicy returns a StrictWhiteListPolicy but will panic
// if one of the provided IP addresses or IP ranges is invalid.
//
// Deprecated: use ConnMustStrictWhiteListPolicy instead.
func MustStrictWhiteListPolicy(allowed []string) PolicyFunc {
	_ = "STUB: not implemented"
	return *new(PolicyFunc)
}

func connWhitelistPolicy(allowed []func(net.IP) bool, def Policy) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}

// something is wrong with the source IP, better reject the connection

func parse(allowed []string) ([]func(net.IP) bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ipFromAddr(upstream net.Addr) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// TrustProxyHeaderFrom returns a ConnPolicyFunc which can be used to decide
// whether to use or reject PROXY headers based on the source IP of the
// connection. This policy ensures that only trusted sources can set the PROXY
// header. Connections from IPs not in the trusted list will be rejected.
func TrustProxyHeaderFrom(trustedIPs ...net.IP) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}

// IgnoreProxyHeaderNotOnInterface returns a ConnPolicyFunc which can be used to
// decide whether to use or ignore PROXY headers depending on the connection
// being made on specific interfaces. This policy can be used when the server
// is bound to multiple interfaces but wants to allow on one or more interfaces.
func IgnoreProxyHeaderNotOnInterface(allowedIP net.IP) ConnPolicyFunc {
	_ = "STUB: not implemented"
	return *new(ConnPolicyFunc)
}
