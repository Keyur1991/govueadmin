package config

import (
	fwconfig "github.com/Keyur1991/go-shreeva/config"
)

func RateLimit() *fwconfig.RateLimit {
	return &fwconfig.RateLimit{
		RequestsPerMinute: 60,
	}
}