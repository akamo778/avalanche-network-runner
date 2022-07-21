package local

func GetDefaultFlags() map[string]interface{} {
	return map[string]interface{}{
		"network-peer-list-gossip-frequency": "250ms",
		"network-max-reconnect-delay":        "1s",
		"public-ip":                          "127.0.0.1",
		"health-check-frequency":             "2s",
		"api-admin-enabled":                  true,
		"api-ipcs-enabled":                   true,
		"index-enabled":                      true,
	}
}
