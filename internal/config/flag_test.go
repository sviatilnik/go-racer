package config

import (
	"testing"
	"time"
)

func TestFlagConfig_Init(t *testing.T) {
	conf := NewFlagConfig()
	_, err := conf.Init([]string{})
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if conf.GetTimeout() != time.Second*5 {
		t.Errorf("Incorrect timeout %v, want 5s", conf.GetTimeout())
	}
}

func TestFlagConfig_InitWithTimeout(t *testing.T) {
	conf := NewFlagConfig()
	remaining, err := conf.Init([]string{"-timeout=10s"})
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if conf.GetTimeout() != time.Second*10 {
		t.Errorf("Incorrect timeout %v, want 10s", conf.GetTimeout())
	}

	if len(remaining) != 0 {
		t.Errorf("Expected no remaining args, got %v", remaining)
	}
}

func TestFlagConfig_InitReturnsRemainingArgs(t *testing.T) {
	conf := NewFlagConfig()
	remaining, err := conf.Init([]string{"-timeout=3s", "https://a.com", "https://b.com"})
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	if len(remaining) != 2 {
		t.Errorf("Expected 2 remaining args, got %d: %v", len(remaining), remaining)
	}
}
