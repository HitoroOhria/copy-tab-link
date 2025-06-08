package model_test

import (
	"testing"

	"github.com/HitoroOhria/copy_tab_link/model/value"
)

func parseURL(t *testing.T, rawURL string) *value.URL {
	u, err := value.NewURL(rawURL)
	if err != nil {
		t.Fatalf("value.NewURL: %v", err)
	}

	return u
}
