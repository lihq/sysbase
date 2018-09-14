package metricsNet

import (
	"bytes"
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

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

func GetLANIpAddrs() (result []string, err error) {
	m := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return m, err
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		//if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !isPrivateSubnet(ipnet.IP) {
		if ipnet, ok := address.(*net.IPNet); ok && isPrivateSubnet(ipnet.IP) {
			if ipnet.IP.To4() != nil {
				ipAddr := ipnet.IP.String()
				m = append(m, ipAddr)
			}
		}
	}

	return m, nil
}
