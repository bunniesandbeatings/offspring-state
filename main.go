package main

import (
	"os"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Debug bool `short:"v" long:"verbose" description:"Debug mode, great for killer regexes" `
}

var globalOptions Options

var parser = flags.NewParser(&globalOptions, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

}
