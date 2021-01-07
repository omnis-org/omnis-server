package utils

import (
	"net"

	"github.com/brotherpowers/ipsubnet"
)

// GetNetworkPart allows to get the network subnet of an
// ip address : 192.168.1.1/24 => 192.168.1.0
// if invalid ip or mask <= 0 return "0.0.0.0"
// if mask >= 32 return ip
func GetNetworkPart(ip string, mask int) string {
	if net.ParseIP(ip) == nil || mask <= 0 {
		return "0.0.0.0"
	}

	if mask >= 32 {
		return ip
	}

	sub := ipsubnet.SubnetCalculator(ip, mask)
	return sub.GetNetworkPortion()
}
