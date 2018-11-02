package metricsNet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"net"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

// GetLANIpAddrs returns all IPv4 private subnet IP addresses.
func GetLANIpAddrs() (result []string, err error) {
	m := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return m, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && isPrivateSubnet(ipnet.IP) {
			if ipnet.IP.To4() != nil {
				ipAddr := ipnet.IP.String()
				m = append(m, ipAddr)
			}
		}
	}

	return m, nil
}

// GetWANIpAddrs returns all IPv4 non-private subnet IP addresses.
func GetWANIpAddrs() (result []string, err error) {
	m := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return m, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !isPrivateSubnet(ipnet.IP) {
			if ipnet.IP.To4() != nil {
				ipAddr := ipnet.IP.String()
				m = append(m, ipAddr)
			}
		}
	}

	return m, nil
}

// GetFirstIP returns the minimum (in uint32) one IPv4 address.
// In most case, all machines connect to each other in internal network on production,
// multiple machines maybe share a same LAN address such as 192.168.0.1, but share WAN IP will be conflict,
// we prefer use WAN addr to LAN addr.
func GetFirstIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	wans := []net.IP{}
	lans := []net.IP{}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() == nil {
				continue
			}
			if ipnet.IP.IsLoopback() {
				continue
			}

			if isPrivateSubnet(ipnet.IP) {
				lans = append(lans, ipnet.IP)
			} else {
				wans = append(wans, ipnet.IP)
			}
		}
	}

	minOne := uint32(math.MaxUint32)
	for _, ip := range wans {
		inUint32 := binary.BigEndian.Uint32(ip.To4())
		if inUint32 < minOne {
			minOne = inUint32
		}
	}

	if minOne == uint32(math.MaxUint32) {
		for _, ip := range lans {
			inUint32 := binary.BigEndian.Uint32(ip.To4())

			if inUint32 < minOne {
				minOne = inUint32
			}
		}
	}

	if minOne < uint32(math.MaxUint32) {
		ipByte := make([]byte, 4)
		binary.BigEndian.PutUint32(ipByte, minOne)
		return net.IP(ipByte).String(), nil
	}

	return "", errors.New("parse failed")
}
