package socketio

import "fmt"

// RedisAdapterOptions is configuration to create new adapter
type RedisAdapterOptions struct {
	// deprecated. Usage Addr options
	Host string
	// deprecated. Usage Addr options
	Port     string
	Addr     string
	Username string
	Password string
	Prefix   string
	Network  string
}

func (ro *RedisAdapterOptions) getAddr() string {
	if ro.Addr == "" {
		ro.Addr = fmt.Sprintf("%s:%s", ro.Host, ro.Port)
	}
	return ro.Addr
}

func defaultOptions() *RedisAdapterOptions {
	return &RedisAdapterOptions{
		Addr:    "127.0.0.1:6379",
		Prefix:  "socket.io",
		Network: "tcp",
	}
}

func getOptions(opts *RedisAdapterOptions) *RedisAdapterOptions {
	options := defaultOptions()

	if opts != nil {
		if opts.Host != "" {
			options.Host = opts.Host
		}

		if opts.Port != "" {
			options.Port = opts.Port
		}

		if opts.Addr != "" {
			options.Addr = opts.Addr
		}

		if opts.Username != "" {
			options.Username = opts.Username
		}

		if opts.Password != "" {
			options.Password = opts.Password
		}

		if opts.Prefix != "" {
			options.Prefix = opts.Prefix
		}

		if opts.Network != "" {
			options.Network = opts.Network
		}
	}

	return options
}
