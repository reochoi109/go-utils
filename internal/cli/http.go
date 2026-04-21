package cli

import (
	"flag"
	"fmt"
	"io"
)

type HTTPConfig struct {
	URL  string
	Verb string
}

func HandleHTTP(w io.Writer, args []string) error {
	cfg := HTTPConfig{}

	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&cfg.Verb, "verb", "GET", "HTTP method")
	fs.Usage = func() {
		printHTTPUsage(w, fs)
	}

	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	cfg.URL = fs.Arg(0)

	fmt.Fprintln(w, "Executing http command")
	return nil
}
