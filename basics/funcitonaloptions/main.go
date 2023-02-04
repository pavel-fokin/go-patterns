package main

import (
	"fmt"
)

type Options struct {
	Host string
	Port int
}

type Service struct {
	options *Options
}

type Setter func(*Options)

func WithHost(host string) Setter {
	return func(opts *Options) {
		opts.Host = host
	}
}

func WithPort(port int) Setter {
	return func(opts *Options) {
		opts.Port = port
	}
}

func New(setters ...Setter) *Service {
	options := &Options{
		Host: "url",
		Port: 8080,
	}

	for _, set := range setters {
		set(options)
	}

	return &Service{
		options: options,
	}
}

func main() {
	fmt.Println("Functional Patterns")
	service := New(
		WithHost("127.0.0.1"),
		WithPort(5000),
	)
	fmt.Printf(
		"%s:%d\n",
		service.options.Host,
		service.options.Port,
	)
}
