package cli

import (
	"flag"
	"fmt"
	"io"
)

type GRPCConfig struct {
	Server string
	Method string
	Body   string
}

func HandleGRPC(w io.Writer, args []string) error {
	cfg := GRPCConfig{}

	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&cfg.Method, "method", "", "Method to call")
	fs.StringVar(&cfg.Body, "body", "", "Body of request")
	fs.Usage = func() {
		printGRPCUsage(w, fs)
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	cfg.Server = fs.Arg(0)

	fmt.Fprintln(w, "Executing grpc command")
	return nil
}
