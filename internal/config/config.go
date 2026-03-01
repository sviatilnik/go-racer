package config

import "time"

type Config interface {
	Init(args []string) (remainingArgs []string, err error)
	GetTimeout() time.Duration
}
