package cli

import (
	"context"
	"testing"
	"time"
)

func TestExecuteCommand_Timeout(t *testing.T) {
	ctx, cancel := CreateContextWithTimeout(100 * time.Millisecond)
	defer cancel()

	if err := ExecuteCommand(ctx, "sleep", "0.5"); err == nil {
		t.Fatalf("Process Success")
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("DeadlineExceeded : %v", ctx.Err())
	}
}

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
		command string
		arg     string
		wantErr bool
	}{
		{"Timeout", 100 * time.Millisecond, "sleep", "0.1", true},
		{"Success", 2 * time.Second, "sleep", "0.1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := CreateContextWithTimeout(tt.timeout)
			defer cancel()

			err := ExecuteCommand(ctx, tt.command, tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
