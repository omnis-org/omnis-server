package utils

import "github.com/brotherpowers/ipsubnet"

func GetNetworkPart(ip string, mask int) string {
	sub := ipsubnet.SubnetCalculator(ip, mask)
	return sub.GetNetworkPortion()
}
