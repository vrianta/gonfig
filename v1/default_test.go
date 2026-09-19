package gonfig

import (
	"os"
	"testing"
)

// ----------------------------------------------------------------------
// Test Suite 3: Default Values
// ----------------------------------------------------------------------

var testDConfig = New[struct {
	Host    string `env:"HOST" default:"localhost"`
	Port    int    `default:"8080"`
	Enabled bool   `default:"true"`
	Timeout string `default:"30s"`
}](true)

func TestParse_DefaultValues(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Ensure no CLI args interfere with the test
	os.Args = []string{"cmd"}

	tests := []struct {
		field string
		got   any
		want  any
	}{
		{"Host", testDConfig.Host, "localhost"},
		{"Port", testDConfig.Port, 8080},
		{"Enabled", testDConfig.Enabled, true},
		{"Timeout", testDConfig.Timeout, "30s"},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("expected %s to default to %v (%T), got %v (%T)", tt.field, tt.want, tt.want, tt.got, tt.got)
			}
		})
	}
}
