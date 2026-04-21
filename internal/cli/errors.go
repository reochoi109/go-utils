package cli

import "errors"

var (
	ErrNoServerSpecified = errors.New("you have to specify the remote server")
	ErrInvalidSubcommand = errors.New("invalid sub-command specified")
)
