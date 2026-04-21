package cli

import (
	"bytes"
	"testing"
)

func TestHandleGRPC(t *testing.T) {
	testCases := []struct {
		name      string
		args      []string
		output    string
		wantErr   bool
		errString string
	}{
		{
			name:      "missing server",
			args:      []string{},
			wantErr:   true,
			errString: ErrNoServerSpecified.Error(),
		},
		{
			name:      "help",
			args:      []string{"-h"},
			wantErr:   true,
			errString: "flag: help requested",
			output:    "grpc: A gRPC client.\n\ngrpc: <options> server\n\nOptions:\n  -body string\n    \tBody of request\n  -method string\n    \tMethod to call\n",
		},
		{
			name:    "success",
			args:    []string{"-method", "service.host.local/method", "-body", "{}", "localhost:50051"},
			wantErr: false,
			output:  "Executing grpc command\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			byteBuf := new(bytes.Buffer)

			err := HandleGRPC(byteBuf, tc.args)

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
