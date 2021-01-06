package utils

import (
	"testing"
)

func TestGetNetworkPart(t *testing.T) {
	networkPart := GetNetworkPart("192.168.1.1", 24)

	if networkPart != "192.168.1.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", 16)

	if networkPart != "10.1.0.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", 8)

	if networkPart != "10.0.0.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", 32)

	if networkPart != "10.1.2.3" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", 31)

	if networkPart != "10.1.2.2" {
		t.Error("GetNetworkPart failed")
	}

}

func TestGetNetworkPartBadArgument(t *testing.T) {

	networkPart := GetNetworkPart("10.1.2.3", 33)

	if networkPart != "10.1.2.3" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", 0)

	if networkPart != "0.0.0.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.3", -1)

	if networkPart != "0.0.0.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("10.1.2.", 8)

	if networkPart != "0.0.0.0" {
		t.Error("GetNetworkPart failed")
	}

	networkPart = GetNetworkPart("", 8)

	if networkPart != "0.0.0.0" {
		t.Error("GetNetworkPart failed")
	}

}
