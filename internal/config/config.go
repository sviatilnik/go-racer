package config

import "time"

type Config interface {
	Init() error
	GetTimeout() time.Duration
}
