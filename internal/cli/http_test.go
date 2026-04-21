package cli

import (
	"bytes"
	"testing"
)

func TestHandleHTTP(t *testing.T) {
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
			output:    "http: A HTTP client\n\nhttp: <options> server\n\nOptions:\n  -verb string\n    \tHTTP method (default \"GET\")\n",
		},
		{
			name:    "success default verb",
			args:    []string{"http://localhost"},
			wantErr: false,
			output:  "Executing http command\n",
		},
		{
			name:    "success custom verb",
			args:    []string{"-verb", "POST", "http://localhost"},
			wantErr: false,
			output:  "Executing http command\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			byteBuf := new(bytes.Buffer)

			err := HandleHTTP(byteBuf, tc.args)

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
