package model_test

import (
	"net/url"
	"testing"
)

func parseURL(t *testing.T, rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("url.Parse: %v", err)
	}

	return u
}
