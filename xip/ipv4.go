package xip

import (
	"encoding/binary"
	"errors"
	"net"
)

// ErrInvalidIPV4 indicates the given ip string is invalid
var ErrInvalidIPV4 = errors.New("invalid ipv4")

// IPV4 struct represent inet type in postgres sql data type.
// It is a wrapper of net.IP
type IPV4 net.IP

// NewINet creates an INet from net.IP
func NewIPV4(ip net.IP) (IPV4, error) {
	if !IsIPV4(ip) {
		return nil, ErrInvalidIPV4
	}
	return IPV4(ip), nil
}

// BatchToIPV4 converts list of IP to list of INet
func BatchToIPV4(ips []net.IP) ([]IPV4, error) {
	res := make([]IPV4, len(ips))
	for i := range ips {
		if !IsIPV4(ips[i]) {
			return nil, ErrInvalidIPV4
		}
		res[i] = IPV4(ips[i])
	}
	return res, nil
}

func (t IPV4) ToUint32() uint32 {
	return binary.BigEndian.Uint32(t[12:16])
}

// ParseUint32IPV4 converts integer to ipv4.
func ParseUint32IPV4(src uint32) IPV4 {
	segs := make([]byte, 4)
	// convert to
	for i := 0; i < 4; i++ {
		tempInt := src & 0xff
		segs[4-i-1] = byte(tempInt)
		src = src >> 8
	}
	return segs
}

// MarshalText served for json marshalling
func (t IPV4) MarshalText() ([]byte, error) {
	return ((net.IP)(t)).MarshalText()
}

// UnmarshalText try to parse json string
func (t *IPV4) UnmarshalText(text []byte) error {
	ip := net.ParseIP(string(text))
	if ip == nil {
		return ErrInvalidIPV4
	}
	*t = IPV4(ip)
	return nil
}
func (t IPV4) String() string {
	return (net.IP(t)).String()
}
func IsIPV4(ip net.IP) bool {
	return ip.To4() == nil
}
func LocalIPV4() IPV4 {
	var src = Local()
	// IP 没实现==,但是实现了 Equal
	if net.IPv4zero.Equal(src) {
		return IPV4(net.IPv4zero)
	}
	return IPV4(src)
}
func Local() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IPv4zero
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return net.IPv4zero
}
