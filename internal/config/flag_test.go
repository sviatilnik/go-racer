package config

import (
	"os"
	"testing"
	"time"
)

func TestFlagConfig_Init(t *testing.T) {
	conf := NewFlagConfig()
	conf.Init()

	if conf.GetTimeout() != time.Second*5 {
		t.Errorf("Incorrect timeout %d", conf.GetTimeout())
	}
}

func TestFlagConfig_GetTimeout(t *testing.T) {
	originalArgs := os.Args

	defer func() {
		os.Args = originalArgs
	}()

	os.Args = []string{
		originalArgs[0],
		"-timeout=10s",
	}

	conf := NewFlagConfig()
	err := conf.Init()
	if err != nil {
		t.Errorf("Error initializing config: %v", err)
	}

	timeout := conf.GetTimeout()
	if timeout != time.Second*10 {
		t.Errorf("Incorrect timeout %d", timeout)
	}
}
