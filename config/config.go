package config

import (
	"flag"
	"io"
)

type Config interface {
	Mode() string
	Get(key string) string
}

type ConfigSetter interface {
	Set(key string, val string)
}

type option struct {
	mode  string
	extra map[string]string
}

// config
func (o *option) Mode() string          { return o.mode }
func (o *option) Get(key string) string { return o.extra[key] }

// config setter
func (o *option) Set(k string, v string) { o.extra[k] = v }

type OptionFunc func(ConfigSetter)

func New(w io.Writer, args []string, opts ...OptionFunc) (Config, error) {
	opt := &option{
		mode:  "dev",
		extra: make(map[string]string),
	}

	fs := flag.NewFlagSet("config", flag.ContinueOnError)
	fs.SetOutput(w)

	// default option
	fs.StringVar(&opt.mode, "m", opt.mode, "service mode ( dev | prod )")

	// args parse
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	// custom option
	for _, fn := range opts {
		fn(opt)
	}
	return opt, nil
}
