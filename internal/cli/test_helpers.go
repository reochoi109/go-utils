package cli

import "errors"

func flagHelpError() error {
	return errors.New("flag: help requested")
}
