package config

import (
	"flag"
	"os"
	"time"
)

type FlagConfig struct {
	timeout time.Duration
}

func NewFlagConfig() *FlagConfig {
	return &FlagConfig{}
}

func (c *FlagConfig) Init() error {

	set := flag.NewFlagSet("racer", flag.ContinueOnError)
	timeoutStr := set.String("timeout", "5s", "Race timeout (default - 5s). Example values: 5s, 1m, 2h")
	set.Parse(os.Args[1:])

	timeout, err := time.ParseDuration(*timeoutStr)
	if err != nil {
		return err
	}

	c.timeout = timeout

	return nil
}

func (c *FlagConfig) GetTimeout() time.Duration {
	return c.timeout
}
