package utils

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func IsURLValid(URL string) bool {
	URL = strings.TrimSpace(URL)
	if URL == "" {
		return false
	}

	parsed, err := url.Parse(URL)
	return err == nil && parsed.Host != "" && parsed.Scheme != ""
}

func CheckConnection() bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	_, err := client.Head("http://clients3.google.com/generate_204")
	return err == nil
}
