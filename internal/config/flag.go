package config

import (
	"flag"
	"time"
)

type FlagConfig struct {
	timeout time.Duration
}

func NewFlagConfig() *FlagConfig {
	return &FlagConfig{}
}

func (c *FlagConfig) Init(args []string) ([]string, error) {
	set := flag.NewFlagSet("racer", flag.ContinueOnError)
	timeoutStr := set.String("timeout", "5s", "Race timeout (default - 5s). Example values: 5s, 1m, 2h")
	if err := set.Parse(args); err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		return nil, err
	}

	c.timeout = timeout
	return set.Args(), nil
}

func (c *FlagConfig) GetTimeout() time.Duration {
	return c.timeout
}
