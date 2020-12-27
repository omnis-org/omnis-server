package utils

import "github.com/brotherpowers/ipsubnet"

// GetNetworkPart should have a comment.
func GetNetworkPart(ip string, mask int) string {
	sub := ipsubnet.SubnetCalculator(ip, mask)
	return sub.GetNetworkPortion()
}
