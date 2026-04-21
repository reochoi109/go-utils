package cli

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	testCases := []struct {
		name      string
		args      []string
		output    string
		wantErr   bool
		errString string
	}{
		{
			name:      "no args",
			args:      []string{},
			wantErr:   true,
			errString: ErrInvalidSubcommand.Error(),
			output:    "invalid sub-command specified\n" + rootUsage,
		},
		{
			name:    "help",
			args:    []string{"-h"},
			wantErr: false,
			output:  rootUsage,
		},
		{
			name:      "invalid subcommand",
			args:      []string{"foo"},
			wantErr:   true,
			errString: ErrInvalidSubcommand.Error(),
			output:    "invalid sub-command specified\n" + rootUsage,
		},
		{
			name:    "http success",
			args:    []string{"http", "http://localhost"},
			wantErr: false,
			output:  "Executing http command\n",
		},
		{
			name:    "grpc success",
			args:    []string{"grpc", "-method", "service.host.local/method", "-body", "{}", "localhost:50051"},
			wantErr: false,
			output:  "Executing grpc command\n",
		},
		{
			name:      "http missing server",
			args:      []string{"http"},
			wantErr:   true,
			errString: ErrNoServerSpecified.Error(),
			output:    "you have to specify the remote server\n" + rootUsage,
		},
		{
			name:      "grpc missing server",
			args:      []string{"grpc"},
			wantErr:   true,
			errString: ErrNoServerSpecified.Error(),
			output:    "you have to specify the remote server\n" + rootUsage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			byteBuf := new(bytes.Buffer)

			err := HandleCommand(byteBuf, tc.args)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if err.Error() != tc.errString {
					t.Fatalf("expected error %q, got %q", tc.errString, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("expected nil error, got %v", err)
				}
			}

			gotOutput := byteBuf.String()
			if gotOutput != tc.output {
				t.Fatalf("expected output:\n%q\n\ngot:\n%q", tc.output, gotOutput)
			}
		})
	}
}
