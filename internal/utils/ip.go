package utils

import (
	"net"

	"github.com/brotherpowers/ipsubnet"
)

// GetNetworkPart should have a comment.
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
