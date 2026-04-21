package cli

import (
	"errors"
	"fmt"
	"io"
)

func HandleCommand(w io.Writer, args []string) error {
	var err error

	if len(args) == 0 {
		err = ErrInvalidSubcommand
	} else {
		switch args[0] {
		case "http":
			err = HandleHTTP(w, args[1:])
		case "grpc":
			err = HandleGRPC(w, args[1:])
		case "-h", "-help", "--help":
			PrintRootUsage(w)
			return nil
		default:
			err = ErrInvalidSubcommand
		}
	}

	if errors.Is(err, ErrNoServerSpecified) || errors.Is(err, ErrInvalidSubcommand) {
		fmt.Fprintln(w, err)
		PrintRootUsage(w)
	}

	return err
}
