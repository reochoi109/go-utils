package cli

import (
	"flag"
	"fmt"
	"io"
)

const rootUsage = `Usage: mync [http|grpc] -h
http: A HTTP client

http: <options> server

Options:
  -verb string
        HTTP method (default "GET")

grpc: A gRPC client.

grpc: <options> server

Options:
  -body string
        Body of request
  -method string
        Method to call
`

func PrintRootUsage(w io.Writer) {
	fmt.Fprint(w, rootUsage)
}

func printHTTPUsage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintln(w, "http: A HTTP client")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "http: <options> server")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Options:")
	fs.PrintDefaults()
}

func printGRPCUsage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintln(w, "grpc: A gRPC client.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "grpc: <options> server")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Options:")
	fs.PrintDefaults()
}
