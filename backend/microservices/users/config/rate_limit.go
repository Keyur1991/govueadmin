package config

import (
	fwconfig "govueadmin/framework/config"
)

func RateLimit() *fwconfig.RateLimit {
	return &fwconfig.RateLimit{
		RequestsPerMinute: 60,
	}
}