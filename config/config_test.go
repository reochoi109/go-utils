package config

import (
	"io"
	"testing"
)

// check args
func TestConfig_New(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantMode string
	}{
		{"defaultMode", nil, "dev"},
		{"productMode", []string{"-m", "prod"}, "prod"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt, err := New(io.Discard, tt.args)
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			if opt.Mode() != tt.wantMode {
				t.Fatalf("Mode() = %q , want = %q", opt.Mode(), tt.wantMode)
			}
		})
	}
}

func TestConfig_CustomFunc(t *testing.T) {
	tests := []struct {
		name string
		args []string
		fn   func(ConfigSetter)
		key  string
		want string
	}{
		{"timeout", nil, func(cs ConfigSetter) {
			cs.Set("t", "30")
		}, "t", "30"},
		{"port", nil, func(cs ConfigSetter) {
			cs.Set("p", "8080")
		}, "p", "8080"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt, err := New(io.Discard, tt.args, tt.fn)
			if err != nil {
				t.Fatalf("New() error = %v ", err)
			}
			if got := opt.Get(tt.key); got != tt.want {
				t.Fatalf("GetExtra(%q) = %v, want = %v", tt.key, got, tt.want)
			}
		})
	}

}
