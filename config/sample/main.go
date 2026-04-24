package main

import (
	"fmt"
	"io"
	"os"
	"utils/config"
)

func main() {
	opt, err := config.New(io.Discard, os.Args[1:],
		func(s config.ConfigSetter) {
			s.Set("timeout", "30")
		},

		func(s config.ConfigSetter) {
			s.Set("port", "8080")
		},

		func(s config.ConfigSetter) {
			s.Set("db_name", "production_db")
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(opt.Get("timeout"))
}
